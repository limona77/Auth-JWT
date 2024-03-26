import { FC } from "react";
import { Box } from "@mantine/core";

interface IDemoProps {
  display: string | undefined;
  justifyContent: string | undefined;
  height: string | undefined;
  alignItems: string | undefined;
  fontSize: string | undefined;
  children: React.ReactNode ;
}

const Demo: FC<IDemoProps> = ({
  children,
  display,
  justifyContent,
  height,
  alignItems,
  fontSize,
}) => {
  return (
    <Box
      style={{
        display,
        justifyContent,
        height,
        alignItems,
        fontSize,
      }}
    >
      {children}
    </Box>
  );
};
//❌ Вы должны авторизироваться, чтобы просматривать данную страницу
//✅ Вы вошли в демонстрационную версию
export default Demo;
