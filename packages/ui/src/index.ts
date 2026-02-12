// packages/ui/src/index.ts

'use client';
// export Gluestack components
export * from "../components/ui/gluestack-ui-provider";

// Export basic components
export * from "../components/ui/box";
export * from "../components/ui/vstack";
export * from "../components/ui/button";
export * from "../components/ui/input";
export * from "../components/ui/heading";
export * from "../components/ui/text";
export * from "../tailwind.config";

// shared components
export { AuthForm } from "./components/auth/auth-form";