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
import { Link } from "wouter";

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
            <Link href="https://t.me/PocketBudBot">
              <Button className="mt-2 rounded-full bg-[#24A1DE] text-white">
                Go to Bot <FaTelegramPlane size={20} className="ml-2" />
              </Button>
            </Link>
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};
