import React, { useState } from "react";
import { DiffEditor } from "@monaco-editor/react";
import { Button } from "antd";
import { JSONSort } from "../../bindings/github.com/zrcoder/ttoy/service/service";
import { App as AntdApp } from "antd";

const JsonDiffer = () => {
  const { modal } = AntdApp.useApp();
  const [original, setOriginal] = useState("");
  const [modified, setModified] = useState("");

  const handleCompare = async () => {
    try {
      const [sortedOriginal, sortedModified] = await Promise.all([
        original.trim() ? JSONSort(original) : Promise.resolve(""),
        modified.trim() ? JSONSort(modified) : Promise.resolve(""),
      ]);
      setOriginal(sortedOriginal);
      setModified(sortedModified);
    } catch (err) {
      modal.error({ content: err.toString() });
    }
  };

  return (
    <div
      style={{
        height: "calc(100vh - 40px)",
        margin: 0,
        padding: "10px",
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
          padding: "0 4px",
        }}
      >
        <header
          style={{
            padding: "8px",
            borderBottom: "1px solid #d9d9d9",
            backgroundColor: "transparent",
          }}
        >
          <h4 style={{ margin: 0 }}>JSON Difference</h4>
        </header>
        <DiffEditor
          height="calc(100% - 40px)"
          language="json"
          theme="vs-dark"
          original={original}
          modified={modified}
          options={{
            originalEditable: true,
            readOnly: false,
            scrollBeyondLastLine: false,
          }}
          onMount={(editor) => {
            const originalEditor = editor.getOriginalEditor();
            originalEditor.onDidChangeModelContent(() => {
              setOriginal(originalEditor.getValue());
            });
            const modifiedEditor = editor.getModifiedEditor();
            modifiedEditor.onDidChangeModelContent(() => {
              setModified(modifiedEditor.getValue());
            });
          }}
        />
        <div style={{ textAlign: "center", marginTop: "16px" }}>
          <Button type="primary" onClick={handleCompare}>
            Sort and Compare
          </Button>
        </div>
      </div>
    </div>
  );
};

export default JsonDiffer;
