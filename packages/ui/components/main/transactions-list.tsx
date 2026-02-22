"use client";
import { withObservables } from "@nozbe/watermelondb/react";
import { Transaction } from "@dilocash/database/local/model/transaction";
import { Button, ButtonText } from "../ui/button";
import { Box } from "../ui/box";
import { HStack } from "../ui/hstack";
import { Text } from "../ui/text";
import { VStack } from "../ui/vstack";
import {
  Accordion,
  AccordionItem,
  AccordionHeader,
  AccordionTrigger,
  AccordionTitleText,
  AccordionIcon,
  AccordionContent,
  AccordionContentText,
} from "../ui/accordion";
import { ChevronDownIcon, ChevronUpIcon } from "../ui/icon";
import { useState } from "react";
import {
  Table,
  TableBody,
  TableHeader,
  TableRow,
  TableHead,
  TableData,
  TableFooter,
} from "../ui/table";

const TransactionsList = ({
  transactions,
}: {
  transactions: Transaction[];
}) => {
  const [isExpanded, setIsExpanded] = useState(false);
  return (
    <Box className="w-full">
      <Accordion
        size="sm"
        variant="unfilled"
        type="single"
        isCollapsible={true}
        isDisabled={false}
        className="mt-4 border border-outline-200 border-b-2 rounded-lg"
      >
        <AccordionItem value="1">
          <AccordionHeader>
            <AccordionTrigger>
              {({ isExpanded }: { isExpanded: boolean }) => {
                return (
                  <>
                    <AccordionTitleText>Transaction Details</AccordionTitleText>
                    {isExpanded ? (
                      <AccordionIcon as={ChevronUpIcon} className="ml-3" />
                    ) : (
                      <AccordionIcon as={ChevronDownIcon} className="ml-3" />
                    )}
                  </>
                );
              }}
            </AccordionTrigger>
          </AccordionHeader>
          <AccordionContent>
            <AccordionContentText>
              <Table className="w-full">
                <TableHeader>
                  <TableRow>
                    <TableHead>Description</TableHead>
                    <TableHead>Currency</TableHead>
                    <TableHead>Amount</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {transactions.map((transaction) => (
                    <TableRow key={transaction.id}>
                      <TableData>{transaction.description}</TableData>
                      <TableData>{transaction.currency}</TableData>
                      <TableData>{transaction.amount}</TableData>
                    </TableRow>
                  ))}
                </TableBody>
                <TableFooter>
                  <TableRow>
                    <TableHead>Total</TableHead>
                    <TableHead></TableHead>
                    <TableHead>
                      {transactions.reduce(
                        (acc: number, transaction: Transaction) => {
                          return acc + parseFloat(transaction.amount);
                        },
                        0,
                      )}
                    </TableHead>
                  </TableRow>
                </TableFooter>
              </Table>
            </AccordionContentText>
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </Box>
  );
};

export default withObservables(["transactions"], ({ transactions }) => ({
  transactions,
}))(TransactionsList);
