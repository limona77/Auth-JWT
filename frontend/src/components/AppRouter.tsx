import { useContext } from "react";

import { Route, Routes } from "react-router-dom";

import { IsAuthContext } from "../context";
import { routes } from "../routes";
import ErrorAuthPage from "../pages/ErrorAuthPage.tsx";
import AuthPage from "../pages/AuthPage.tsx";

const AppRouter = () => {
  const isAuth = useContext(IsAuthContext);
  return isAuth ? (
    <Routes>
      {routes.map((route) => (
        <Route key={route.path} path={route.path} element={route.component} />
      ))}
    </Routes>
  ) : (
    <Routes>
      <Route element={<ErrorAuthPage />} path={"/demo"} />
      <Route element={<AuthPage />} path={"/"} />
    </Routes>
  );
};

export default AppRouter;
