import { AuthService } from "../../services/auth";

export const fetchRegister = async (email: string, password: string) => {
  const response = await AuthService.register(email, password);
  localStorage.setItem("token", response.data.accessToken);
  alert("Вы зарегистрированы!");
};

export const fetchLogin = async (email: string, password: string) => {
  const response = await AuthService.login(email, password);
  localStorage.setItem("token", response.data.accessToken);
  alert("Вы успешно вошли в свой аккаунт!");
};
