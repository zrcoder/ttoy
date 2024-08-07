import React from "react";
import { Editor as MonacoEditor } from "@monaco-editor/react";
import LangSelector from "./LangSelector";

const Editor = ({
  height,
  language,
  value,
  readOnly = false,
  onTextChange,
  editorDidMount,
  languages,
  onLanguageChange,
}) => {
  const handleEditorDidMount = (editor, monaco) => {
    if (editorDidMount) {
      editorDidMount(editor, monaco);
    }
    editor.updateOptions({ readOnly: readOnly });
  };

  return (
    <div style={{ height: "100%", display: "flex", flexDirection: "column" }}>
      {languages && languages.length > 0 && (
        <LangSelector
          value={language}
          onChange={onLanguageChange}
          languages={languages}
        />
      )}
      <MonacoEditor
        loading={null} // 去掉默认的 loading 文字
        height={height}
        language={language}
        value={value}
        onChange={onTextChange}
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
