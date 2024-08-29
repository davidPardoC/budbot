import { CONFIG } from "@/constants/config";
import axios from "axios";
import { Transaction, TransactionsTypes } from "./interfaces/transaction";


export const getUserTransactions = async (userId: number, month:number, year:number) => {
    const { data } = await axios<Transaction[]>(
        `${CONFIG.baseApiUrl}/api/v1/users/${userId}/transactions?month=${month}&year=${year}`
      );
      return data;
}

export const getTransactionsGroupedByCategory = async (userId: number, month:number, year:number, filter: TransactionsTypes ) => {
    const { data = [] } = await axios<Transaction[]>(
        `${CONFIG.baseApiUrl}/api/v1/users/${userId}/transactions-grouped?month=${month}&year=${year}`
      );
      return data.filter((transaction) => transaction.type === filter);
}