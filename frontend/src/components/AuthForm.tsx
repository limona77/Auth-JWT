import { useForm } from "@mantine/form";
import { TextInput, Button, Box, PasswordInput } from "@mantine/core";

import { useNavigate } from "react-router-dom";

import { useAppDispatch } from "../store";
import {
  fetchLogin,
  fetchRegister,
} from "../store/slices/auth/asyncActions.ts";
const AuthForm = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const form = useForm({
    initialValues: { email: "", password: "" },
    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      password: (value) => {
        return value.length < 5
          ? "Длина пароля должна быть минимум 5 символов"
          : null;
      },
    },

    validateInputOnBlur: true,
  });

  const handleLoginEvent = async () => {
    const res = await dispatch(
      fetchLogin({
        email: form.values.email,
        password: form.values.password,
      }),
    );

    if (res.meta.requestStatus == "fulfilled") {
      navigate("/demo");
    }
  };
  const handleRegisterEvent = async () => {
    const res = await dispatch(
      fetchRegister({
        email: form.values.email,
        password: form.values.password,
      }),
    );
    if (res.meta.requestStatus == "fulfilled") {
      navigate("/demo");
    }
  };
  return (
    <Box
      style={{
        display: "flex",
        justifyContent: "center",
        height: "100vh",
        alignItems: "center",
      }}
    >
      <Box
        style={{
          minWidth: 140,
          width: 340,
          maxHeight: 240,
        }}
      >
        <form>
          <TextInput
            mt="sm"
            withAsterisk
            size="md"
            label="Почта"
            placeholder="Введите свою почту"
            {...form.getInputProps("email")}
          />
          <PasswordInput
            mt="sm"
            size="md"
            label="Пароль"
            withAsterisk
            placeholder="Введите свой пароль"
            error={form.getInputProps("password").error}
            {...form.getInputProps("password")}
          />
          <Button mt="sm" fullWidth={true} onClick={handleRegisterEvent}>
            Зарегистрироваться
          </Button>
          <Button mt="sm" fullWidth={true} onClick={handleLoginEvent}>
            Войти
          </Button>
        </form>
      </Box>
    </Box>
  );
};

export default AuthForm;
