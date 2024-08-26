import { YEARS } from "@/constants/years";
import { DateTime } from "luxon";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "../ui/select";

type Props = { defaultYear?: string };

export const YearSelector = ({defaultYear = DateTime.now().year.toString()}:Props) => {
    console.log(defaultYear)
  return (
    <Select defaultValue={defaultYear}>
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Select a month" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {YEARS.map((year) => (
            <SelectItem key={year} value={year}>
              {year}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  )
}
