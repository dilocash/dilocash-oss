"use client";
import { Intent } from "@dilocash/database/local/model/intent";
import { Box } from "../ui/box";
import { Text } from "../ui/text";
import {
  AlertCircleIcon,
  ClockIcon,
  Icon,
} from "../ui/icon";
import { HStack } from "../ui/hstack";
import { useObservable } from "../../hooks/useQuery";
import { Observable } from "@nozbe/watermelondb/utils/rx";

const IntentsList = ({ intents: intentsObservable, className }: { intents: Observable<Intent[]>, className?: string }) => {
  const intents = useObservable(intentsObservable);

  return (
    <Box className={`p-4 w-full ${className}`}>
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
};

export default IntentsList;
