import { synchronize } from "@nozbe/watermelondb/sync";
import { create } from "@bufbuild/protobuf";
import { TimestampSchema } from "@bufbuild/protobuf/wkt";
import { PullCommandsRequestSchema, PullCommandsResponse } from "@dilocash/gen/ts/transport/dilocash/v1/command_types_pb";
import { useDatabase } from "@nozbe/watermelondb/react";
import { useState } from "react";

import { createClient } from "@connectrpc/connect";
import { CommandService } from '@dilocash/gen/ts/transport/dilocash/v1/command_service_pb';

import { createConnectTransport } from "@connectrpc/connect-web";

const BASE_URL = "http://localhost:8000/sync";

const useSync = () => {
  const transport = createConnectTransport({
    baseUrl: BASE_URL,
  });
  const client = createClient(CommandService, transport);
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
        pullChanges: async ({ lastPulledAt }) => {
          // 1. Llamada a Go vía gRPC
          const lastUpdatedAt = lastPulledAt
            ? create(TimestampSchema, {
                seconds: BigInt(Math.floor(lastPulledAt / 1000)),
                nanos: (lastPulledAt % 1000) * 1_000_000,
              })
            : undefined;

          const response = await client.pullCommands(
            create(PullCommandsRequestSchema, {
              lastUpdatedAt,
              limit: 100,
            })
          ) as PullCommandsResponse;
          // Watermelon espera un objeto con: { changes, timestamp }
          // response.commands contains the synced commands from the server
          // response.checkpointUpdatedAt is the new checkpoint (milliseconds)
          return {
            changes: { commands: { created: response.commands, updated: [], deleted: [] } },
            timestamp: Number(response.checkpointUpdatedAt),
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