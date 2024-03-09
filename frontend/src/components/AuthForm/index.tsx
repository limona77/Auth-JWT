import { useForm } from "@mantine/form";
import { TextInput, Button, Box, PasswordInput } from "@mantine/core";
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
        // backgroundColor:'orange'
      }}
    >
      <Box
        style={{
          minWidth: 140,
          width: 340,
          maxHeight: 240,
          // backgroundColor:'red'
        }}
      >
        <form
          onSubmit={form.onSubmit(() => {
            return;
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
