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
    editor.updateOptions({ readOnly: readOnly });
  };

  return (
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
  );
};

export default Editor;
