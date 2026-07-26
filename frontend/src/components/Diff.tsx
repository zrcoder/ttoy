import { useState } from "react";
import { DiffEditor } from "@monaco-editor/react";
import { Button } from "antd";
import { Sort as SvcSort } from "../../bindings/github.com/zrcoder/ttoy/service";
import { App as AntdApp } from "antd";
import { useTheme } from "../contexts/ThemeContext";
import AppTabs from "./common/AppTabs";
import { contentHeight } from "./common/layout";

type DiffTabProps = {
  language: string;
  original?: string;
  modified?: string;
  onOriginalChange?: (value: string) => void;
  onModifiedChange?: (value: string) => void;
};

const DiffTab = ({
  language,
  original = "",
  modified = "",
  onOriginalChange,
  onModifiedChange,
}: DiffTabProps) => {
  const { isDark } = useTheme();
  return (
    <div
      style={{
        height: "100%",
        margin: 0,
        boxSizing: "border-box",
        backgroundColor: "transparent",
      }}
    >
      <DiffEditor
        height={contentHeight}
        language={language}
        theme={isDark ? "vs-dark" : "vs"}
        original={original}
        modified={modified}
        options={{
          originalEditable: true,
          readOnly: false,
          scrollBeyondLastLine: false,
          renderOverviewRuler: false,
        }}
        onMount={(editor) => {
          editor
            .getOriginalEditor()
            .onDidChangeModelContent(() =>
              onOriginalChange?.(editor.getOriginalEditor().getValue()),
            );
          editor
            .getModifiedEditor()
            .onDidChangeModelContent(() =>
              onModifiedChange?.(editor.getModifiedEditor().getValue()),
            );
        }}
      />
    </div>
  );
};

const JsonDiffTab = () => {
  const { modal } = AntdApp.useApp();
  const [original, setOriginal] = useState("");
  const [modified, setModified] = useState("");

  const handleCompare = async () => {
    try {
      const [sortedOriginal, sortedModified] = await Promise.all([
        original.trim() ? SvcSort.JSON(original) : Promise.resolve(""),
        modified.trim() ? SvcSort.JSON(modified) : Promise.resolve(""),
      ]);
      setOriginal(sortedOriginal ?? "");
      setModified(sortedModified ?? "");
    } catch (err: unknown) {
      modal.error({ content: (err as Error).toString() });
    }
  };

  return (
    <div style={{ position: "relative", height: "100%" }}>
      <DiffTab
        language="json"
        original={original}
        modified={modified}
        onOriginalChange={setOriginal}
        onModifiedChange={setModified}
      />
      <Button
        type="primary"
        onClick={handleCompare}
        style={{
          position: "absolute",
          bottom: 24,
          left: "50%",
          transform: "translateX(-50%)",
        }}
      >
        Sort
      </Button>
    </div>
  );
};

const Diff = () => {
  const items = [
    { key: "text", label: "Text", children: <DiffTab language="text" /> },
    { key: "json", label: "JSON", children: <JsonDiffTab /> },
  ];
  return <AppTabs defaultActiveKey="text" items={items} />;
};

export default Diff;
