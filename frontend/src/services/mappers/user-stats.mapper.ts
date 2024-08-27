import { StatsResponseInterface } from "../interfaces/stats-response.interface";

export const mappUserStats = (data: StatsResponseInterface) => {
  const mappedData = Object.entries(data).map(([key, value]) => {
    return {
      title: key.replace(/_/g, " ").toUpperCase(),
      amount: value.toFixed(2),
      subtitle: `${key.replace(/_/g, " ").toUpperCase()} in current mount`,
      isMoney: true,
    };
  });

  return mappedData
};
