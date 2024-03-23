import "./App.css";

import { MantineProvider } from "@mantine/core";

import { createBrowserRouter, RouterProvider } from "react-router-dom";

import AuthForm from "./components/AuthForm.tsx";

import "@mantine/core/styles.css";

import NavBar from "./components/NavBar.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: [<NavBar />, <AuthForm />],
  },
  {
    path: "/demo",
    element: <NavBar />,
  },
]);
function App() {
  return (
    <MantineProvider>
      <RouterProvider router={router} />
    </MantineProvider>
  );
}

export default App;
