import { DateTime } from "luxon";
import { create } from "zustand";

interface DashboardStore {
  currentMonth: string;
  currentYear: string;
  setCurrentDate: (month: string, year: string) => void;
}

export const useAuthStore = create<DashboardStore>()((set) => ({
  currentMonth: DateTime.now().month.toString(),
  currentYear: DateTime.now().year.toString(),
  setCurrentDate: (month: string, year: string) =>
    set({ currentMonth: month, currentYear: year }),
}));
