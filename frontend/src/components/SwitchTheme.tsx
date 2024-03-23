import {
  ActionIcon,
  useComputedColorScheme,
  useMantineColorScheme,
} from "@mantine/core";
import { IconSun, IconMoon } from "@tabler/icons-react";
import { useContext } from "react";

import { SwitchThemeContext } from "../context";
export const SwitchTheme = () => {
  const { setColorScheme } = useMantineColorScheme();
  const computedColorScheme = useComputedColorScheme("light");
  const switchThemeProps = useContext(SwitchThemeContext);
  const Switch = () => {
    setColorScheme(computedColorScheme === "light" ? "dark" : "light");
  };
  return (
    <ActionIcon
      onClick={Switch}
      variant={switchThemeProps?.variant}
      size={switchThemeProps?.size}
    >
      {computedColorScheme === "light" ? <IconSun /> : <IconMoon />}
    </ActionIcon>
  );
};

export default SwitchTheme;
