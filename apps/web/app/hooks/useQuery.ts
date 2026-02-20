import { Observable } from "@nozbe/watermelondb/utils/rx";
import { useEffect, useState } from "react";
import { Command } from "@dilocash/database/local/model/commmand";
import { Intent } from "@dilocash/database/local/model/intent";
import { Transaction } from "@dilocash/database/local/model/transaction";
import { useDatabase } from "@nozbe/watermelondb/react";
import { Q } from "@nozbe/watermelondb";

const defaultObservable = <T>(): Observable<T[]> =>
  new Observable<T[]>((observer) => {
    observer.next([]);
  });

export const useGetCommands = () => {
  const database = useDatabase();
  const [commands, setCommands] = useState<Observable<Command[]>>(defaultObservable);

  useEffect(() => {
    setCommands(database.get<Command>(Command.table).query().observe());
  }, [database]);

  return commands;
};
export const useGetIntents = (commandId: string) => {
  const database = useDatabase();
  const [intents, setIntents] =
    useState<Observable<Intent[]>>(defaultObservable);

  useEffect(() => {
    setIntents(
      database
        .get<Intent>(Intent.table)
        .query(Q.where("command_id", commandId))
        .observe()
    );
  }, [database, commandId]);

  return intents;
};