import { CaretLeftFilled, CaretRightFilled } from "@ant-design/icons";
import { Button } from "antd";
import { useState } from "react";
import Editor from "./Editor";
import { contentHeight } from "./layout";

type DualEditorProps = {
  leftLanguage: string;
  rightLanguage: string;
  leftContent?: string;
  rightContent?: string;
  buttonAction?: (
    l: string,
    r: string,
    setL: (v: string) => void,
    setR: (v: string) => void,
  ) => void;
  reverseButtonAction?: (
    l: string,
    r: string,
    setL: (v: string) => void,
    setR: (v: string) => void,
  ) => void;
  leftReadOnly?: boolean;
  rightReadOnly?: boolean;
};

const DualEditor = ({
  leftLanguage,
  rightLanguage,
  leftContent,
  rightContent,
  buttonAction,
  reverseButtonAction,
  leftReadOnly = false,
  rightReadOnly = true,
}: DualEditorProps) => {
  const [leftValue, setLeftValue] = useState(leftContent ?? "");
  const [rightValue, setRightValue] = useState(rightContent ?? "");

  return (
    <div
      style={{
        display: "flex",
        height: contentHeight,
        alignItems: "stretch",
        boxSizing: "border-box",
      }}
    >
      <div style={{ flex: 1, minWidth: 0 }}>
        <Editor
          height="100%"
          language={leftLanguage}
          value={leftValue}
          onTextChange={setLeftValue}
          readOnly={leftReadOnly}
        />
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          gap: 8,
          padding: "0 8px",
        }}
      >
        <Button
          onClick={() =>
            buttonAction?.(leftValue, rightValue, setLeftValue, setRightValue)
          }
          icon={<CaretRightFilled />}
          type="primary"
        />
        {reverseButtonAction && (
          <Button
            onClick={() =>
              reverseButtonAction(
                leftValue,
                rightValue,
                setLeftValue,
                setRightValue,
              )
            }
            icon={<CaretLeftFilled />}
            type="primary"
          />
        )}
      </div>
      <div style={{ flex: 1, minWidth: 0 }}>
        <Editor
          height="100%"
          language={rightLanguage}
          value={rightValue}
          onTextChange={setRightValue}
          readOnly={rightReadOnly}
        />
      </div>
    </div>
  );
};

export default DualEditor;
