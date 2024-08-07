import React, { useState, useRef } from "react";
import { Button, Col, Row } from "antd";
import Editor from "./Editor";

const DualEditor = ({
  leftLabel,
  rightLabel,
  leftLanguage,
  rightLanguage,
  leftReadOnly = false,
  rightReadOnly = false,
  leftContent,
  rightContent,
  onLeftChange,
  onRightChange,
  firstButtonLabel,
  firstButtonAction,
  secondButtonLabel,
  secondButtonAction,
}) => {
  const [leftValue, setLeftValue] = useState(leftContent || "");
  const [rightValue, setRightValue] = useState(rightContent || "");
  const leftEditorRef = useRef(null);
  const rightEditorRef = useRef(null);

  const handleLeftChange = (value) => {
    setLeftValue(value || "");
    if (onLeftChange) onLeftChange(value || "");
  };

  const handleRightChange = (value) => {
    setRightValue(value || "");
    if (onRightChange) onRightChange(value || "");
  };

  const handleFirstButtonClick = () => {
    if (firstButtonAction) {
      firstButtonAction(leftValue, rightValue, setLeftValue, setRightValue);
    }
  };

  const handleSecondButtonClick = () => {
    if (secondButtonAction) {
      secondButtonAction(leftValue, rightValue, setLeftValue, setRightValue);
    }
  };

  return (
    <div
      style={{
        padding: "20px",
        height: "calc(100vh - 40px)",
        boxSizing: "border-box",
      }}
    >
      <Row gutter={8} style={{ height: "calc(100% - 40px)", margin: 0 }}>
        <Col span={12} style={{ paddingRight: "8px", height: "100%" }}>
          <Editor
            height="calc(100% - 32px)"
            readOnly={leftReadOnly}
            language={leftLanguage}
            value={leftValue}
            onChange={handleLeftChange}
            editorDidMount={(editor) => (leftEditorRef.current = editor)}
            label={leftLabel}
          />
        </Col>
        <Col span={12} style={{ paddingLeft: "8px", height: "100%" }}>
          <Editor
            height="calc(100% - 32px)"
            language={rightLanguage}
            readOnly={rightReadOnly}
            value={rightValue}
            onChange={handleRightChange}
            editorDidMount={(editor) => (rightEditorRef.current = editor)}
            label={rightLabel}
          />
        </Col>
      </Row>
      <div style={{ textAlign: "center", marginTop: "16px" }}>
        {firstButtonLabel && (
          <Button
            type="primary"
            onClick={handleFirstButtonClick}
            style={{ marginRight: "8px" }}
          >
            {firstButtonLabel}
          </Button>
        )}
        {secondButtonLabel && (
          <Button type="default" onClick={handleSecondButtonClick}>
            {secondButtonLabel}
          </Button>
        )}
      </div>
    </div>
  );
};

export default DualEditor;
