import React from "react";
import { Editor as MonacoEditor } from "@monaco-editor/react";

const Editor = ({
  height,
  language,
  value,
  readOnly = false,
  onChange,
  editorDidMount,
  label,
}) => {
  const handleEditorDidMount = (editor, monaco) => {
    if (editorDidMount) {
      editorDidMount(editor, monaco);
    }
    editor.updateOptions({ readOnly: readOnly });
  };

  return (
    <div style={{ height: "100%", display: "flex", flexDirection: "column" }}>
      {label && (
        <div style={{ padding: "8px", borderBottom: "1px solid #d9d9d9" }}>
          <h4 style={{ margin: 0 }}>{label}</h4>
        </div>
      )}
      <MonacoEditor
        loading={null} // 去掉默认的 loading 文字
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
          readOnly: readOnly,
        }}
        editorDidMount={handleEditorDidMount}
      />
    </div>
  );
};

export default Editor;
