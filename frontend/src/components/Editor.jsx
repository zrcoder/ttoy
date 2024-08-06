import React from "react";
import { Editor as MonacoEditor } from "@monaco-editor/react";

const Editor = ({
  height,
  language,
  value,
  readOnly = false,
  onChange,
  editorDidMount,
}) => {
  const handleEditorDidMount = (editor, monaco) => {
    if (editorDidMount) {
      editorDidMount(editor, monaco);
    }
    // 确保 readOnly 选项被正确应用
    editor.updateOptions({ readOnly: readOnly });
  };

  return (
    <MonacoEditor
      height={height}
      language={language}
      value={value}
      onChange={onChange}
      theme="vs-dark"
      options={{
        minimap: { enabled: false },
        scrollBeyondLastLine: false,
        lineNumbers: "on",
        renderWhitespace: "none",
        renderControlCharacters: false,
        overviewRulerLanes: 0,
        readOnly: readOnly, // 将 readOnly 传递到 options 中
      }}
      editorDidMount={handleEditorDidMount}
    />
  );
};

export default Editor;
