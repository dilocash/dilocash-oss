import { synchronize, SyncPullResult } from "@nozbe/watermelondb/sync";
import { create } from "@bufbuild/protobuf";
import { TimestampSchema, Timestamp } from "@bufbuild/protobuf/wkt";
import { Changes, ChangesSchema, PullChangesRequest, PullChangesResponse, PushChangesRequest, PushChangesRequestSchema, PushChangesResponse } from "@dilocash/gen/ts/transport/dilocash/v1/sync_types_pb";
import { useDatabase } from "@nozbe/watermelondb/react";
import { useState } from "react";

import { createClient, Transport } from "@connectrpc/connect";
import { SyncService } from '@dilocash/gen/ts/transport/dilocash/v1/sync_service_pb';
import { PullChangesRequestSchema } from "@dilocash/gen/ts/transport/dilocash/v1/sync_types_pb";


const useSync = (transport: Transport) => {
  const client = createClient(SyncService, transport);
  const [isSyncing, setIsSyncing] = useState(false);
  const database = useDatabase();

  const sync = async () => {
    console.log("before sync");
    if (isSyncing) {
      console.info("Sync already in progress");
      return;
    }
    console.log("after sync");

    console.info("Syncing...");

    setIsSyncing(true);
    try {
      await synchronize({
        database,
        sendCreatedAsUpdated: true,
        pullChanges: async ({ lastPulledAt, schemaVersion, migration }) => {
          console.info("pullChanges", {lastPulledAt, schemaVersion, migration});
          const lastPulledAtTimestamp = lastPulledAt
            ? toProtoTimestamp(lastPulledAt)
            : undefined;

          const pullRequest: PullChangesRequest = create(PullChangesRequestSchema, {
            lastPulledAt: lastPulledAtTimestamp,
            schemaVersion: schemaVersion,
            migration: JSON.stringify(migration),
            limit: 100,
          });

          const response: PullChangesResponse = await client.pullChanges(
            pullRequest
          );

          // WatermelonDB expects an object with { changes, timestamp }
          // We must ensure 'changes' is not undefined and map the proto timestamp to a number (ms)
          const changes = {
            commands: response.changes?.commands ?? { created: [], updated: [], deleted: [] },
            intents: response.changes?.intents ?? { created: [], updated: [], deleted: [] },
            transactions: response.changes?.transactions ?? { created: [], updated: [], deleted: [] },
          };

          const timestamp = response.timestamp
            ? Number(response.timestamp.seconds) * 1000 + Math.floor(response.timestamp.nanos / 1_000_000)
            : Date.now();

          return { changes, timestamp } as SyncPullResult;
        },
        pushChanges: async ({ changes, lastPulledAt }) => {
          console.info("pushChanges", {lastPulledAt, changes});

          const lastPulledAtTimestamp = lastPulledAt
            ? toProtoTimestamp(lastPulledAt)
            : undefined;

          const pushRequest: PushChangesRequest = create(PushChangesRequestSchema, {
            lastPulledAt: lastPulledAtTimestamp,
            changes: changesToProto(changes),
          });

          const response: PushChangesResponse = await client.pushChanges(
            pushRequest
          );
          if (!response.ok) {
            throw new Error(response.conflictIds.join(", "));
          }
          
        },
      });
    } catch (e) {
      console.error(e);
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
const dbObjToProtoObj = (obj: any): any => {
    const camelCaseObj = Object.fromEntries(
        Object.entries(obj).map(([key, value]) => [toCamelCase(key), value])
    );
    if (typeof camelCaseObj.createdAt === 'number') {
        camelCaseObj.createdAt = toProtoTimestamp(camelCaseObj.createdAt);
    }
    if (typeof camelCaseObj.updatedAt === 'number') {
        camelCaseObj.updatedAt = toProtoTimestamp(camelCaseObj.updatedAt);
    }
    return camelCaseObj;
}

const toProtoTimestamp = (timestamp: number): Timestamp => {
    return create(TimestampSchema, {
        seconds: BigInt(Math.floor(timestamp / 1000)),
        nanos: (timestamp % 1000) * 1_000_000,
      });
}

const toCamelCase = (str: string): string => {
    return str.replace(/_([a-z])/g, (g) => g[1].toUpperCase());
}

const changesToProto = (obj: any): any => {
    return {
      commands : {
        created: obj.commands.created.map((command: any) => dbObjToProtoObj(command)),
        updated: obj.commands.updated.map((command: any) => dbObjToProtoObj(command)),
        deleted: obj.commands.deleted,
      },
      intents : {
        created: obj.intents.created.map((intent: any) => dbObjToProtoObj(intent)),
        updated: obj.intents.updated.map((intent: any) => dbObjToProtoObj(intent)),
        deleted: obj.intents.deleted,
      },
      transactions : {
        created: obj.transactions.created.map((transaction: any) => dbObjToProtoObj(transaction)),
        updated: obj.transactions.updated.map((transaction: any) => dbObjToProtoObj(transaction)),
        deleted: obj.transactions.deleted,
      },
    };
}

export default useSync;