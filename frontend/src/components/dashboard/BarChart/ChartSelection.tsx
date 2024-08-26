
  
  import {
    ToggleGroup,
    ToggleGroupItem,
} from "@/components/ui/toggle-group"
import { ChartNoAxesColumnIncreasing, ChartPie } from "lucide-react"
import { HTMLProps } from "react"

 type Props = HTMLProps<HTMLDivElement>

  export function ChartSelection(props: Props) {
    return (
      <ToggleGroup  type={"multiple"} variant="outline" className={props.className} >
        <ToggleGroupItem value="bold" aria-label="Toggle bold">
          <ChartNoAxesColumnIncreasing className="h-4 w-4" />
        </ToggleGroupItem>
        <ToggleGroupItem value="italic" aria-label="Toggle italic">
          <ChartPie className="h-4 w-4" />
        </ToggleGroupItem>
      </ToggleGroup>
    )
  }
  