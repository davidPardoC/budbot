import { TRANSACTIONS_CACHE_KEY } from "@/constants/cache";
import { Transaction as TransactionInterface } from "@/services/interfaces/transaction";
import { getUserTransactions } from "@/services/transactions-services";
import { useDashboardStore } from "@/stores/auth.store";
import { getUserFromToken } from "@/utils/auth";
import { DateTime } from "luxon";
import { useQuery } from "react-query";
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Skeleton } from "../ui/skeleton";
import "./TransactionsHistory.css";




export const TransactionsHistory = () => {
  const { user_id } = getUserFromToken();
  const {currentMonth, currentYear} = useDashboardStore()
  const { data, isLoading } = useQuery([TRANSACTIONS_CACHE_KEY, currentMonth, currentYear], () => getUserTransactions(user_id, parseInt(currentMonth), parseInt(currentYear)));

  return (
    <Card className="w-full md:w-1/3 p-2 mt-2 md:mt-0 ">
      <CardHeader>
        <CardTitle className="text-lg">Recent Transactions</CardTitle>
      </CardHeader>
      <CardContent className="h-[40vh] overflow-y-scroll">
        {data && data.map((transaction, index) => (
          <Transaction key={index} {...transaction} />
        ))}
        {isLoading && (
          <>
            <TransactionItemSkeleton />
            <TransactionItemSkeleton />
            <TransactionItemSkeleton />
            <TransactionItemSkeleton />
            <TransactionItemSkeleton />
          </>
        )}
      </CardContent>
    </Card>
  );
};

const Transaction = ({ type, created_at, amount , description}: TransactionInterface) => {
  const isExpense = type === "expense";
  const color = isExpense ? "text-red-500" : "text-green-500";

  const sign = isExpense ? "-" : "+";

  return (
    <div className="flex justify-between items-center py-2 border-b border-gray-200">
      <div>
        <p className="text-sm capitalize">{description}</p>
        <p className="text-xs text-gray-500">{DateTime.fromISO(created_at).toLocaleString(DateTime.DATETIME_SHORT)}</p>
      </div>
      <p className={`text-sm ${color}`}>
        {sign}${amount}
      </p>
    </div>
  );
};

const TransactionItemSkeleton = () => {
  return (
    <div className="flex justify-between items-center py-2 border-b border-gray-200">
       <div>
        <Skeleton className="w-20 h-4 mb-1" />
        <Skeleton className="w-10 h-3" />
      </div>
      <Skeleton className="w-10 h-3" />
    </div>
  )
}