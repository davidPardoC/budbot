import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import { Plus } from "lucide-react";
import { FaTelegramPlane } from "react-icons/fa";

export const AddNewTrandactionModal = () => {
  return (
    <Dialog>
      <DialogTrigger>
        <Button className="rounded-full">
          Add New <Plus size={20} className="ml-1" />{" "}
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Comming Soon</DialogTitle>
          <DialogDescription>
            <p>This feature is not available yet, but it will be soon.</p>
            <p>Register transacations in the telegram bot.</p>
            <a
              href={`https://t.me/${
                import.meta.env.PROD ? "PocketBudBot" : "DevPocketBud_bot"
              }`}
              className="flex items-center w-fit gap-1 text-white bg-[#24A1DE] px-4 py-3 rounded-full mt-3 text-md mx-auto"
            >
              Go to bot <FaTelegramPlane size={20} />
            </a>
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};
