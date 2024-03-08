import { useForm } from '@mantine/form';
import { TextInput, Button, Box } from '@mantine/core';
const AuthForm = () => {
    const form = useForm({
        initialValues: { name: '', email: '', age: 0 },

        // functions will be used to validate values at corresponding key
        validate: {
            name: (value) => (value.length < 2 ? 'Name must have at least 2 letters' : null),
            email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
        },
    });

    return (
        <Box maw={340} m="auto auto"  >
            <form onSubmit={form.onSubmit(console.log)}>
                <TextInput label="Name"  size="md" placeholder="Name" {...form.getInputProps('name')} />
                <TextInput mt="sm" size="md"  label="Email" placeholder="Email" {...form.getInputProps('email')} />
                <Button type="submit" mt="sm">
                    Submit
                </Button>
            </form>
        </Box>
    );
};

export default AuthForm;
