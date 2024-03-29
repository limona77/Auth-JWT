import {
  ActionIcon,
  useComputedColorScheme,
  useMantineColorScheme,
} from "@mantine/core";
import { IconSun, IconMoon } from "@tabler/icons-react";
import { FC } from "react";

export interface ISwitchThemeProps {
  size: string;
  variant: string;
}
export const SwitchTheme: FC<ISwitchThemeProps> = ({ size, variant }) => {
  const { setColorScheme } = useMantineColorScheme();
  const computedColorScheme = useComputedColorScheme("light");
  const Switch = () => {
    setColorScheme(computedColorScheme === "light" ? "dark" : "light");
  };
  return (
    <ActionIcon onClick={Switch} variant={variant} size={size}>
      {computedColorScheme === "light" ? <IconSun /> : <IconMoon />}
    </ActionIcon>
  );
};

export default SwitchTheme;
