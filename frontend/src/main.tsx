import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ThemeProvider } from "./components/common/ThemeProvider.tsx";
import "./index.css";
import MainRouter from "./router/MainRouter.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <MainRouter />
    </ThemeProvider>
  </StrictMode>
);
