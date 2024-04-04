import { Box, Button } from "@mantine/core";

import { useSelector } from "react-redux";

import { useNavigate } from "react-router-dom";

import { selectIsAuth } from "../store/slices/auth/selectors.ts";

import { useAppDispatch } from "../store";

import { fetchLogout } from "../store/slices/auth/asyncActions.ts";

import { SwitchTheme } from "./SwitchTheme.tsx";
import MyLoader from "./MyLoader.tsx";

const NavBar = () => {
  const { isAuth, isLoading } = useSelector(selectIsAuth);
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const handleLogoutEvent = async () => {
    const res = await dispatch(fetchLogout());
    if (res.meta.requestStatus == "fulfilled") {
      navigate("/");
    }
  };
  return isLoading ? (
    <MyLoader />
  ) : isAuth ? (
    <Box
      style={{
        display: "flex",
        right: 0,
        position: "absolute",
        padding: "10px",
        gap: "10px",
        alignItems: "center",
      }}
    >
      <SwitchTheme size={"xl"} variant={"light"} />
      <Button
        onClick={handleLogoutEvent}
        variant="filled"
        color="rgba(242, 41, 41, 1)"
      >
        Выйти
      </Button>
    </Box>
  ) : (
    <Box
      style={{
        display: "flex",
        right: 0,
        position: "absolute",
        padding: "10px",
        gap: "10px",
        alignItems: "center",
      }}
    >
      <SwitchTheme size={"xl"} variant={"light"} />
    </Box>
  );
};
export default NavBar;
