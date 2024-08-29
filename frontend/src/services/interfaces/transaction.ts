export interface Transaction {
  amount: number;
  created_at: string;
  created_by: number;
  description: string;
  id: number;
  type: TransactionsTypes;
}

export interface TransactionGrouped {
  amount: number;
  description: string;
  type: TransactionsTypes;
}

export type TransactionsTypes = "income" | "expense";
