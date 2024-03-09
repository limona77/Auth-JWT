import "./App.css";

import { MantineProvider } from "@mantine/core";

import AuthForm from "./components/AuthForm";
import "@mantine/core/styles.css";

function App() {
  return (
    <MantineProvider>
      <AuthForm />
    </MantineProvider>
  );
}

export default App;
