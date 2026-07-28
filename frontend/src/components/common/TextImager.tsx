import { CaretRightFilled } from "@ant-design/icons";
import type { OnMount } from "@monaco-editor/react";
import { App as AntdApp, Button, Spin } from "antd";
import { useRef, useState } from "react";
import Editor from "./Editor";
import { ImageWithDownload } from "./ImageWithDownload";
import { contentHeight } from "./layout";

type TextImagerProps = {
  editorContent?: string;
  onTextChange?: (text: string) => void;
  imageGenerator?: (input: string) => Promise<string>;
  lang: string;
  filename?: string;
};

const TextImager = ({
  editorContent,
  onTextChange,
  imageGenerator,
  lang,
  filename,
}: TextImagerProps) => {
  const { modal } = AntdApp.useApp();
  const [text, setText] = useState(editorContent ?? "");
  const [image, setImage] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const editorRef = useRef<Parameters<OnMount>[0] | null>(null);

  const handleTextChange = (value: string | undefined) => {
    const newText = value ?? "";
    setText(newText);
    if (onTextChange) onTextChange(newText);
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
        .catch((err: unknown) => {
          modal.error({ content: (err as Error).toString() });
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
        height: contentHeight,
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
          autoFocus
          editorDidMount={(editor) => {
            editorRef.current = editor;
          }}
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
        <Button
          onClick={handleButtonClick}
          loading={loading}
          icon={<CaretRightFilled />}
          type="primary"
        />
      </div>
      <div
        style={{
          flex: 1,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          background: "#FFFFFF20",
          overflow: "auto",
          position: "relative",
        }}
      >
        {loading && <Spin size="large" />}
        {image && (
          <ImageWithDownload src={image} alt="Generated" filename={filename} />
        )}
        {!image && !loading && <span style={{ color: "#aaa" }}>No image</span>}
      </div>
    </div>
  );
};

export default TextImager;
