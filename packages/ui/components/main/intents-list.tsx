"use client";
import { withObservables } from "@nozbe/watermelondb/react";
import { Intent } from "@dilocash/database/local/model/intent";
import { Box } from "../ui/box";

const IntentsList = ({ intents }: { intents: Intent[] }) => (
  <Box className="p-4 border border-orange-200 rounded-lg border-b-4 w-full">
    {intents.map((intent) => (
      <Box key={intent.id}>{intent.textMessage}</Box>
    ))}
  </Box>
);

export default withObservables(["intents"], ({ intents }) => ({
  intents,
}))(IntentsList);
