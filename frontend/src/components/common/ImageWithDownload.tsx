import { DownloadOutlined } from "@ant-design/icons";
import { Button, message } from "antd";
import { Dialogs } from "@wailsio/runtime";
import React from "react";
import { Save as SvcSave } from "../../../bindings/github.com/zrcoder/ttoy/service";

export const ImageWithDownload: React.FC<{
  src: string;
  alt?: string;
  filename?: string;
}> = ({ src, alt = "image", filename }) => {
  if (!src) return null;

  const handleDownload = async () => {
    const suggestedName =
      filename || alt.replace(/[^a-zA-Z0-9.-]/g, "_") || "image.png";
    const filePath = await Dialogs.SaveFile({
      Title: "Save Image",
      Filename: suggestedName,
      Filters: [
        { DisplayName: "Images", Pattern: "*.png;*.jpg;*.jpeg;*.gif;*.webp" },
      ],
    });

    if (!filePath) return;

    try {
      await SvcSave.Image(filePath, src);
      message.success("Image saved successfully");
    } catch (err) {
      message.error(`Failed to save image: ${err}`);
    }
  };

  return (
    <div
      style={{
        position: "relative",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
        height: "100%",
      }}
    >
      <img
        src={src}
        alt={alt}
        style={{
          maxWidth: "100%",
          maxHeight: "100%",
          objectFit: "contain",
          userSelect: "none",
        }}
        onContextMenu={(e) => e.preventDefault()}
        draggable={false}
      />
      <Button
        type="primary"
        shape="circle"
        icon={<DownloadOutlined />}
        onClick={handleDownload}
        style={{
          position: "absolute",
          top: 16,
          right: 16,
          zIndex: 10,
        }}
      />
    </div>
  );
};
