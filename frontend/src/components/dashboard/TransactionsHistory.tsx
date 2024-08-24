import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import './TransactionsHistory.css';

type Props = {
  transactions: Transaction[];
};

type Transaction = {
    type: 'income' | 'expense';
    date: string;
    amount: number;
}

export const TransactionsHistory = ({ transactions }: Props) => {
  return (
    <Card className="w-full md:w-1/3 p-2 mt-2 md:mt-0 ">
      <CardHeader>
        <CardTitle className="text-lg">Recent Transactions</CardTitle>
      </CardHeader>
      <CardContent className="h-[40vh] overflow-y-scroll">
        {transactions.map((transaction, index) => (
          <Transaction key={index} {...transaction} />
        ))}
      </CardContent>
    </Card>
  );
};

const Transaction = ({type, date, amount}:Transaction) => {

    const isExpense = type === 'expense';
    const color = isExpense ? 'text-red-500' : 'text-green-500';

    const sign = isExpense ? '-' : '+';

  return (
    <div className="flex justify-between items-center py-2 border-b border-gray-200">
      <div>
        <p className="text-sm">Netflix</p>
        <p className="text-xs text-gray-500">{date}</p>
      </div>
      <p className={`text-sm ${color}`}>{sign}${amount}</p>
    </div>
  );
};
