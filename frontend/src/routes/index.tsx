import DemoPage from "../pages/DemoPage.tsx";
import AuthPage from "../pages/AuthPage.tsx";

interface IRoute {
  path: string;
  component: React.ReactNode;
}
export const routes: IRoute[] = [
  {
    path: "/",
    component: <AuthPage />,
  },
  {
    path: "/demo",
    component: <DemoPage />,
  },
];
