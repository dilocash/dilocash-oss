import { HStack } from "../ui/hstack";
import { Button, ButtonIcon, ButtonSpinner, ButtonText } from "../ui/button";
import useSync from "../../hooks/useSync";
import { Transport } from "@connectrpc/connect";
import { Input, InputField } from "../ui/input";
import { useDatabase } from "@nozbe/watermelondb/react";
import { useTranslation } from "react-i18next";
import { useState } from "react";
import { withDatabase } from "@nozbe/watermelondb/react";
import { Database } from "@nozbe/watermelondb";
import { Intent } from "@dilocash/database/local/model/intent";
import { Command } from "@dilocash/database/local/model/commmand";
import { Transaction } from "@dilocash/database/local/model/transaction";
import {
  AddIcon,
  CheckCircleIcon,
  CircleIcon,
  ClockIcon,
  CloseCircleIcon,
  LoaderIcon,
  MessageCircleIcon,
  RepeatIcon,
} from "../ui/icon";

const CommandsBar = ({ transport }: { transport: Transport }) => {
  const [commandText, setCommandText] = useState("");
  const { t } = useTranslation();
  const { sync, isSyncing } = useSync(transport);

  const AddCommandButton = withDatabase(
    ({ database }: { database: Database }) => {
      const handleClick = async () => {
        await database.write(async () => {
          let newCommand = database
            .get<Command>(Command.table)
            .prepareCreate((command) => {
              command.commandStatus = 0;
            });

          let newIntent = database
            .get<Intent>(Intent.table)
            .prepareCreate((intent) => {
              intent.textMessage = commandText;
              intent.intentStatus = 0;
              intent.command.set(newCommand);
            });

          let newTransaction = database
            .get<Transaction>(Transaction.table)
            .prepareCreate((transaction) => {
              const pieces = commandText.split(" ");

              transaction.amount = pieces[0];
              transaction.currency = pieces[1];
              transaction.description = pieces.slice(2).join(" ");
              transaction.command.set(newCommand);
            });

          await database.batch(newCommand, newIntent, newTransaction);
        });
        setCommandText("");
        await sync();
      };

      return (
        <Button onPress={handleClick}>
          <ButtonIcon size="md" as={AddIcon} />
        </Button>
      );
    },
  );

  return (
    <HStack
      space="md"
      className="fixed bottom-0 left-0 right-0 z-50 bg-white shadow-lg p-2"
    >
      <Input variant="rounded" className="grow">
        <InputField
          value={commandText}
          onChangeText={(text) => setCommandText(text)}
          type="text"
          placeholder={t("commands.command_placeholder")}
        />
      </Input>
      <AddCommandButton />
      <Button>
        <ButtonIcon size="md" as={MessageCircleIcon} />
      </Button>
      <Button onPress={sync}>
        {isSyncing ? (
          <ButtonSpinner color="orange" />
        ) : (
          <ButtonIcon size="md" as={RepeatIcon} />
        )}
      </Button>
    </HStack>
  );
};

export default CommandsBar;
