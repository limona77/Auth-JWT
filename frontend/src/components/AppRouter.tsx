import { Route, Routes } from "react-router-dom";

import { useSelector } from "react-redux";

import { routes } from "../routes";
import ErrorAuthPage from "../pages/ErrorAuthPage.tsx";
import AuthPage from "../pages/AuthPage.tsx";

import { selectIsAuth } from "../store/slices/auth/selectors.ts";

const AppRouter = () => {
  const { isAuth } = useSelector(selectIsAuth);

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
