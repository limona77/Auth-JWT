import { Box, Button } from "@mantine/core";

import { SwitchTheme } from "./SwitchTheme.tsx";

const NavBar = () => {
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
      <SwitchTheme size={"xl"} variant={"light"} />
      <Button variant="filled" color="rgba(242, 41, 41, 1)">
        Выйти
      </Button>
    </Box>
  );
};

export default NavBar;
