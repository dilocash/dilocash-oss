"use client";

import { withObservables } from "@nozbe/watermelondb/react";
import {
  useGetCommands,
  useGetIntents,
  useGetTransactions,
} from "../../hooks/useQuery";
import IntentsList from "./intents-list";
import TransactionsList from "./transactions-list";
import { Command } from "@dilocash/database/local/model/commmand";
import { VStack } from "../ui/vstack";
import { Button, ButtonText } from "../ui/button";
import { Box } from "../ui/box";
import { HStack } from "../ui/hstack";
import { Text } from "../ui/text";
import {
  Accordion,
  AccordionContent,
  AccordionContentText,
  AccordionHeader,
  AccordionIcon,
  AccordionItem,
  AccordionTitleText,
  AccordionTrigger,
} from "../ui/accordion";
import { ChevronDownIcon, ChevronUpIcon, Icon, TrashIcon } from "../ui/icon";

const CommandsListView = () => {
  const commands = useGetCommands();
  return <EnhancedCommandsList commands={commands}></EnhancedCommandsList>;
};

export default CommandsListView;

const CommandsList = ({ commands }: { commands: Command[] }) => (
  <VStack>
    {commands.map((command, i) => (
      <CommandItem key={i} command={command} />
    ))}
  </VStack>
);

const EnhancedCommandsList = withObservables(["commands"], ({ commands }) => ({
  commands,
}))(CommandsList);

const CommandItem = ({ command }: { command: Command }) => {
  const intents = useGetIntents(command.id); // we need to use this hook for observability to work.
  const transactions = useGetTransactions(command.id); // we need to use this hook for observability to work.

  return (
    <Box className="p-4 m-2 border border-gray-200 rounded-lg border-b-4 bg-white">
      <HStack className="flex w-full items-center">
        <IntentsList intents={intents} className="grow" />
        <Button className="m-2" onPress={() => command.delete()}>
          <Icon as={TrashIcon} className="w-4 h-4 text-typography-100" />
        </Button>
      </HStack>
      <TransactionsList className="w-full" transactions={transactions} />
    </Box>
  );
};
