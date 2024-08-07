import React, { useState, useRef } from "react";
import { Button, Col, Row, Spin } from "antd";
import Editor from "./Editor";
import { App as AntdApp } from "antd";

const TextImager = ({
  langs = [],
  editorContent,
  onTextChange,
  imageGenerator, // Function to generate image based on text
  buttonLabel = "Generate→",
  lang,
  onLangChange,
}) => {
  const { modal } = AntdApp.useApp();
  const [text, setText] = useState(editorContent || "");
  const [image, setImage] = useState(null);
  const [loading, setLoading] = useState(false);
  const editorRef = useRef(null);

  const handleTextChange = (value) => {
    setText(value || "");
    if (onTextChange) onTextChange(value || "");
  };

  const handleButtonClick = () => {
    setLoading(true);
    setImage(null);
    if (imageGenerator) {
      imageGenerator(text)
        .then((res) => {
          setImage(res);
        })
        .catch((err) => {
          modal.error({
            content: err.toString(),
          });
        })
        .finally(() => {
          setLoading(false);
        });
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
          <div
            style={{ height: "100%", display: "flex", flexDirection: "column" }}
          >
            <Editor
              height="calc(100% - 32px)" // Adjust height based on the label
              language={lang}
              languages={langs}
              onLanguageChange={onLangChange}
              value={text}
              onChange={handleTextChange}
              editorDidMount={(editor) => (editorRef.current = editor)}
            />
          </div>
        </Col>
        <Col span={12} style={{ paddingLeft: "8px", height: "100%" }}>
          <div
            style={{ height: "100%", display: "flex", flexDirection: "column" }}
          >
            <div
              style={{
                height: "calc(100% - 32px)",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                position: "relative", // Make the container relative for positioning the spinner
                overflow: "hidden", // Hide overflow to avoid scrollbars
              }}
            >
              {loading && (
                <div
                  style={{
                    position: "absolute",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                    width: "100%",
                    height: "100%",
                    backgroundColor: "rgba(0, 0, 0, 0)", // Transparent background
                    zIndex: 1, // Ensure spinner is on top of image
                  }}
                >
                  <Spin size="large" />
                </div>
              )}
              {image ? (
                <img
                  src={image}
                  alt="Generated"
                  style={{
                    maxWidth: "100%",
                    maxHeight: "100%",
                    objectFit: "contain", // Ensures the image keeps its aspect ratio
                    zIndex: 0, // Ensure image is behind the spinner
                  }}
                />
              ) : (
                !loading && (
                  <div style={{ textAlign: "center", color: "#aaa" }}>
                    No image
                  </div>
                )
              )}
            </div>
          </div>
        </Col>
      </Row>
      <div style={{ textAlign: "center", marginTop: "16px" }}>
        <Button type="primary" onClick={handleButtonClick}>
          {buttonLabel}
        </Button>
      </div>
    </div>
  );
};

export default TextImager;
