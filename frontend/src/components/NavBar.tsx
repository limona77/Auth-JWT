import { Box } from "@mantine/core";

import { SwitchThemeContext } from "../context";

import { SwitchTheme } from "./SwitchTheme.tsx";
import Demo from "./Demo.tsx";

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
      <Demo />
      <SwitchThemeContext.Provider value={switchThemeProps}>
        <SwitchTheme />
      </SwitchThemeContext.Provider>
    </Box>
  );
};

export default NavBar;
