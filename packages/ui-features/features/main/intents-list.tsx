/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

"use client";
import { Intent } from "@dilocash/database/local/model/intent";
import { Box } from "@dilocash/ui-components/components/ui//box";
import { Text } from "@dilocash/ui-components/components/ui//text";
import {
  AlertCircleIcon,
  ClockIcon,
  Icon,
} from "@dilocash/ui-components/components/ui//icon";
import { HStack } from "@dilocash/ui-components/components/ui//hstack";
import { useObservable } from "../../hooks/useQuery";
import { Observable } from "@nozbe/watermelondb/utils/rx";

const IntentsList = ({ intents: intentsObservable, className }: { intents: Observable<Intent[]>, className?: string }) => {
  const intents = useObservable(intentsObservable);

  return (
    <Box className={`grow p-4 ${className}`}>
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
