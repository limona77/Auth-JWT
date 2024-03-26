import NavBar from "../components/NavBar.tsx";
import Demo from "../components/Demo.tsx";

const ErrorAuthPage = () => {
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
        ❌ Вы должны авторизироваться, чтобы просматривать данную страницу
      </Demo>
    </>
  );
};

export default ErrorAuthPage;
