import { createRxDatabase, addRxPlugin } from 'rxdb';
import { RxDBDevModePlugin } from 'rxdb/plugins/dev-mode';
import { TransactionSchema } from './schemas/transaction';
import { setupTransactionReplication } from './replication';
import { Transport } from '@connectrpc/connect';

if (process.env.NODE_ENV === 'development') {
  addRxPlugin(RxDBDevModePlugin);
}

export const initDB = async (adapter: any, transport: Transport, session: string) => {
  const db = await createRxDatabase({
    name: 'dilocash_db' + session,
    storage: adapter,
  });

  await db.addCollections({
    transactions: { schema: TransactionSchema }
  });

  setupTransactionReplication(db.transactions, transport, session);

  return db;
};