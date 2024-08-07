import React, { useState, useRef } from "react";
import { Button, Col, Row } from "antd";
import Editor from "./Editor";

const DualEditor = ({
  leftLanguage,
  rightLanguage,
  leftContent,
  rightContent,
  onLeftChange,
  onRightChange,
  buttonLabel,
  buttonAction,
  languages = [],
  onLeftLangChange,
  onRightLangChange,
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

  const handleButtonClick = () => {
    if (buttonAction) {
      buttonAction(leftValue, rightValue, setLeftValue, setRightValue);
    }
  };

  return (
    <div
      style={{
        padding: "10px",
        height: "calc(100vh - 40px)",
        boxSizing: "border-box",
      }}
    >
      <Row gutter={8} style={{ height: "calc(100% - 40px)", margin: 0 }}>
        <Col span={12} style={{ paddingRight: "8px", height: "100%" }}>
          <Editor
            height="calc(100% - 32px)"
            language={leftLanguage}
            value={leftValue}
            onTextChange={handleLeftChange}
            editorDidMount={(editor) => (leftEditorRef.current = editor)}
            readOnly={false} // Always writable
            languages={languages} // Pass languages to Editor component
            onLanguageChange={onLeftLangChange}
          />
        </Col>
        <Col span={12} style={{ paddingLeft: "8px", height: "100%" }}>
          <Editor
            height="calc(100% - 32px)"
            language={rightLanguage}
            value={rightValue}
            onTextChange={handleRightChange}
            onLanguageChange={onRightLangChange}
            editorDidMount={(editor) => (rightEditorRef.current = editor)}
            readOnly={true} // Always read-only
            languages={languages} // Pass languages to Editor component
          />
        </Col>
      </Row>
      <div style={{ textAlign: "center", marginTop: "16px" }}>
        {buttonLabel && (
          <Button
            type="primary"
            onClick={handleButtonClick}
            style={{ marginRight: "8px" }}
          >
            {buttonLabel}
          </Button>
        )}
      </div>
    </div>
  );
};

export default DualEditor;
