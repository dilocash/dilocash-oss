"use client";

import { Transport } from "@connectrpc/connect";
import CommandsBar from "./commands-bar";
import CommandsListView from "./commands-list";
import { VStack } from "../ui/vstack";
import { isWeb } from "@gluestack-ui/utils/nativewind-utils";
import { Box } from "../ui/box";
import { Session } from "@supabase/supabase-js";
import { HStack } from "../ui/hstack";
import { AlertCircleIcon, CheckIcon, Icon } from "../ui/icon";
import { Text } from "../ui/text";

const CommandsView = ({ transport, session, className }: { transport: Transport, session: Session | null, className?: string }) => {
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
