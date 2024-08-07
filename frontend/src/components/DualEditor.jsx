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

  // Handle editor content changes
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
    <Row
      style={{
        padding: "20px",
        height: "calc(100vh - 40px)",
      }}
    >
      <Col
        span={11}
        style={{ paddingRight: "8px", height: "100%", border: "none" }}
      >
        <div
          style={{
            height: "100%",
            textAlign: "left",
          }}
        >
          <h4>{leftLabel}</h4>
          <Editor
            height="calc(100% - 32px)" // Subtract the height of the label
            readOnly={leftReadOnly}
            language={leftLanguage}
            value={leftValue}
            onChange={handleLeftChange}
            editorDidMount={(editor) => (leftEditorRef.current = editor)}
          />
        </div>
      </Col>
      <Col span={2} style={{ textAlign: "center", height: "100%" }}>
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            justifyContent: "center",
            height: "100%",
          }}
        >
          {firstButtonLabel && (
            <Button
              type="primary"
              onClick={handleFirstButtonClick}
              style={{ marginBottom: "8px" }}
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
      </Col>
      <Col span={11} style={{ paddingLeft: "8px", height: "100%" }}>
        <div style={{ height: "100%", textAlign: "left" }}>
          <h4>{rightLabel}</h4>
          <Editor
            height="calc(100% - 32px)"
            language={rightLanguage}
            readOnly={rightReadOnly}
            value={rightValue}
            onChange={handleRightChange}
            editorDidMount={(editor) => (rightEditorRef.current = editor)}
          />
        </div>
      </Col>
    </Row>
  );
};
export default DualEditor;
