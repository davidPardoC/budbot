import { ReactNode } from "react";
import { Card } from "../ui/card";
import { Skeleton } from "../ui/skeleton";

type Props = {
  title: string;
  amount: number;
  isMoney?: boolean;
  subtitle: string | null;
  icon?: ReactNode;
  isLoading?: boolean;
  type: "money" | "percentage";
};

const mapType = {
  money: "$",
  percentage: "%",
};

export const StatCard = ({
  amount,
  title,
  isLoading,
  type,
  icon = null,
  subtitle = null,
}: Props) => {
  return (
    <Card className="w-full p-3">
      <div className="flex">
        {isLoading ? <Skeleton className="h-4 w-full" /> : title}
        {isLoading ? (
          <Skeleton className="h-4" />
        ) : (
          <p className="ml-1">{icon}</p>
        )}
      </div>
      <div className="flex justify-between">
        <p className="font-black flex w-full mt-2 items-center gap-1">
          {mapType[type]}
          {isLoading ? <Skeleton className="h-4 w-full" /> : amount.toFixed(2)}
        </p>
      </div>
      <p className="text-xs mt-2">
        {isLoading ? <Skeleton className="h-4 w-full" /> : subtitle}
      </p>
    </Card>
  );
};
