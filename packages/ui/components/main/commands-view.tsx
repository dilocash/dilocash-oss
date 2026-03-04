"use client";

import { Transport } from "@connectrpc/connect";
import CommandsBar from "./commands-bar";
import CommandsListView from "./commands-list";
import { VStack } from "../ui/vstack";

const CommandsView = ({ transport, className }: { transport: Transport, className?: string }) => {
  return (
    <VStack className={`h-full ${className}`}>
      <CommandsListView className="flex-1" />
      <CommandsBar transport={transport} className="px-2 pt-2" />
    </VStack>
  );
};

export default CommandsView;
