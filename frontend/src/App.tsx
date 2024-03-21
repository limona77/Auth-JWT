import "./App.css";

import { MantineProvider } from "@mantine/core";

import AuthForm from "./components/AuthForm.tsx";
import SwitchTheme from "./components/SwitchTheme.tsx";

import "@mantine/core/styles.css";

function App() {
  return (
    <MantineProvider defaultColorScheme="dark">
      <SwitchTheme />
      <AuthForm />
    </MantineProvider>
  );
}

export default App;
