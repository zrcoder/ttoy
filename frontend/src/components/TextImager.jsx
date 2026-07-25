import React, { useState, useRef } from "react";
import { Button, Spin } from "antd";
import Editor from "./Editor";
import { App as AntdApp } from "antd";
import { CaretRightFilled } from "@ant-design/icons";
const TextImager = ({
  editorContent,
  onTextChange,
  imageGenerator,
  lang,
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
    if (!text.trim()) {
      modal.warning({ content: "Please input content first" });
      return;
    }
    setLoading(true);
    setImage(null);
    if (imageGenerator) {
      imageGenerator(text)
        .then((res) => {
          setImage(res);
        })
        .catch((err) => {
          modal.error({ content: err.toString() });
        })
        .finally(() => {
          setLoading(false);
        });
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
          language={lang}
          value={text}
          onTextChange={handleTextChange}
          editorDidMount={(editor) => (editorRef.current = editor)}
        />
      </div>
      <div
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          padding: "0 8px",
        }}
      >
        <Button onClick={handleButtonClick} loading={loading} icon={<CaretRightFilled />}>
        </Button>
      </div>
      <div
        style={{
          flex: 1,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          background: "transparent",
          overflow: "auto",
          position: "relative",
        }}
      >
        {loading && <Spin size="large" />}
        {image && (
          <img
            src={image}
            alt="Generated"
            style={{
              maxWidth: "100%",
              maxHeight: "100%",
              objectFit: "contain",
            }}
          />
        )}
        {!image && !loading && <span style={{ color: "#aaa" }}>No image</span>}
      </div>
    </div>
  );
};

export default TextImager;
