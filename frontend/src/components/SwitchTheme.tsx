import {
  useMantineColorScheme,
  useComputedColorScheme,
  ActionIcon,
} from "@mantine/core";
import { IconSun, IconMoon } from "@tabler/icons-react";

const SwitchTheme = () => {
  const { setColorScheme } = useMantineColorScheme();
  const computedColorScheme = useComputedColorScheme("light");

  const Switch = () => {
    setColorScheme(computedColorScheme === "light" ? "dark" : "light");
  };
  return (
    <ActionIcon
      style={{
        position: "absolute",
        top: "50%",
      }}
      onClick={Switch}
      variant="default"
      size="xl"
      aria-label="Toggle color scheme"
    >
      {computedColorScheme === "light" ? <IconSun /> : <IconMoon />}
    </ActionIcon>
  );
};
export default SwitchTheme;
