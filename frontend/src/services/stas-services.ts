import { CONFIG } from "@/constants/config";
import axios from "axios";
import { StatsResponseInterface } from "./interfaces/stats-response.interface";

export const getUserStats = async (userId: number, month:number, year:number) => {
  const { data } = await axios<StatsResponseInterface[]>(
    `${CONFIG.baseApiUrl}/api/v1/users/${userId}/stats?month=${month}&year=${year}`
  );

  return data;
};
