import { Box, Title } from "@mantine/core";

const NotFound = () => {
  return (
    <Box
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <Title order={1}>Такой страницы не существует 👀 </Title>
    </Box>
  );
};

export default NotFound;
