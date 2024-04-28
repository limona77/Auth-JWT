import NavBar from "../components/NavBar.tsx";
import Demo from "../components/Demo.tsx";

const DemoPage = () => {
  return (
    <>
      <NavBar />,
      <Demo
        display={"flex"}
        justifyContent={"center"}
        height={"100vh"}
        alignItems={"center"}
        fontSize={"40px"}
      >
        ✅ Вы авторизованы
      </Demo>
      ,
    </>
  );
};

export default DemoPage;
