import { replicateRxCollection } from 'rxdb/plugins/replication';
import { createClient } from "@connectrpc/connect";
import { TransactionService } from "@dilocash/gen/ts/transport/dilocash/v1/transaction_service_pb";

export function setupTransactionReplication(collection: any, transport: any, session: any) {
  const client = createClient(TransactionService, transport);

  return replicateRxCollection({
    collection,
    replicationIdentifier: 'dilocash-sync-grpc',
    // PUSH: Envía cambios locales al servidor Go
    push: {
      handler: async (docs) => {
        try {
          //await client.syncTransactions({ transactions: docs });
          console.log("Sincronizando con gRPC:", docs);
          return []; // Si no hay errores, devolvemos array vacío
        } catch (err) {
          console.error("Error sincronizando con gRPC:", err);
          throw err;
        }
      },
      batchSize: 5
    },
    // PULL: Trae nuevos cambios desde el servidor Go
    pull: {
      handler: async (lastCheckpoint) => {
        // const response = await client.getUpdates({ 
        //   lastTimestamp: lastCheckpoint?.updatedAt || 0 
        // });
        // return {
        //   documents: response.transactions,
        //   checkpoint: { updatedAt: response.newTimestamp }
        // };
        console.log("Obteniendo actualizaciones desde gRPC:", lastCheckpoint);
        return {
          documents: [],
          checkpoint: { updatedAt: 0 }
        };
      }
    },
    live: true, // Replicación en tiempo real
    retryTime: 5000 // Si falla el gRPC, reintenta en 5s
  });
}