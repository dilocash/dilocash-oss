"use client";

import { Transport } from "@connectrpc/connect";
import CommandsBar from "./commands-bar";
import CommandsListView from "./commands-list";
import { VStack } from "../ui/vstack";
import { isWeb } from "@gluestack-ui/utils/nativewind-utils";
import { Box } from "../ui/box";
import { HStack } from "../ui/hstack";
import { AlertCircleIcon, CheckIcon, Icon } from "../ui/icon";
import { Text } from "../ui/text";
import { useAuth } from "../../auth/provider";

const CommandsView = ({ transport, className }: { transport: Transport, className?: string }) => {
  const { session } = useAuth();
  return (
    <VStack className={`${className}`}>
      <HStack className="p-5">
        <Box className="grow"> session: {session?.user?.email}</Box>
        {session ?
          <Box><Icon color="green" as={CheckIcon} /></Box> :
          <Box><Text className="text-gray-500">Offline</Text><Icon color="red" as={AlertCircleIcon} /></Box>}
      </HStack>
      <CommandsListView className="flex-1" />
      <CommandsBar transport={transport} className={`px-2 pt-2 ${isWeb ? "m-2" : ""}`} />
    </VStack>
  );
};

export default CommandsView;
