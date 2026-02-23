"use client";
import { withObservables } from "@nozbe/watermelondb/react";
import { Intent } from "@dilocash/database/local/model/intent";
import { Box } from "../ui/box";
import { Text } from "../ui/text";
import {
  AlertCircleIcon,
  ClockIcon,
  CloseCircleIcon,
  Icon,
  RemoveIcon,
} from "../ui/icon";
import { HStack } from "../ui/hstack";

const IntentsList = ({ intents }: { intents: Intent[] }) => (
  <Box className="p-4 w-full">
    {intents.map((intent) => (
      <HStack key={intent.id}>
        {intent.intentStatus === 0 && (
          <Icon className="text-gray-200" as={ClockIcon} />
        )}
        {intent.intentStatus === 4 && (
          <Icon className="text-red-500" as={AlertCircleIcon} />
        )}
        <Text className="px-2">{intent.textMessage}</Text>
      </HStack>
    ))}
  </Box>
);

export default withObservables(["intents"], ({ intents }) => ({
  intents,
}))(IntentsList);
