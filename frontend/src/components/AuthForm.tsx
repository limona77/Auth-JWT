import { useForm } from "@mantine/form";
import { TextInput, Button, Box, PasswordInput } from "@mantine/core";

import { fetchLogin, fetchRegister } from "../api/auth";
const AuthForm = () => {
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
            label="Email"
            placeholder="Email"
            {...form.getInputProps("email")}
          />
          <PasswordInput
            size="md"
            label="Password"
            withAsterisk
            placeholder="Input placeholder"
            {...form.getInputProps("password")}
          />
          <Button
            mt="sm"
            fullWidth={true}
            onClick={() => {
              fetchRegister(form.values.email, form.values.password);
            }}
          >
            Зарегистрироваться
          </Button>
          <Button
            mt="sm"
            fullWidth={true}
            onClick={() => {
              fetchLogin(form.values.email, form.values.password);
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
