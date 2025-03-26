import { useCodeMirror, basicSetup, EditorView } from "@uiw/react-codemirror";
import { useCallback, useState } from "react";
import { markdown, markdownLanguage } from "@codemirror/lang-markdown";
import { languages } from "@codemirror/language-data";
import { HighlightStyle, syntaxHighlighting } from "@codemirror/language";
import { githubDark, githubLight } from "@uiw/codemirror-theme-github";
import { tags as t } from "@lezer/highlight";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { useTheme } from "./theme-provider";
import "github-markdown-css";

const markdownHighlightStyle = HighlightStyle.define([
  { tag: t.heading1, fontSize: "2em", fontWeight: "bold" },
  { tag: t.heading2, fontSize: "1.75em", fontWeight: "bold" },
  { tag: t.heading3, fontSize: "1.5em", fontWeight: "bold" },
]);

const myTheme = EditorView.theme({
  "&": {
    backgroundColor: "transparent !important",
  },
});

function Editor() {
  const { theme } = useTheme();
  const resolvedTheme =
    theme === "system"
      ? window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light"
      : theme;

  const [value, setValue] = useState<string>("# Welcome to Redoed!");
  const handleChange = useCallback((val: string) => setValue(val), []);

  const { setContainer } = useCodeMirror({
    value,
    height: "90vh",
    extensions: [
      basicSetup(),
      markdown({
        base: markdownLanguage,
        codeLanguages: languages,
        addKeymap: true,
      }),
      syntaxHighlighting(markdownHighlightStyle),
      EditorView.lineWrapping,
      myTheme,
    ],
    theme: resolvedTheme === "dark" ? githubDark : githubLight,
    onChange: handleChange,
  });

  return (
    <div className="w-screen sm:grid sm:grid-cols-2 py-2">
      <div>
        <div ref={setContainer} className="border-r-1" />
      </div>
      <div>
        <div className="h-[90vh] markdown-body p-2 markdown-preview overflow-y-auto scrollbar hidden sm:block !bg-background">
          <ReactMarkdown remarkPlugins={[remarkGfm]}>{value}</ReactMarkdown>
        </div>
      </div>
    </div>
  );
}

export default Editor;
