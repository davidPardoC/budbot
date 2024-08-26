import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { MONTHS } from "@/constants/months";
import { DateTime } from "luxon";

type Props = { defaultMonth?: number };

export const MonthSelector = ({
  defaultMonth = DateTime.now().month - 1,
}: Props) => {
  return (
    <Select defaultValue={defaultMonth.toString()}>
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Select a month" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {MONTHS.map((month, idx) => (
            <SelectItem key={month} value={idx.toString()}>
              {month}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
