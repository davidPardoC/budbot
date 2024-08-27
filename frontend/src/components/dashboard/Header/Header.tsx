import { ModeToggle } from "@/components/common/ModeToggle";
import { MonthSelector } from "@/components/common/MonthSelector";
import { YearSelector } from "@/components/common/YearSelector";
import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/stores/auth.store";
import { Settings } from "lucide-react";
import { Link } from "wouter";

export const Header = () => {
  const { currentMonth, currentYear } = useAuthStore();

  return (
    <header className="grid grid-cols-1 md:grid-cols-2 ">
      <h1 className="text-2xl font-bold flex items-center">
        Dashboard{" "}
        <Link href="/settings">
          <Button variant="outline" size="icon" className="ml-2">
            <Settings className="h-4 w-4" />
          </Button>
        </Link>
      </h1>
      <div className="mt-1 md:ml-auto flex gap-4">
        <YearSelector defaultYear={currentYear} />
        <MonthSelector defaultMonth={currentMonth} />
        <span className="ml-auto">
          <ModeToggle />
        </span>
      </div>
    </header>
  );
};
