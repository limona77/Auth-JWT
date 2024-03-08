
import './App.css'
import AuthForm from "./components/AuthForm";
import '@mantine/core/styles.css';

import { MantineProvider } from '@mantine/core';
function App() {

  return (<MantineProvider><AuthForm/></MantineProvider>)
}

export default App
