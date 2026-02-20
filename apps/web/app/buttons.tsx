"use client";
import { useDatabase } from "@nozbe/watermelondb/react";
import { Command } from "@dilocash/database/local/model/commmand";
import { Intent } from "@dilocash/database/local/model/intent";
import { Transaction } from "@dilocash/database/local/model/transaction";
import { withDatabase } from "@nozbe/watermelondb/DatabaseProvider";
import { Database } from "@nozbe/watermelondb";
import useSync from "./hooks/useSync";
import { Button, ButtonText } from "@dilocash/ui/components/ui/button";

// Method 1 to access the database:
const CreateCommandButton = () => {
  const database = useDatabase();
  const handleClick = async () => {
    await database.write(async () => {
      let newCommand = await database.get<Command>(Command.table).create((command) => {
        command.status = "pending";
      });
      await database.get<Intent>(Intent.table).create((intent) => {
          intent.textMessage = "A intent";
          intent.status = "pending";
          intent.command.set(newCommand);
        });
      await database.get<Transaction>(Transaction.table).create((transaction) => {
          transaction.amount = "100";
          transaction.currency = "USD";
          transaction.description = "A transaction";
          transaction.command.set(newCommand);
        });
    });
  };

  return (
    <Button onPress={handleClick}>
      <ButtonText>Create a command</ButtonText>
    </Button>
  );
};

// Method 2 to access the database:
const CreateIntentButton = withDatabase(
  ({ database, command }: { database: Database; command: Command }) => {
    const handleClick = async () => {
      await database.write(async () => {
        await database.get<Intent>(Intent.table).create((intent) => {
          intent.textMessage = "A intent";
          intent.status = "pending";
          intent.command.set(command);
        });
      });
    };

    return (
      <Button onPress={handleClick}>
        <ButtonText>Add an intent</ButtonText>
      </Button>
    );
  }
);

const SyncButton = () => {
  const { sync, isSyncing } = useSync();

  return (
    <Button onPress={sync}>
      <ButtonText>{isSyncing ? "Syncing..." : "Click to sync"}</ButtonText>
    </Button>
  );
};

export { CreateCommandButton, CreateIntentButton, SyncButton };