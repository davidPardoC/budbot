import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group";
import { TRANSACTION_GROUPED_CACHE_KEY } from "@/constants/cache";
import { TransactionsTypes } from "@/services/interfaces/transaction";
import { getTransactionsGroupedByCategory } from "@/services/transactions-services";
import { useDashboardStore } from "@/stores/auth.store";
import { getUserFromToken } from "@/utils/auth";
import { useState } from "react";
import { useQuery } from "react-query";
import { Cell, Pie, PieChart, ResponsiveContainer, Tooltip } from "recharts";
import { Card } from "../../ui/card";

const COLORS = [
  "#0088FE", // Azul brillante
  "#00C49F", // Verde
  "#FFBB28", // Amarillo
  "#FF8042", // Naranja
  "#8884D8", // Púrpura suave
  "#82ca9d", // Verde menta
  "#A4DE6C", // Verde claro
  "#D0ED57", // Amarillo lima
  "#FF4560", // Rojo vibrante
  "#546E7A",
];

const RADIAN = Math.PI / 180;
const renderCustomizedLabel = (props: Record<string, number & string>) => {
  const { cx, cy, midAngle, innerRadius, outerRadius, percent, fill } = props;
  const radius = 25 + innerRadius + (outerRadius - innerRadius);
  const x = cx + radius * Math.cos(-midAngle * RADIAN);
  const y = cy + radius * Math.sin(-midAngle * RADIAN);

  return (
    <text
      x={x}
      y={y}
      fill={fill}
      textAnchor={x > cx ? "start" : "end"}
      dominantBaseline="central"
      fontSize={20}
    >
      {`${(percent * 100).toFixed(0)}%`}
    </text>
  );
};

export const DashboardChart = () => {
  const { user_id } = getUserFromToken();
  const { currentMonth, currentYear } = useDashboardStore();
  const [filter, setFilter] = useState<TransactionsTypes>("expense");
  const { data = [] } = useQuery(
    [TRANSACTION_GROUPED_CACHE_KEY, currentMonth, currentYear, filter],
    () =>
      getTransactionsGroupedByCategory(
        user_id,
        parseInt(currentMonth),
        parseInt(currentYear),
        filter
      )
  );
  return (
    <Card className="w-full md:w-2/3 p-2">
      <ToggleGroup
        type="single"
        variant="outline"
        className="flex justify-start"
        value={filter}
      >
        <ToggleGroupItem
          value="expense"
          aria-label="Toggle italic"
          onClick={() => setFilter("expense")}
        >
          Expenses
        </ToggleGroupItem>
        <ToggleGroupItem
          value="income"
          aria-label="Toggle bold"
          onClick={() => setFilter("income")}
        >
          Incomes
        </ToggleGroupItem>
      </ToggleGroup>
      <div className="h-[25vh] md:h-[50vh] flex justify-center items-center">
        <ResponsiveContainer width={700} height="80%">
          <PieChart>
            <Tooltip contentStyle={{ textTransform: "capitalize" }} />
            <Pie
              data={data}
              dataKey="amount"
              nameKey={"description"}
              cx="50%"
              cy="50%"
              innerRadius={120}
              outerRadius={180}
              fill="#82ca9d"
              label={renderCustomizedLabel}
            >
              {data.map((_, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={COLORS[index % COLORS.length]}
                />
              ))}
            </Pie>
          </PieChart>
        </ResponsiveContainer>
      </div>
    </Card>
  );
};
