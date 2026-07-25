import React from "react";
import { DiffEditor } from "@monaco-editor/react";
import { useTheme } from "../contexts/ThemeContext";

const TextDiffer = () => {
  const { isDark } = useTheme();
  return (
    <div
      style={{
        height: "calc(100vh - 40px)",
        margin: 0,
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
        }}
      >
        <header
          style={{
            padding: "8px",
            borderBottom: "1px solid #d9d9d9",
            backgroundColor: "transparent",
          }}
        >
          <h4 style={{ margin: 0 }}>Text Difference</h4>
        </header>
        <DiffEditor
          height="calc(100% - 40px)" // Adjust height to account for the label
          language="text"
          theme={isDark ? "vs-dark" : "vs"}
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
