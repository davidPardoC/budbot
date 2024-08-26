import {
  ChartConfig,
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
import { Card } from "../../ui/card";
import { ChartSelection } from "./ChartSelection";

type Props = {
  chartConfig: ChartConfig;
  chartData: Record<string, string | number>[];
};

export const DashboardBarChart = ({ chartConfig, chartData }: Props) => {
  return (
    <Card className="w-full md:w-2/3 p-2">
      <div className="flex">
        <ChartSelection  className="ml-auto"/>
      </div>
      <ChartContainer config={chartConfig} className="min-h-[200px] ">
        <BarChart accessibilityLayer data={chartData}>
          <CartesianGrid vertical={false} />
          <XAxis
            dataKey="month"
            tickLine={false}
            tickMargin={10}
            axisLine={false}
            tickFormatter={(value) => value.slice(0, 3)}
          />
          <ChartTooltip content={<ChartTooltipContent />} />
          <ChartLegend content={<ChartLegendContent />} />
          <Bar dataKey="desktop" fill="var(--color-desktop)" radius={4} />
          <Bar dataKey="mobile" fill="var(--color-mobile)" radius={4} />
        </BarChart>
      </ChartContainer>
    </Card>
  );
};
