# ADR 044: Local Database Selection — WatermelonDB over RxDB

## Status

- **Status**: Accepted
- **Date**: 2026-03-07
- **Authors**: @jalbarran
- **Technical Domain**: Frontend / Mobile / Web
- **Supersedes**: ADR 043 (partially — storage layer only)

## Context

ADR 043 established an offline-first architecture using **RxDB** as the local database, with `storage-dexie` (IndexedDB) for Web and `storage-sqlite` (via `expo-sqlite`) for Mobile.

During implementation, two blockers emerged with RxDB:

1. **Licensing:** RxDB's advanced features (including the SQLite storage adapter needed for Expo/Mobile) are gated behind the **RxDB Premium** license. The free/OSS tier does not include the adapters required for React Native / Expo environments, making it incompatible with Dilocash's Open Core model and OSS repository requirements.

2. **Cost:** The RxDB Premium license carries a recurring per-developer or per-project fee that is not sustainable for a bootstrapped, early-stage product. This conflicts directly with ADR 032 (Cost-Optimized Infrastructure) and ADR 021 (Dual Licensing and Repository Separation).

We therefore need to select an alternative local database that:

- Has a fully **Open Source license** (no premium tiers for core adapters).
- Supports **React Native / Expo** (SQLite-backed) for mobile.
- Supports **Web (Next.js / PWA)** without native dependencies.
- Is production-proven and actively maintained.

## Decision

We will replace RxDB with **WatermelonDB** as the local database layer across the stack.

The storage adapters will be platform-specific, consistent with the polyglot persistence strategy established in ADR 043:

* **Web (Next.js):** **LokiJS** — an in-memory / localStorage-backed database that operates fully in the browser with no native bindings. It is MIT-licensed and requires zero native compilation.
* **Mobile (Android/iOS via Expo):** **SQLite** — via `@nozbe/watermelondb/adapters/sqlite` using `expo-sqlite`, the same underlying storage as originally planned in ADR 043.

WatermelonDB itself is **MIT-licensed**, fulfilling the OSS requirement for the Open Core repository.

### 1. Bundler Constraint & Turbopack Compatibility

The [official WatermelonDB documentation](https://watermelondb.dev/docs/Installation) states that WatermelonDB requires Webpack due to its reliance on **TypeScript decorators** (`@field`, `@text`, `@date`, `@children`, `@relation`, `@writer`). These decorators use `emitDecoratorMetadata`, a TypeScript-only feature that Turbopack's SWC-based compiler does not support.

**We resolved this constraint** by eliminating the decorator dependency entirely from the WatermelonDB model layer, making Turbopack fully compatible.

#### Solution: Decorator-Free WatermelonDB Models

All WatermelonDB Model classes in `packages/database/local/model/` were refactored to use plain TypeScript getters, setters, and explicit method calls that replicate the exact internal behavior of each decorator:

| Decorator | Replacement Strategy |
| --- | --- |
| `@field("column")` | `get` / `set` using `_getRaw()` / `_setRaw()` |
| `@text("column")` | Same as `@field` with `?? ""` null coercion |
| `@readonly @date("column")` | Read-only `get` converting raw `number` → `new Date()` |
| `@children("table")` | `this.collections.get("table").query(Q.where(foreignKey, this.id))` |
| `@relation("table", "col")` | `new Relation(this, table, col, { isImmutable: false })` with instance cache |
| `@writer async method()` | `await this.database.write(async () => { ... }, "Model.method")` |

The `experimentalDecorators` and `emitDecoratorMetadata` flags were also removed from `apps/web/tsconfig.json`.

Turbopack was then enabled in `apps/web/next.config.ts` via the top-level `turbopack: {}` property (the stable API in Next.js 16):

```ts
// apps/web/next.config.ts
const nextConfig: NextConfig = {
  turbopack: {},
  // ...
};
```

This approach was verified to work: the development server starts as **`▲ Next.js 16.1.6 (Turbopack)`** with no build errors.

> **Note:** The mobile app (`apps/mobile`) is bundled by **Metro** (Expo's bundler), which fully supports TypeScript decorators via Babel. The decorator-free refactor applies **only** to `packages/database` — the shared model package consumed by the web app. The mobile `database-provider.tsx` is unaffected.

### 2. Polyglot Persistence (Storage Layer)

| Platform | Adapter | Underlying Storage | Bundler |
| --- | --- | --- | --- |
| **Web (Next.js)** | WatermelonDB + LokiJS Adapter | In-memory / localStorage | **Turbopack** ✅ |
| **Mobile (Expo)** | WatermelonDB + SQLite Adapter | `expo-sqlite` | Metro |

### 3. Data Contract (Single Source of Truth)

The local database schema continues to be derived from **Protocol Buffer (.proto)** definitions, as established in ADR 043.

* WatermelonDB **Model** classes will implement the same field structure as the Protobuf-generated TypeScript types.
* Schema changes on the Go backend will still propagate as type mismatches in the frontend, enforcing schema discipline.

### 4. Replication Mechanism

The custom replication logic will be re-implemented using WatermelonDB's `synchronize()` API and Connect-go services:

* **Push:** Local `created`, `updated`, and `deleted` records are sent in batches to the gRPC endpoint.
* **Pull:** The client fetches server changes based on a `lastPulledAt` timestamp (checkpoint).
* **Frequency:** Live synchronization when online; exponential backoff during outages.

### 5. Service Worker Integration (PWA)

The **Serwist** integration established in ADR 043 remains unchanged. Background Sync tasks will trigger the WatermelonDB `synchronize()` call when the browser tab is closed or the connection is restored.

## Consequences

### Positive

* **Fully OSS:** WatermelonDB (MIT), LokiJS (MIT), and `expo-sqlite` carry no licensing fees or premium tier restrictions.
* **Cost Elimination:** Removes the RxDB Premium licensing cost entirely, consistent with ADR 032.
* **Turbopack Enabled:** By removing the decorator dependency, the Next.js web app benefits from Turbopack's significantly faster incremental compilation and HMR during local development.
* **Performance:** WatermelonDB is designed from the ground up for React Native performance, with lazy loading and direct SQLite access on mobile avoiding the overhead of RxDB's reactive pipeline.
* **Proven in Production:** WatermelonDB is used in production by Nozbe and other large-scale apps; LokiJS is widely used for in-browser storage.

### Negative / Risks

* **LokiJS Persistence on Web:** The LokiJS adapter stores data in memory by default; persistence to `localStorage` must be explicitly configured. Data may be lost if the user clears browser storage. This risk is mitigated by cloud sync ensuring data is always recoverable from the server.
* **Decorator-Free Verbosity:** The refactored models are more verbose than the decorator-based originals. Any new model property must be written as an explicit getter/setter pair. This is a one-time ergonomic cost with no runtime trade-off.
* **No Built-in Reactivity:** Unlike RxDB's Observable streams, WatermelonDB uses a `withObservables` HOC pattern for reactivity. UI components must be wrapped accordingly.
* **Different Migration API:** WatermelonDB uses a migration runner distinct from RxDB's. The migration versioning strategy from ADR 043 must be adapted to WatermelonDB's `addColumns` / `createTable` migration DSL.

## Technical Implementation Blueprint

| Component | Technology | Monorepo Location |
| --- | --- | --- |
| **Schemas / Models** | WatermelonDB Models (decorator-free) | `packages/database/local/model/` |
| **Web Adapter** | WatermelonDB + LokiJS | `apps/web/lib/database-provider.tsx` |
| **Mobile Adapter** | WatermelonDB + SQLite | `apps/mobile/src/app/lib/database-provider.tsx` |
| **Transport** | Connect-go (HTTP/2) | `packages/api-client` |
| **UI State** | `withObservables` HOC | `apps/web` & `apps/mobile` |
| **PWA Logic** | Serwist + Workbox | `apps/web/sw.ts` |

---

### Synchronization Workflow

1. **User Input:** Form → WatermelonDB (Local Insert via `database.write()`).
2. **Reactivity:** `withObservables` HOC → UI Update (Shows "Pending" icon).
3. **Sync Trigger:** `synchronize()` → Connect-go Client → AWS Fargate (Go).
4. **Acknowledgment:** Fargate → Postgres (Commit) → gRPC Response.
5. **Finality:** WatermelonDB (marks records as synced) → UI Update (Shows "Check" icon).

---
