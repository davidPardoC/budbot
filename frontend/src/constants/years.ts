export const generateYearsUntilNow = () => {
  const currentYear = new Date().getFullYear();
  const years = [];
  for (let i = 2024; i <= currentYear; i++) {
    years.push(i.toString());
  }
  return years;
};

export const YEARS = generateYearsUntilNow();
