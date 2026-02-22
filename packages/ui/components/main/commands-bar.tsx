import { HStack } from "../ui/hstack";
import { Button, ButtonText } from "../ui/button";
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

const CommandsBar = ({ transport }: { transport: Transport }) => {
  const [commandText, setCommandText] = useState("");
  const { t } = useTranslation();
  const { sync, isSyncing } = useSync(transport);

  const AddCommandButton = withDatabase(
    ({ database }: { database: Database }) => {
      const handleClick = async () => {
        console.log("before write");
        await database.write(async () => {
          console.log("inside write");
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
          console.log("after batch");
        });
        console.log("after write");
        setCommandText("");
        console.log("before sync");
        await sync();
        console.log("after sync");
      };

      return (
        <Button className="w-32" onPress={handleClick}>
          <ButtonText>{t("commands.add")}</ButtonText>
        </Button>
      );
    },
  );

  return (
    <>
      <HStack className="fixed bottom-0 left-0 right-0 z-50 bg-white shadow-lg">
        <Input className="grow">
          <InputField
            value={commandText}
            onChangeText={(text) => setCommandText(text)}
            type="text"
            placeholder={t("commands.command_placeholder")}
          />
        </Input>
        <AddCommandButton />
        <Button className="w-32">
          <ButtonText>{t("commands.record")}</ButtonText>
        </Button>
        <Button className="w-32" onPress={sync}>
          <ButtonText>
            {isSyncing ? t("commands.syncing") : t("commands.sync")}
          </ButtonText>
        </Button>
      </HStack>
    </>
  );
};

export default CommandsBar;
