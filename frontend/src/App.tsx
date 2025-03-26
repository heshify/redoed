import Editor from "@/components/markdown-editor";
import Header from "@/components/header";
import { ThemeProvider } from "./components/theme-provider";

function App() {
  return (
    <ThemeProvider>
      <main className="h-fit">
        <Header />
        <Editor />
      </main>
    </ThemeProvider>
  );
}

export default App;
