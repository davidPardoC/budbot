import { StatCard } from "@/components/common/StatCard";
import { DashboardBarChart } from "@/components/dashboard/BarChart/BarChart";
import { Header } from "@/components/dashboard/Header/Header";
import { TransactionsHistory } from "@/components/dashboard/TransactionsHistory";
import { type ChartConfig } from "@/components/ui/chart";

const Cards = [
  {
    title: "Spent",
    amount: 1000,
    subtitle: "Spent in current mount",
    isMoney: true,
  },
  {
    title: "Income",
    amount: 1000,
    subtitle: "Spent in current mount",
    isMoney: true,
  },
  {
    title: "Budget",
    amount: 1000,
    subtitle: "Spent in current mount",
    isMoney: true,
  },
  {
    title: "Spent",
    amount: 1000,
    subtitle: "Spent in current mount",
    isMoney: true,
  },
];

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
  return (
    <main className="container pt-10">
      <Header/>
      <section className="grid grid-cols-2 md:grid-cols-4 gap-4 mt-2">
        {Cards.map((card, index) => (
          <StatCard key={index} {...card} />
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
