import ReactDOM from "react-dom/client";

import { MantineProvider } from "@mantine/core";
import { BrowserRouter } from "react-router-dom";

import { Provider } from "react-redux";

import App from "./App.tsx";

import { store } from "./store";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <MantineProvider>
    <BrowserRouter>
      <Provider store={store}>
        <App />
      </Provider>
    </BrowserRouter>
  </MantineProvider>,
);
