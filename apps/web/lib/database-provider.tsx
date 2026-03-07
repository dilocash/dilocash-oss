/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

"use client";
import React, { PropsWithChildren, useEffect, useState } from 'react';
import { Database } from '@nozbe/watermelondb';
import LokiJSAdapter from '@nozbe/watermelondb/adapters/lokijs';
import schema from '@dilocash/database/local/model/schema';
import { Intent } from '@dilocash/database/local/model/intent';
import { Command } from '@dilocash/database/local/model/commmand';
import { Transaction } from '@dilocash/database/local/model/transaction';
import { DatabaseProvider } from '@nozbe/watermelondb/react';
import { setGenerator } from "@nozbe/watermelondb/utils/common/randomId";
import { v4 as uuidv4 } from "uuid";

const CustomeDatabaseProvider = ({ children }: PropsWithChildren) => {
  const [database, setDatabase] = useState<Database | null>(null);

  useEffect(() => {
    (async () => {
      setGenerator(() => uuidv4());

      const adapter = new LokiJSAdapter({
        useWebWorker: false,
        useIncrementalIndexedDB: true,
        dbName: "dilocash",
        schema,
        onSetUpError: (error: Error) => {
          console.log("error setting up database", error);
        },
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