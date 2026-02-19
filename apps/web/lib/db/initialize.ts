// apps/web/lib/db/initialize.ts

import { getRxStorageDexie } from "rxdb/plugins/storage-dexie";
import { initDB as initLocalDB } from "@dilocash/db-local/database";
import { createConnectTransport } from "@connectrpc/connect-web";

const transport = createConnectTransport({
  baseUrl: "https://demo.connectrpc.com",
});

export const initDB = async (session: string) => {
  const db = await initLocalDB(getRxStorageDexie(), transport, session);
  return db;
};