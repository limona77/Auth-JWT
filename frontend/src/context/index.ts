import { createContext } from "react";

import { ISwitchThemeProps } from "../components/NavBar.tsx";

export const SwitchThemeContext = createContext<ISwitchThemeProps | undefined>(
  undefined,
);

export const IsAuthContext = createContext<boolean>(false);
