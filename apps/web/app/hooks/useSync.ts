import { synchronize } from "@nozbe/watermelondb/sync";
import { create } from "@bufbuild/protobuf";
import { TimestampSchema } from "@bufbuild/protobuf/wkt";
import { PullChangesRequest, PullChangesResponse } from "@dilocash/gen/ts/transport/dilocash/v1/sync_types_pb";
import { useDatabase } from "@nozbe/watermelondb/react";
import { useState } from "react";

import { createClient } from "@connectrpc/connect";
import { SyncService } from '@dilocash/gen/ts/transport/dilocash/v1/sync_service_pb';
import { getSupabaseClient } from '@dilocash/ui/auth/client';
import { createConnectTransport } from "@connectrpc/connect-web";
import { PullChangesRequestSchema } from "@dilocash/gen/ts/transport/dilocash/v1/sync_types_pb";

const BASE_URL = "http://localhost:8080";

const useSync = () => {
  const transport = createConnectTransport({
    baseUrl: BASE_URL,
    interceptors: [
    (next) => async (req) => {
      const supabase = getSupabaseClient(
        process.env.NEXT_PUBLIC_SUPABASE_URL!, 
        process.env.NEXT_PUBLIC_SUPABASE_PUBLISHABLE_KEY!,
        localStorage);
      const { data } = await supabase.auth.getSession();
      
      if (data.session?.access_token) {
        req.header.set("Authorization", `Bearer ${data.session.access_token}`);
      }
      return await next(req);
    },
  ],
  });
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
        pullChanges: async ({ lastPulledAt, schemaVersion, migration }) => {
          // 1. Llamada a Go vía gRPC
          const lastPulledAtTimestamp = lastPulledAt
            ? create(TimestampSchema, {
                seconds: BigInt(Math.floor(lastPulledAt / 1000)),
                nanos: (lastPulledAt % 1000) * 1_000_000,
              })
            : undefined;

          const pullRequest: PullChangesRequest = create(PullChangesRequestSchema, {
            lastPulledAt: lastPulledAtTimestamp,
            schemaVersion: schemaVersion,
            migration: JSON.stringify(migration),
            limit: 100,
          });

          const response = await client.pullChanges(
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

          return {
            changes: changes as any, // Cast to any to bypass strict proto vs watermelondb type comparison
            timestamp,
          };
        },
        pushChanges: async ({ changes, lastPulledAt }) => {
          // const response = await fetch(
          //   `${BASE_URL}?last_pulled_at=${lastPulledAt}`,
          //   {
          //     method: "POST",
          //     body: JSON.stringify(changes),
          //   }
          // );
          // if (!response.ok) {
          //   throw new Error(await response.text());
          // }
        },
      });
    } catch (e) {
      console.log(e);
    }

    setIsSyncing(false);
  };

  return { isSyncing, sync };
};

export default useSync;