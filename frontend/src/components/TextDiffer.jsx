import React from "react";
import { DiffEditor } from "@monaco-editor/react";

const TextDiffer = () => {
  return (
    <div
      style={{
        height: "calc(100vh - 40px)",
        margin: 0,
        padding: "20px",
        boxSizing: "border-box",
        backgroundColor: "transparent",
      }}
    >
      <div
        style={{
          height: "100%",
          display: "flex",
          flexDirection: "column",
          backgroundColor: "transparent",
          padding: "0 4px", // Adjust padding to match DualEditor
        }}
      >
        <header
          style={{
            padding: "8px",
            borderBottom: "1px solid #d9d9d9",
            backgroundColor: "transparent",
          }}
        >
          <h4 style={{ margin: 0 }}>Text Diff</h4>
        </header>
        <DiffEditor
          height="calc(100% - 40px)" // Adjust height to account for the label
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
