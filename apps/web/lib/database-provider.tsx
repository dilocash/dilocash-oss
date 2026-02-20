"use client";
import React, { PropsWithChildren, useEffect, useState } from 'react';
import { Database } from '@nozbe/watermelondb';
import LokiJSAdapter from '@nozbe/watermelondb/adapters/lokijs';
import schema from '@dilocash/database/local/model/schema';
import { Intent } from '@dilocash/database/local/model/intent';
import { Command } from '@dilocash/database/local/model/commmand';
import { Transaction } from '@dilocash/database/local/model/transaction';
import { DatabaseProvider } from '@nozbe/watermelondb/react';

const CustomeDatabaseProvider = ({ children }: PropsWithChildren) => {
  const [database, setDatabase] = useState<Database | null>(null);

  useEffect(() => {
    (async () => {

      const adapter = new LokiJSAdapter({
        useWebWorker: false,
        useIncrementalIndexedDB: true,
        dbName: "dilocash",
        schema,
      });

      const db = new Database({
        adapter,
        modelClasses: [Command, Intent, Transaction],
      });

      setDatabase(db);
    })();
  }, []);

  return (
    database && (
      <DatabaseProvider database={database}>{children}</DatabaseProvider>
    )
  );
};

export default CustomeDatabaseProvider;