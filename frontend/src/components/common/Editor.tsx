import { Editor as MonacoEditor, OnMount } from "@monaco-editor/react";
import { useTheme } from "../../contexts/ThemeContext";
import { CopyButton } from "./CopyButton";

type EditorProps = {
  height: string;
  language: string;
  value: string;
  readOnly?: boolean;
  onTextChange?: (value: string) => void;
  editorDidMount?: OnMount;
};

const Editor = ({
  height,
  language,
  value,
  readOnly = false,
  onTextChange,
  editorDidMount,
}: EditorProps) => {
  const { isDark } = useTheme();

  const handleEditorDidMount: OnMount = (editor, monaco) => {
    if (editorDidMount) {
      editorDidMount(editor, monaco);
    }
    editor.updateOptions({ readOnly: readOnly });
  };

  const handleChange = (val: string | undefined) => {
    if (onTextChange) {
      onTextChange(val ?? "");
    }
  };

  return (
    <div
      style={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        minHeight: "300px",
        position: "relative",
      }}
    >
      <CopyButton text={value} />
      <MonacoEditor
        loading={null}
        height={height}
        language={language}
        value={value}
        onChange={handleChange}
        theme={isDark ? "vs-dark" : "vs"}
        options={{
          minimap: { enabled: false },
          scrollBeyondLastLine: false,
          lineNumbers: "on",
          renderWhitespace: "none",
          renderControlCharacters: false,
          overviewRulerLanes: 0,
        }}
        onMount={handleEditorDidMount}
      />
    </div>
  );
};

export default Editor;
