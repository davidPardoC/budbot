import { StatCard } from "@/components/common/StatCard";
import { DashboardChart } from "@/components/dashboard/BarChart/DashboardChart";
import { Header } from "@/components/dashboard/Header/Header";
import { TransactionsHistory } from "@/components/dashboard/TransactionsHistory";
import { STATS_CACHE_KEY } from "@/constants/cache";
import { getUserStats } from "@/services/stas-services";
import { useDashboardStore } from "@/stores/auth.store";
import { getUserFromToken } from "@/utils/auth";
import { useQuery } from "react-query";




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
        <DashboardChart />
        <TransactionsHistory />
      </section>
    </main>
  );
};
