import React, { useState, useRef } from "react";
import { Button, Col, Row, Spin } from "antd";
import Editor from "./Editor";
import { App as AntdApp } from "antd";

const TextImageGenerator = ({
  editorLabel,
  imageLabel,
  editorLanguage,
  editorContent,
  onEditorChange,
  generator, // Function to generate image based on text
  buttonLabel = "Generate",
}) => {
  const { modal } = AntdApp.useApp();
  const [text, setText] = useState(editorContent || "");
  const [image, setImage] = useState(null);
  const [loading, setLoading] = useState(false);
  const editorRef = useRef(null);

  const handleTextChange = (value) => {
    setText(value || "");
    if (onEditorChange) onEditorChange(value || "");
  };

  const handleButtonClick = () => {
    setLoading(true); // Start loading indicator
    setImage(null);

    if (generator) {
      generator(text)
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
        padding: "20px",
        height: "calc(100vh - 40px)",
        boxSizing: "border-box",
      }}
    >
      <Row gutter={8} style={{ height: "calc(100% - 40px)", margin: 0 }}>
        <Col span={12} style={{ paddingRight: "8px", height: "100%" }}>
          <div
            style={{ height: "100%", display: "flex", flexDirection: "column" }}
          >
            <div style={{ padding: "8px", borderBottom: "1px solid #d9d9d9" }}>
              <h4 style={{ margin: 0 }}>{editorLabel}</h4>
            </div>
            <Editor
              height="calc(100% - 32px)" // Adjust height based on the label
              language={editorLanguage}
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
            <div style={{ padding: "8px", borderBottom: "1px solid #d9d9d9" }}>
              <h4 style={{ margin: 0 }}>{imageLabel}</h4>
            </div>
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

export default TextImageGenerator;
