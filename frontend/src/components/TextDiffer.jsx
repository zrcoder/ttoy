import React from "react";
import { DiffEditor } from "@monaco-editor/react";

const TextDiffer = () => {
  return (
    <div style={{ height: "100vh", margin: 0 }}>
      <div
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          padding: "20px",
          height: "calc(100vh - 40px)",
        }}
      >
        <DiffEditor
          height="calc(100% - 32px)"
          language="text"
          theme="vs-dark"
          options={{
            originalEditable: true,
            readOnly: false,
            scrollBeyondLastLine: false,
          }}
        />
      </div>
    </div>
  );
};

export default TextDiffer;
