"use client";

import { useDatabase } from "@nozbe/watermelondb/react";
import {
  useGetCommands,
  useGetIntents,
  useGetTransactions,
  useObservable,
} from "../../hooks/useQuery";
import IntentsList from "./intents-list";
import TransactionsList from "./transactions-list";
import { Command } from "@dilocash/database/local/model/commmand";
import { Button, ButtonIcon } from "../ui/button";
import { HStack } from "../ui/hstack";
import { Card } from "../ui/card";
import { TrashIcon } from "../ui/icon";
import { useEffect, useRef } from "react";
import { Platform, ScrollView } from 'react-native';

const CommandsListView = ({ className }: { className?: string }) => {
  const commandsObservable = useGetCommands();
  const commands = useObservable(commandsObservable);

  // Create a ref for the bottom-most element
  const bottomOfPanelRef = useRef<HTMLDivElement>(null);
  const isWeb = Platform.OS === 'web';
  if (isWeb) {
    // Use useEffect to scroll to the ref whenever messages change
    useEffect(() => {
      // The scrollIntoView method handles the actual scrolling
      bottomOfPanelRef.current?.scrollIntoView({ behavior: "smooth" });
    }, [commands]); // Dependency array includes messages to trigger on update
  }

  return (
    <ScrollView
      style={{ height: "100%" }}
      contentContainerStyle={{ flexGrow: 1 }}
    >
      {commands.map((command, i) => (
        <CommandItem key={i} command={command} />
      ))}
      {isWeb && <div ref={bottomOfPanelRef} />}
    </ScrollView>
  );
};

export default CommandsListView;

const CommandItem = ({ command }: { command: Command }) => {
  const intents = useGetIntents(command.id); // we need to use this hook for observability to work.
  const transactions = useGetTransactions(command.id); // we need to use this hook for observability to work.

  const database = useDatabase();
  const handleDelete = async () => {
    console.debug("Deleting command", command.id);
    await database.write(async () => {
      await command.markAsDeleted();
      (await command.intents).forEach((intent) => intent.markAsDeleted());
      (await command.transactions).forEach((transaction) =>
        transaction.markAsDeleted(),
      );
    });
  };

  return (
    <Card
      className="p-4 m-2 border-b-4 bg-white border-b-orange-300"
      variant="outline"
    >
      <HStack className="flex w-full items-center">
        <IntentsList intents={intents} />
        <Button onPress={handleDelete}>
          <ButtonIcon as={TrashIcon} className="w-4 h-4text-typography-100" />
        </Button>
      </HStack>
      <TransactionsList className="w-full" transactions={transactions} />
    </Card>
  );
};
