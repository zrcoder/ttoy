import { useState } from "react";
import { Button, Select, Spin, Input } from "antd";
import { CopyOutlined, CheckOutlined } from "@ant-design/icons";
import { App as AntdApp } from "antd";
import { contentHeight } from "./common/layout";
import {
  AsciiArt,
  Hash,
} from "../../bindings/github.com/zrcoder/ttoy/service/service";
import AppTabs from "./common/AppTabs";

const FONTS = ["big", "block", "chunky", "coinstak", "colossal", "cricket", "cyberlarge", "cybermedium", "doh", "doom", "isometric1", "isometric3", "larry3d", "marquee", "ogre", "pawp", "puffy", "rectangles", "rounded", "slant", "small", "standard", "starwars", "stop"];

type CopyButtonProps = {
  copied: boolean;
  onClick: () => void;
};

const CopyButton = ({ copied, onClick }: CopyButtonProps) => (
  <Button type="text" size="small" icon={copied ? <CheckOutlined style={{ color: "#52c41a" }} /> : <CopyOutlined />} onClick={onClick} />
);

type LabelRowProps = {
  label: string;
  copyButton: React.ReactNode;
};

const LabelRow = ({ label, copyButton }: LabelRowProps) => (
  <div style={{ display: "flex", alignItems: "center", gap: 8, marginBottom: 4 }}>
    <span style={{ fontWeight: "bold" }}>{label}</span>
    {copyButton}
  </div>
);

const AsciiArtTab = () => {
  const { modal } = AntdApp.useApp();
  const [input, setInput] = useState("");
  const [font, setFont] = useState("standard");
  const [output, setOutput] = useState("");
  const [loading, setLoading] = useState(false);
  const [copied, setCopied] = useState(false);

  const handleGenerate = () => {
    setCopied(false);
    if (!input.trim()) {
      modal.warning({ content: "Please input content first" });
      return;
    }
    setLoading(true);
    AsciiArt(input, font)
      .then((res) => {
        setOutput(res);
      })
      .catch((err: unknown) => {
        modal.error({ content: (err as Error).toString() });
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <div style={{ display: "flex", flexDirection: "column", height: contentHeight }}>
      <div style={{ marginBottom: 16 }}>
        <Input value={input} onChange={(e) => setInput(e.target.value.replace(/[^\x00-\x7F]/g, ""))} placeholder="Enter text (ASCII only)" />
      </div>
      <div style={{ display: "flex", gap: 8, marginBottom: 16, alignItems: "center" }}>
        <Select value={font} onChange={setFont} options={FONTS.map((f) => ({ value: f, label: f }))} style={{ width: 200 }} />
        <Button type="primary" onClick={handleGenerate} loading={loading}>Generate</Button>
      </div>
      <div style={{ flex: 1, minHeight: 0 }}>
        {loading ? <div style={{ display: "flex", justifyContent: "center", alignItems: "center", height: "100%" }}><Spin size="large" /></div> : output && (
          <>
            <LabelRow label="Output" copyButton={<CopyButton copied={copied} onClick={() => { navigator.clipboard.writeText(output); setCopied(true); }} />} />
            <Input.TextArea disabled value={output} style={{ fontFamily: "monospace", height: "96%" }} />
          </>
        )}
      </div>
    </div>
  );
};

type HashResults = Record<string, string> | null;

const HashTab = () => {
  const { modal } = AntdApp.useApp();
  const [input, setInput] = useState("");
  const [results, setResults] = useState<HashResults>(null);
  const [loading, setLoading] = useState(false);
  const [copiedKey, setCopiedKey] = useState<string | null>(null);

  const handleHash = () => {
    setCopiedKey(null);
    if (!input.trim()) {
      modal.warning({ content: "Please input content first" });
      return;
    }
    setLoading(true);
    Hash(input)
      .then((res) => {
        setResults(res as HashResults);
      })
      .catch((err: unknown) => {
        modal.error({ content: (err as Error).toString() });
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleCopy = (hashKey: string) => {
    setCopiedKey(hashKey);
  };

  type HashResultProps = {
    label: string;
    value: string | undefined;
    hashKey: string;
    onCopy: (key: string) => void;
  };

  const HashResult = ({ label, value, hashKey, onCopy }: HashResultProps) => (
    <div style={{ marginBottom: 16 }}>
      <LabelRow label={label} copyButton={<CopyButton copied={copiedKey === hashKey} onClick={() => { navigator.clipboard.writeText(value || ""); onCopy(hashKey); }} />} />
      <Input disabled value={value || "-"} style={{ fontFamily: "monospace" }} />
    </div>
  );

  return (
    <div style={{ display: "flex", flexDirection: "column", height: contentHeight }}>
      <div style={{ marginBottom: 16, height: 500 }}>
        <Input.TextArea value={input} onChange={(e) => setInput(e.target.value)} placeholder="Enter text" style={{ height: "100%" }} />
      </div>
      <Button type="primary" onClick={handleHash} loading={loading} style={{ marginBottom: 16, alignSelf: "flex-start" }}>Hash</Button>
      <div style={{ flex: 1, overflow: "auto" }}>
        {loading ? <div style={{ display: "flex", justifyContent: "center", alignItems: "center" }}><Spin size="large" /></div> : results && (
          ["MD5", "SHA1", "SHA256", "SHA512"].map((algo) => <HashResult key={algo} label={algo} value={results[algo.toLowerCase()]} hashKey={algo.toLowerCase()} onCopy={handleCopy} />)
        )}
      </div>
    </div>
  );
};

const Generator = () => (
  <AppTabs defaultActiveKey="1" items={[
    { key: "1", label: "AsciiArt", children: <AsciiArtTab /> },
    { key: "2", label: "Hash", children: <HashTab /> },
  ]} />
);

export default Generator;
