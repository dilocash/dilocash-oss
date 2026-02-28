"use client";
import { Transaction } from "@dilocash/database/local/model/transaction";
import { Box } from "../ui/box";
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
import {
  Table,
  TableBody,
  TableHeader,
  TableRow,
  TableHead,
  TableData,
  TableFooter,
} from "../ui/table";
import { useObservable } from "../../hooks/useQuery";
import { Observable } from "@nozbe/watermelondb/utils/rx";

const TransactionsList = ({
  transactions: transactionsObservable,
  className,
}: {
  transactions: Observable<Transaction[]>;
  className?: string;
}) => {
  const transactions = useObservable(transactionsObservable);

  return (
    transactions.length > 0 && (
      <Box className={`w-full ${className}`}>
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
                      <AccordionTitleText>
                        Transaction Details
                      </AccordionTitleText>
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
    )
  );
};

export default TransactionsList;

