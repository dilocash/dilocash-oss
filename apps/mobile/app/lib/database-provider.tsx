"use client";
import React, { PropsWithChildren, useEffect, useState } from "react";
import { Database } from "@nozbe/watermelondb";
import SQLiteAdapter from "@nozbe/watermelondb/adapters/sqlite";
import schema from "@dilocash/database/local/model/schema";
import { Intent } from "@dilocash/database/local/model/intent";
import { Command } from "@dilocash/database/local/model/commmand";
import { Transaction } from "@dilocash/database/local/model/transaction";
import { DatabaseProvider } from "@nozbe/watermelondb/react";
import { setGenerator } from "@nozbe/watermelondb/utils/common/randomId";
import "react-native-get-random-values";
import { v4 as uuidv4 } from "uuid";

const CustomeDatabaseProvider = ({ children }: PropsWithChildren) => {
  const [database, setDatabase] = useState<Database | null>(null);

  useEffect(() => {
    (async () => {
      setGenerator(() => uuidv4());

      const adapter = new SQLiteAdapter({
        dbName: "dilocash",
        schema,
        //migrations: [],
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
