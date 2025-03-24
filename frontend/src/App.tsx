import Editor from "@/components/markdown-editor";
import Header from "@/components/header";
import { ThemeProvider } from "./components/theme-provider";

function App() {
  return (
    <ThemeProvider>
      <div>
        <Header />
        <Editor />
      </div>
    </ThemeProvider>
  );
}

export default App;
