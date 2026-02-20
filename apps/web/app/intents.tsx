"use client";
import { withObservables } from "@nozbe/watermelondb/react";
import { Intent } from "@dilocash/database/local/model/intent";
import { Button, ButtonText } from "@dilocash/ui/components/ui/button";

const IntentsList = ({ intents }: { intents: Intent[] }) => (
  <div>
    {intents.map((intent, i) => (
      <Button
        key={i}
        onPress={() => intent.delete()}
      ><ButtonText>{intent.textMessage}</ButtonText>
      </Button>
    ))}
  </div>
);

export default withObservables(["intents"], ({ intents }) => ({
  intents,
}))(IntentsList);