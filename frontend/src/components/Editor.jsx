import React from "react";
import { Editor as MonacoEditor } from "@monaco-editor/react";
import { useTheme } from "../contexts/ThemeContext";

const Editor = ({
  height,
  language,
  value,
  readOnly = false,
  onTextChange,
  editorDidMount,
}) => {
  const { isDark } = useTheme();

  const handleEditorDidMount = (editor, monaco) => {
    if (editorDidMount) {
      editorDidMount(editor, monaco);
    }
    editor.updateOptions({ readOnly: readOnly });
  };

  return (
    <div style={{ height: "100%", display: "flex", flexDirection: "column", minHeight: "300px" }}>
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
          readOnly: readOnly,
        }}
        editorDidMount={handleEditorDidMount}
      />
    </div>
  );
};

export default Editor;
