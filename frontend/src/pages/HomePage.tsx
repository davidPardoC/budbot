import { StatCard } from "@/components/common/StatCard";
import { DashboardBarChart } from "@/components/dashboard/BarChart/BarChart";
import { Header } from "@/components/dashboard/Header/Header";
import { TransactionsHistory } from "@/components/dashboard/TransactionsHistory";
import { type ChartConfig } from "@/components/ui/chart";
import { STATS_CACHE_KEY } from "@/constants/cache";
import { getUserStats } from "@/services/stas-services";
import { useDashboardStore } from "@/stores/auth.store";
import { getUserFromToken } from "@/utils/auth";
import { useQuery } from "react-query";

const chartData = [
  { month: "January", desktop: 186, mobile: 80 },
  { month: "February", desktop: 305, mobile: 200 },
  { month: "March", desktop: 237, mobile: 120 },
  { month: "April", desktop: 73, mobile: 190 },
  { month: "May", desktop: 209, mobile: 130 },
  { month: "June", desktop: 214, mobile: 140 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#2563eb",
  },
  mobile: {
    label: "Mobile",
    color: "#60a5fa",
  },
} satisfies ChartConfig;

export const HomePage = () => {
  const { currentMonth, currentYear } = useDashboardStore();

  const { user_id } = getUserFromToken();
  const { data, isLoading } = useQuery(
    [STATS_CACHE_KEY, currentMonth, currentYear],
    () => getUserStats(user_id, parseInt(currentMonth), parseInt(currentYear))
  );
  return (
    <main className="container pt-10 pb-10">
      <Header />
      <section className="grid grid-cols-2 md:grid-cols-4 gap-4 mt-2">
        {data &&
          data.map((card, index) => (
            <StatCard key={index} {...card} isLoading={isLoading} isMoney />
          ))}
      </section>
      <section className="flex flex-col md:flex-row mt-5 gap-4">
        <DashboardBarChart chartConfig={chartConfig} chartData={chartData} />
        <TransactionsHistory
          transactions={[
            { type: "expense", date: new Date().toDateString(), amount: 100 },
            { type: "income", date: new Date().toDateString(), amount: 100 },
            { type: "expense", date: new Date().toDateString(), amount: 100 },
            { type: "income", date: new Date().toDateString(), amount: 100 },
            { type: "expense", date: new Date().toDateString(), amount: 100 },
            { type: "income", date: new Date().toDateString(), amount: 100 },
            { type: "expense", date: new Date().toDateString(), amount: 100 },
            { type: "income", date: new Date().toDateString(), amount: 100 },
            { type: "expense", date: new Date().toDateString(), amount: 100 },
            { type: "income", date: new Date().toDateString(), amount: 100 },
          ]}
        />
      </section>
    </main>
  );
};
