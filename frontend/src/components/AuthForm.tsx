import { useForm } from "@mantine/form";
import { TextInput, Button, Box, PasswordInput } from "@mantine/core";

import { useAppDispatch } from "../store";
import {
  fetchLogin,
  fetchRegister,
} from "../store/slices/auth/asyncActions.ts";
const AuthForm = () => {
  const dispatch = useAppDispatch();
  const form = useForm({
    initialValues: { email: "", password: "" },
    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      password: (value) =>
        value.length < 5 ? "Длина пароля должна быть минимум 5 символов" : null,
    },
  });
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
            size="md"
            label="Пароль"
            withAsterisk
            placeholder="Введите свой пароль"
            {...form.getInputProps("password")}
          />
          <Button
            mt="sm"
            fullWidth={true}
            onClick={() => {
              dispatch(
                fetchRegister({
                  email: form.values.email,
                  password: form.values.password,
                }),
              );
            }}
          >
            Зарегистрироваться
          </Button>
          <Button
            mt="sm"
            fullWidth={true}
            onClick={() => {
              dispatch(
                fetchLogin({
                  email: form.values.email,
                  password: form.values.password,
                }),
              );
            }}
          >
            Войти
          </Button>
        </form>
      </Box>
    </Box>
  );
};

export default AuthForm;
