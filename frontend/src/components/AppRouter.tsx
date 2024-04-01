import { Route, Routes } from "react-router-dom";

import { useSelector } from "react-redux";

import React, { useEffect } from "react";

import { routes } from "../routes";
import ErrorAuthPage from "../pages/ErrorAuthPage.tsx";
import AuthPage from "../pages/AuthPage.tsx";

import { selectIsAuth } from "../store/slices/auth/selectors.ts";
import { useAppDispatch } from "../store";
import { fetchAuthMe } from "../store/slices/auth/asyncActions.ts";
import NotFound from "../pages/NotFound.tsx";

import MyLoader from "./MyLoader.tsx";

const AppRouter = () => {
  const { isAuth, isLoading } = useSelector(selectIsAuth);
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(fetchAuthMe());
  }, []);
  return isLoading ? (
    <MyLoader />
  ) : isAuth ? (
    <Routes>
      {routes.map((route) => (
        <Route key={route.path} path={route.path} element={route.component} />
      ))}
    </Routes>
  ) : (
    <Routes>
      <Route element={<ErrorAuthPage />} path={"/demo"} />
      <Route element={<AuthPage />} path={"/"} />
      <Route element={<NotFound />} path={"*"} />
    </Routes>
  );
};

export default AppRouter;
