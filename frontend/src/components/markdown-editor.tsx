import { useCodeMirror, basicSetup } from "@uiw/react-codemirror";
import { useCallback, useState } from "react";
import { markdown, markdownLanguage } from "@codemirror/lang-markdown";
import { languages } from "@codemirror/language-data";
import { HighlightStyle, syntaxHighlighting } from "@codemirror/language";
import { tokyoNight } from "@uiw/codemirror-theme-tokyo-night";
import { tags as t } from "@lezer/highlight";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

const markdownHighlightStyle = HighlightStyle.define([
  { tag: t.heading1, fontSize: "2em", fontWeight: "bold" },
  { tag: t.heading2, fontSize: "1.75em", fontWeight: "bold" },
  { tag: t.heading3, fontSize: "1.5em", fontWeight: "bold" },
]);

function Editor() {
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
    ],
    theme: tokyoNight,
    onChange: handleChange,
  });

  return (
    <div className="grid grid-cols-2 py-2">
      <div>
        <div ref={setContainer} className="border rounded p-2" />
      </div>
      <div>
        <div className="border h-[90vh] rounded p-2 markdown-preview overflow-scroll">
          <ReactMarkdown remarkPlugins={[remarkGfm]}>{value}</ReactMarkdown>
        </div>
      </div>
    </div>
  );
}

export default Editor;
