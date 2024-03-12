import { useForm } from "@mantine/form";
import { TextInput, Button, Box, PasswordInput } from "@mantine/core";

import { AuthService } from "../../services/auth";

const AuthForm = () => {
  const form = useForm({
    initialValues: { email: "", password: "" },
    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      password: (value) =>
        value.length < 5 ? "Длина пароля должна быть минимум 5 символов" : null,
    },
  });

  const fetchRegister = async (email: string, password: string) => {
    try {
      await AuthService.register(email, password);
      alert("Вы зарегистрированы!");
    } catch (err) {
      alert(err);
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
        <form
          onSubmit={form.onSubmit(() => {
            fetchRegister(form.values.email, form.values.password);
          })}
        >
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
          <Button type="submit" mt="sm">
            Зарегистрироваться
          </Button>
        </form>
      </Box>
    </Box>
  );
};

export default AuthForm;
