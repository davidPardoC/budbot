import { ReactNode } from "react";
import { Card } from "../ui/card";

type Props = {
  title: string;
  amount: number;
  isMoney?: boolean;
  subtitle: string | null;
  icon?: ReactNode;
};

export const StatCard = ({
  amount,
  isMoney = false,
  title,
  icon = null,
  subtitle = null,
}: Props) => {
  return (
    <Card className="w-full p-5">
      <div className="flex">
        {title}
        {icon}
      </div>
      <div className="flex justify-between">
        <p className="font-black">
          {isMoney ? "$" : ""}
          {amount}
        </p>
      </div>
      <p className="text-xs">{subtitle}</p>
    </Card>
  );
};
