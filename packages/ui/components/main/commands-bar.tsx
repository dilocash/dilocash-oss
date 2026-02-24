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

import {
  createValidator,
  Violation,
  type ValidationResult,
} from "@bufbuild/protovalidate";
import { TransactionSchema } from "@dilocash/gen/ts/transport/dilocash/v1/transaction_types_pb";

import { create } from "@bufbuild/protobuf";
import { MicIcon } from "./icons/mic";

const CommandsBar = ({ transport }: { transport: Transport }) => {
  const validator = createValidator();
  const [commandText, setCommandText] = useState("");
  const { t } = useTranslation();
  const { sync, isSyncing } = useSync(transport);

  const AddCommandButton = withDatabase(
    ({ database }: { database: Database }) => {
      const handleClick = async () => {
        const pieces = commandText.split(" ");
        const parsedTransaction = {
          amount: pieces[0],
          currency: pieces[1],
          description: pieces.slice(2).join(" "),
        };

        let validTransaction = true;

        const validationResult: ValidationResult = validator.validate(
          TransactionSchema,
          create(TransactionSchema, parsedTransaction),
        );
        if (validationResult.kind !== "valid") {
          const cleanViolations = validationResult.violations?.filter(
            (violation: Violation) => {
              return violation.ruleId !== "string.uuid_empty";
            },
          );
          cleanViolations?.forEach((violation: Violation) => {
            console.error("Validation error", violation);
          });
          validTransaction = cleanViolations?.length === 0;
        }
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
              // 0 = unspecified, 4 = failed
              intent.intentStatus = validTransaction ? 0 : 4;
              intent.command.set(newCommand);
            });

          if (validTransaction) {
            let newTransaction = database
              .get<Transaction>(Transaction.table)
              .prepareCreate((transaction) => {
                transaction.amount = parsedTransaction.amount;
                transaction.currency = parsedTransaction.currency;
                transaction.description = parsedTransaction.description;
                transaction.command.set(newCommand);
              });
            await database.batch(newCommand, newIntent, newTransaction);
          } else {
            await database.batch(newCommand, newIntent);
          }
        });
        setCommandText("");
        await sync();
      };

      return (
        <Button size="md" onPress={handleClick}>
          <ButtonIcon as={AddIcon} />
        </Button>
      );
    },
  );

  return (
    <HStack
      space="sm"
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
      <Button size="md">
        <ButtonIcon as={MicIcon} />
      </Button>
      <Button size="md" onPress={sync}>
        {isSyncing ? (
          <ButtonSpinner color="orange" />
        ) : (
          <ButtonIcon as={RepeatIcon} />
        )}
      </Button>
    </HStack>
  );
};

export default CommandsBar;
