import "./App.css";

import { MantineProvider } from "@mantine/core";

import { BrowserRouter } from "react-router-dom";

import "@mantine/core/styles.css";

import AppRouter from "./components/AppRouter.tsx";

function App() {
  return (
    <MantineProvider>
      <BrowserRouter>
        <AppRouter />
      </BrowserRouter>
    </MantineProvider>
  );
}

export default App;
