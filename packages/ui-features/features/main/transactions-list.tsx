/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

"use client";
import { Transaction } from "@dilocash/database/local/model/transaction";
import { Box } from "@dilocash/ui-components/components/ui//box";
import {
  Accordion,
  AccordionItem,
  AccordionHeader,
  AccordionTrigger,
  AccordionTitleText,
  AccordionIcon,
  AccordionContent,
  AccordionContentText,
} from "@dilocash/ui-components/components/ui//accordion";
import { ChevronDownIcon, ChevronUpIcon } from "@dilocash/ui-components/components/ui//icon";
import {
  Table,
  TableBody,
  TableHeader,
  TableRow,
  TableHead,
  TableData,
  TableFooter,
} from "@dilocash/ui-components/components/ui//table";
import { useObservable } from "../../hooks/useQuery";
import { Observable } from "@nozbe/watermelondb/utils/rx";
import { useTranslation } from "react-i18next";

const TransactionsList = ({
  transactions: transactionsObservable,
  className,
}: {
  transactions: Observable<Transaction[]>;
  className?: string;
}) => {
  const transactions = useObservable(transactionsObservable);
  const { t } = useTranslation();

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
                        {t('transactions.table_title')}
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
                      <TableHead className="text-xs md:text-base p-1">{t('transactions.description')}</TableHead>
                      <TableHead className="text-xs md:text-base p-1">{t('transactions.currency')}</TableHead>
                      <TableHead className="text-xs md:text-base p-1">{t('transactions.amount')}</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {transactions.map((transaction) => (
                      <TableRow key={transaction.id}>
                        <TableData className="text-xs md:text-base">{transaction.description}</TableData>
                        <TableData className="text-xs md:text-base">{transaction.currency}</TableData>
                        <TableData className="text-xs md:text-base">{transaction.amount}</TableData>
                      </TableRow>
                    ))}
                  </TableBody>
                  <TableFooter>
                    <TableRow>
                      <TableHead className="text-xs md:text-base">{t('transactions.total')}</TableHead>
                      <TableHead className="text-xs md:text-base"></TableHead>
                      <TableHead className="text-xs md:text-base">
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

