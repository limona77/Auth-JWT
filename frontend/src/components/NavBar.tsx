import { Box, Button } from "@mantine/core";

import { SwitchThemeContext } from "../context";

import { SwitchTheme } from "./SwitchTheme.tsx";

export interface ISwitchThemeProps {
  size: string;
  variant: string;
}
const NavBar = () => {
  const switchThemeProps: ISwitchThemeProps = { size: "xl", variant: "light" };
  return (
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
      <SwitchThemeContext.Provider value={switchThemeProps}>
        <SwitchTheme />
      </SwitchThemeContext.Provider>
      <Button variant="filled" color="rgba(242, 41, 41, 1)">
        Выйти
      </Button>
    </Box>
  );
};

export default NavBar;
