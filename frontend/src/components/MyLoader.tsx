import { Box, Loader } from "@mantine/core";

const MyLoader = () => {
  return (
    <Box
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <Loader color="blue" size="xl" type="dots" />
    </Box>
  );
};

export default MyLoader;
