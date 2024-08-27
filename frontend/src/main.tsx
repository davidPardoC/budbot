import axios from "axios";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { QueryClient, QueryClientProvider } from "react-query";
import { ThemeProvider } from "./components/common/ThemeProvider.tsx";
import { CONFIG } from "./constants/config.ts";
import "./index.css";
import MainRouter from "./router/MainRouter.tsx";

const queryClient = new QueryClient();

axios.defaults.baseURL = CONFIG.baseApiUrl

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <QueryClientProvider client={queryClient}>
        <MainRouter />
      </QueryClientProvider>
    </ThemeProvider>
  </StrictMode>
);
