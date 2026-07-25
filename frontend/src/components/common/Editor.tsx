import { Editor as MonacoEditor, OnMount } from "@monaco-editor/react";
import { useTheme } from "../../contexts/ThemeContext";

type EditorProps = {
  height: string;
  language: string;
  value: string;
  readOnly?: boolean;
  onTextChange?: (value: string | undefined) => void;
  editorDidMount?: (editor: any, monaco: any) => void;
};

const Editor = ({ height, language, value, readOnly = false, onTextChange, editorDidMount }: EditorProps) => {
  const { isDark } = useTheme();

  const handleEditorDidMount: OnMount = (editor, monaco) => {
    if (editorDidMount) {
      editorDidMount(editor, monaco);
    }
    editor.updateOptions({ readOnly: readOnly });
  };

  return (
    <div
      style={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        minHeight: "300px",
      }}
    >
      <MonacoEditor
        loading={null}
        height={height}
        language={language}
        value={value}
        onChange={onTextChange}
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
