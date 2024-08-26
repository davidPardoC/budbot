import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";
import { format } from "date-fns";
import { CalendarIcon } from "lucide-react";
import { DateTime } from "luxon";
import * as React from "react";
import { DateRange } from "react-day-picker";

type Props = {
  onDateSelect: (e: DateRange | undefined) => void;
  defaultDate?: DateRange;
};

export function DateRengePicker({
  className,
  onDateSelect,
  defaultDate,
}: React.HTMLAttributes<HTMLDivElement> & Props) {
  const [date, setDate] = React.useState<DateRange | undefined>(
    defaultDate
      ? defaultDate
      : {
          from: DateTime.now().set({ month: DateTime.now().month, day: 1 }).toJSDate(),
          to: new Date(),
        }
  );

  return (
    <div className={cn("grid gap-2", className)}>
      <Popover>
        <PopoverTrigger asChild>
          <Button
            id="date"
            variant={"outline"}
            className={cn(
              "w-full justify-start text-left font-normal",
              !date && "text-muted-foreground"
            )}
          >
            <CalendarIcon className="mr-2 h-4 w-4" />
            {date?.from ? (
              date.to ? (
                <>
                  {format(date.from, "LLL dd, y")} -{" "}
                  {format(date.to, "LLL dd, y")}
                </>
              ) : (
                format(date.from, "LLL dd, y")
              )
            ) : (
              <span>Pick a date</span>
            )}
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-auto p-0" align="end">
          <Calendar
            initialFocus
            mode="range"
            defaultMonth={date?.from}
            selected={date}
            onSelect={(e) => {
              onDateSelect(e);
              setDate(e);
            }}
            numberOfMonths={2}
          />
        </PopoverContent>
      </Popover>
    </div>
  );
}
