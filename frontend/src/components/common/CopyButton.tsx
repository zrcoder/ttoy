import { CopyOutlined } from "@ant-design/icons";
import { Button, Tooltip } from "antd";
import React, { useRef, useState } from "react";

export const CopyButton: React.FC<{ text: string; inline?: boolean }> = ({
  text,
  inline,
}) => {
  const [open, setOpen] = useState(false);
  const timeoutRef = useRef<ReturnType<typeof setTimeout>>(null);

  if (!text) return null;

  const handleClick = async () => {
    try {
      await navigator.clipboard.writeText(text);
      setOpen(true);
      if (timeoutRef.current) clearTimeout(timeoutRef.current);
      timeoutRef.current = setTimeout(() => setOpen(false), 1000);
    } catch {}
  };

  const wrapperStyle: React.CSSProperties = inline
    ? { display: "inline-flex", verticalAlign: "middle" }
    : { position: "absolute", top: 8, right: 8, zIndex: 10 };

  return (
    <div style={wrapperStyle}>
      <Tooltip open={open} title="Copied!" placement="top">
        <Button
          size="small"
          icon={<CopyOutlined />}
          onClick={handleClick}
        />{" "}
      </Tooltip>
    </div>
  );
};
