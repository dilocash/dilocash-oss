/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { synchronize, SyncPullResult } from "@nozbe/watermelondb/sync";
import { create } from "@bufbuild/protobuf";
import { TimestampSchema, Timestamp } from "@bufbuild/protobuf/wkt";
import {
  Changes,
  PullChangesRequest,
  PullChangesRequestSchema,
  PullChangesResponse,
  PushChangesRequest,
  PushChangesRequestSchema,
  PushChangesResponse,
} from "@dilocash/gen/ts/transport/dilocash/v1/sync_types_pb";
import { useDatabase } from "@nozbe/watermelondb/react";
import { useState } from "react";

import { createClient, Transport } from "@connectrpc/connect";
import { SyncService } from "@dilocash/gen/ts/transport/dilocash/v1/sync_service_pb";

const useSync = (transport: Transport) => {
  const client = createClient(SyncService, transport);
  const [isSyncing, setIsSyncing] = useState(false);
  const database = useDatabase();

  const sync = async () => {
    if (isSyncing) {
      console.info("Sync already in progress");
      return;
    }
    console.info("Syncing...");

    setIsSyncing(true);
    try {
      await synchronize({
        database,
        sendCreatedAsUpdated: false,
        pullChanges: async ({ lastPulledAt, schemaVersion, migration }) => {
          console.info("pullChanges", {
            lastPulledAt,
            schemaVersion,
            migration,
          });
          const lastPulledAtTimestamp = lastPulledAt
            ? toProtoTimestamp(lastPulledAt)
            : undefined;

          const pullRequest: PullChangesRequest = create(
            PullChangesRequestSchema,
            {
              lastPulledAt: lastPulledAtTimestamp,
              schemaVersion: schemaVersion,
              migration: JSON.stringify(migration),
              limit: 100,
            },
          );

          const response: PullChangesResponse =
            await client.pullChanges(pullRequest);

          // WatermelonDB expects an object with { changes, timestamp }
          // We must ensure 'changes' is not undefined and map the proto timestamp to a number (ms)
          const changes = changesToDB(response.changes);

          return {
            changes,
            timestamp: toDbTimestamp(response.timestamp as Timestamp),
          } as SyncPullResult;
        },
        pushChanges: async ({ changes, lastPulledAt }) => {
          console.info("pushChanges", { lastPulledAt, changes });

          const lastPulledAtTimestamp = lastPulledAt
            ? toProtoTimestamp(lastPulledAt)
            : undefined;

          const pushRequest: PushChangesRequest = create(
            PushChangesRequestSchema,
            {
              lastPulledAt: lastPulledAtTimestamp,
              changes: changesToProto(changes),
            },
          );

          const response: PushChangesResponse =
            await client.pushChanges(pushRequest);
          if (!response.ok) {
            throw new Error(response.conflictIds.join(", "));
          }
        },
      });
    } catch (e) {
      console.error('sync error', e);
    }

    setIsSyncing(false);
  };

  return { isSyncing, sync };
};

/**
 * Convert db-like objects (item_status) to protobuf-like object (itemStatus)
 * @param obj db-like object
 * @returns protobuf-like object
 */
const dbToProto = (obj: any): any => {
  const camelCaseObj = Object.fromEntries(
    Object.entries(obj).map(([key, value]) => [toCamelCase(key), value]),
  );
  if (typeof camelCaseObj.createdAt === "number") {
    camelCaseObj.createdAt = toProtoTimestamp(camelCaseObj.createdAt);
  }
  if (typeof camelCaseObj.updatedAt === "number") {
    camelCaseObj.updatedAt = toProtoTimestamp(camelCaseObj.updatedAt);
  }
  return camelCaseObj;
};

const protoToDb = (obj: any): any => {
  const camelCaseObj = Object.fromEntries(
    Object.entries(obj).map(([key, value]) => [toSnakeCase(key), value]),
  );
  if (camelCaseObj.created_at) {
    camelCaseObj.created_at = toDbTimestamp(
      camelCaseObj.created_at as Timestamp,
    );
  }
  if (camelCaseObj.updated_at) {
    camelCaseObj.updated_at = toDbTimestamp(
      camelCaseObj.updated_at as Timestamp,
    );
  }
  return camelCaseObj;
};

const commandToProto = (obj: any): any => {
  return dbToProto(obj);
};

const commandToDb = (obj: any): any => {
  return protoToDb(obj);
};

const intentToProto = (obj: any): any => {
  return dbToProto(obj);
};

const intentToDb = (obj: any): any => {
  return protoToDb(obj);
};

const transactionToProto = (obj: any): any => {
  return dbToProto(obj);
};

const transactionToDb = (obj: any): any => {
  return protoToDb(obj);
};

const toProtoTimestamp = (timestamp: number): Timestamp => {
  return create(TimestampSchema, {
    seconds: BigInt(Math.floor(timestamp / 1000)),
    nanos: (timestamp % 1000) * 1_000_000,
  });
};

const toDbTimestamp = (timestamp: Timestamp): number => {
  return timestamp
    ? Number(timestamp.seconds) * 1000 + Math.floor(timestamp.nanos / 1_000_000)
    : Date.now();
};

const toCamelCase = (str: string): string => {
  return str.replace(/_([a-z])/g, (g) => g[1].toUpperCase());
};

const toSnakeCase = (str: string): string => {
  return str.replace(/[A-Z]/g, (g) => `_${g.toLowerCase()}`);
};

const changesToProto = (obj: any): any => {
  return {
    commands: {
      created: obj.commands.created.map((command: any) =>
        commandToProto(command),
      ),
      updated: obj.commands.updated.map((command: any) =>
        commandToProto(command),
      ),
      deleted: obj.commands.deleted,
    },
    intents: {
      created: obj.intents.created.map((intent: any) => intentToProto(intent)),
      updated: obj.intents.updated.map((intent: any) => intentToProto(intent)),
      deleted: obj.intents.deleted,
    },
    transactions: {
      created: obj.transactions.created.map((transaction: any) =>
        transactionToProto(transaction),
      ),
      updated: obj.transactions.updated.map((transaction: any) =>
        transactionToProto(transaction),
      ),
      deleted: obj.transactions.deleted,
    },
  };
};

const changesToDB = (obj: any): any => {
  return {
    commands: {
      created: obj.commands.created.map((command: any) => commandToDb(command)),
      updated: obj.commands.updated.map((command: any) => commandToDb(command)),
      deleted: obj.commands.deleted,
    },
    intents: {
      created: obj.intents.created.map((intent: any) => intentToDb(intent)),
      updated: obj.intents.updated.map((intent: any) => intentToDb(intent)),
      deleted: obj.intents.deleted,
    },
    transactions: {
      created: obj.transactions.created.map((transaction: any) =>
        transactionToDb(transaction),
      ),
      updated: obj.transactions.updated.map((transaction: any) =>
        transactionToDb(transaction),
      ),
      deleted: obj.transactions.deleted,
    },
  };
};

export default useSync;
