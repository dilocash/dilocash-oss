"use client";

import { Transport } from "@connectrpc/connect";
import { Box } from "../ui/box";
import CommandsBar from "./commands-bar";
import CommandsListView from "./commands-list";

const CommandsView = ({ transport }: { transport: Transport }) => {
  return (
    <Box className="flex flex-col h-full py-10 bg-gray-100">
      <CommandsListView />
      <CommandsBar transport={transport} />
    </Box>
  );
};

export default CommandsView;
