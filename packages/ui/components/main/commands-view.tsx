"use client";

import { Transport } from "@connectrpc/connect";
import CommandsBar from "./commands-bar";
import CommandsListView from "./commands-list";
import { VStack } from "../ui/vstack";
import { isWeb } from "@gluestack-ui/utils/nativewind-utils";
import { Box } from "../ui/box";
import { HStack } from "../ui/hstack";
import { CheckIcon, Icon, InfoIcon } from "../ui/icon";
import { Text } from "../ui/text";
import { useAuth } from "../../auth/provider";
import { Alert, AlertIcon, AlertText } from "../ui/alert";
import { Pressable } from "../ui/pressable";
import { useRouter } from "solito/navigation";
import { useTranslation } from "react-i18next";

const CommandsView = ({ transport, className }: { transport: Transport, className?: string }) => {
  const { session } = useAuth()
  const { replace } = useRouter()
  const { t } = useTranslation();

  const goToLogin = () => {
    replace('/auth/signin')
  };

  return (
    <VStack className={`${className}`}>
      <HStack className="p-2">
        <Box className="grow">{session && <Text>session: {session?.user?.email}</Text>}</Box>
        {session ?
          <Box><Icon color="green" as={CheckIcon} /></Box> :
          <Alert action="warning" variant="outline" >
            <AlertIcon as={InfoIcon} />
            <Pressable onPress={goToLogin}><AlertText>{t('commands.connect')}</AlertText></Pressable>
          </Alert>
        }
      </HStack>
      <CommandsListView className="flex-1" />
      <CommandsBar transport={transport} className={`px-2 pt-2 ${isWeb ? "m-2" : ""}`} />
    </VStack>
  );
};

export default CommandsView;
