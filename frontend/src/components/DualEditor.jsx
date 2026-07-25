import React, { useState, useRef } from "react";
import { Button } from "antd";
import { CaretRightFilled, CaretLeftFilled } from "@ant-design/icons";
import Editor from "./Editor";

const DualEditor = ({
  leftLanguage,
  rightLanguage,
  leftContent,
  rightContent,
  buttonAction,
  reverseButtonAction,
  leftReadOnly = false,
  rightReadOnly = true,
}) => {
  const [leftValue, setLeftValue] = useState(leftContent || "");
  const [rightValue, setRightValue] = useState(rightContent || "");
  const leftEditorRef = useRef(null);
  const rightEditorRef = useRef(null);

  const handleLeftChange = (value) => {
    setLeftValue(value || "");
  };

  const handleRightChange = (value) => {
    setRightValue(value || "");
  };

  const handleButtonClick = () => {
    if (buttonAction) {
      buttonAction(leftValue, rightValue, setLeftValue, setRightValue);
    }
  };

  const handleReverseButtonClick = () => {
    if (reverseButtonAction) {
      reverseButtonAction(leftValue, rightValue, setLeftValue, setRightValue);
    }
  };

  return (
    <div
      style={{
        display: "flex",
        height: "calc(100vh - 120px)",
        alignItems: "stretch",
        boxSizing: "border-box",
      }}
    >
      <div style={{ flex: 1, minWidth: 0 }}>
        <Editor
          height="100%"
          language={leftLanguage}
          value={leftValue}
          onTextChange={handleLeftChange}
          editorDidMount={(editor) => (leftEditorRef.current = editor)}
          readOnly={leftReadOnly}
        />
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          gap: "8px",
          padding: "0 8px",
        }}
      >
        <Button
          onClick={handleButtonClick}
          icon={<CaretRightFilled />}
          type="primary"
        ></Button>
        {reverseButtonAction && (
          <Button
            onClick={handleReverseButtonClick}
            icon={<CaretLeftFilled />}
            type="primary"
          ></Button>
        )}
      </div>
      <div style={{ flex: 1, minWidth: 0 }}>
        <Editor
          height="100%"
          language={rightLanguage}
          value={rightValue}
          onTextChange={handleRightChange}
          editorDidMount={(editor) => (rightEditorRef.current = editor)}
          readOnly={rightReadOnly}
        />
      </div>
    </div>
  );
};

export default DualEditor;
