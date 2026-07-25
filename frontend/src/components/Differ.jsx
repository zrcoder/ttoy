import React, { useState } from "react";
import { DiffEditor } from "@monaco-editor/react";
import { Button } from "antd";
import { JSONSort } from "../../bindings/github.com/zrcoder/ttoy/service/service";
import { App as AntdApp } from "antd";
import { useTheme } from "../contexts/ThemeContext";
import AppTabs from "./AppTabs";

const TextDiffTab = () => {
  const { isDark } = useTheme();
  return (
    <div
      style={{
        height: "calc(100vh - 140px)",
        margin: 0,
        boxSizing: "border-box",
        backgroundColor: "transparent",
      }}
    >
      <DiffEditor
        height="100%"
        language="text"
        theme={isDark ? "vs-dark" : "vs"}
        options={{
          originalEditable: true,
          readOnly: false,
          scrollBeyondLastLine: false,
        }}
      />
    </div>
  );
};

const JsonDiffTab = () => {
  const { isDark } = useTheme();
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
        height: "calc(100vh - 108px)",
        margin: 0,
        boxSizing: "border-box",
        backgroundColor: "transparent",
      }}
    >
      <DiffEditor
        height="calc(100vh - 140px)"
        language="json"
        theme={isDark ? "vs-dark" : "vs"}
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
      <div style={{ textAlign: "center", padding: "8px" }}>
        <Button type="primary" onClick={handleCompare}>
          Sort and Compare
        </Button>
      </div>
    </div>
  );
};

const Differ = () => {
  const items = [
    {
      key: "text",
      label: "Text Diff",
      children: <TextDiffTab />,
    },
    {
      key: "json",
      label: "JSON Diff",
      children: <JsonDiffTab />,
    },
  ];

  return <AppTabs defaultActiveKey="text" items={items} />;
};

export default Differ;