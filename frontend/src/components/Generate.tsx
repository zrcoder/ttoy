import { App as AntdApp, Button, Input, Select, Spin } from "antd";
import { useState } from "react";
import { Generate as SvcGenerate } from "../../bindings/github.com/zrcoder/ttoy/service";
import AppTabs from "./common/AppTabs";
import { CopyButton } from "./common/CopyButton";
import { contentHeight } from "./common/layout";

const FONTS = [
  "big",
  "block",
  "chunky",
  "coinstak",
  "colossal",
  "cricket",
  "cyberlarge",
  "cybermedium",
  "doh",
  "doom",
  "isometric1",
  "isometric3",
  "larry3d",
  "marquee",
  "ogre",
  "pawp",
  "puffy",
  "rectangles",
  "rounded",
  "slant",
  "small",
  "standard",
  "starwars",
  "stop",
];

type LabelRowProps = {
  label: string;
  copyButton: React.ReactNode;
};

const LabelRow = ({ label, copyButton }: LabelRowProps) => (
  <div
    style={{ display: "flex", alignItems: "center", gap: 8, marginBottom: 4 }}
  >
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

  const handleGenerate = () => {
    if (!input.trim()) {
      modal.warning({ content: "Please input content first" });
      return;
    }
    setLoading(true);
    SvcGenerate.AsciiArt(input, font)
      .then((res) => {
        setOutput(res ?? "");
      })
      .catch((err: unknown) => {
        modal.error({ content: (err as Error).toString() });
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: contentHeight,
      }}
    >
      <div style={{ marginBottom: 16 }}>
        <Input
          value={input}
          onChange={(e) =>
            setInput(e.target.value.replace(/[^\x00-\x7F]/g, ""))
          }
          placeholder="Enter text (ASCII only)"
        />
      </div>
      <div
        style={{
          display: "flex",
          gap: 8,
          marginBottom: 16,
          alignItems: "center",
        }}
      >
        <Select
          value={font}
          onChange={setFont}
          options={FONTS.map((f) => ({ value: f, label: f }))}
          style={{ width: 200 }}
        />
        <Button type="primary" onClick={handleGenerate} loading={loading}>
          Generate
        </Button>
      </div>
      <div style={{ flex: 1, minHeight: 0, position: "relative" }}>
        <CopyButton text={output} />
        <Input.TextArea
          value={output}
          readOnly
          style={{ height: "100%", fontFamily: "monospace" }}
        />
      </div>
    </div>
  );
};

type HashResults = Record<string, string>;

const HashTab = () => {
  const { modal } = AntdApp.useApp();
  const [input, setInput] = useState("");
  const [results, setResults] = useState<HashResults | null>(null);
  const [loading, setLoading] = useState(false);

  const handleHash = () => {
    if (!input.trim()) {
      modal.warning({ content: "Please input content first" });
      return;
    }
    setLoading(true);
    SvcGenerate.Hash(input)
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

  type HashResultProps = {
    label: string;
    value: string;
  };

  const HashResult = ({ label, value }: HashResultProps) => (
    <div style={{ marginBottom: 16 }}>
      <LabelRow label={label} copyButton={<CopyButton text={value} inline />} />
      <Input value={value || "-"} style={{ fontFamily: "monospace" }} />
    </div>
  );

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: contentHeight,
      }}
    >
      <div style={{ marginBottom: 16, height: 500 }}>
        <Input.TextArea
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Enter text"
          style={{ height: "100%" }}
        />
      </div>
      <Button
        type="primary"
        onClick={handleHash}
        loading={loading}
        style={{ marginBottom: 16, alignSelf: "flex-start" }}
      >
        Hash
      </Button>
      <div style={{ flex: 1, overflow: "auto" }}>
        {loading ? (
          <div
            style={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
            }}
          >
            <Spin size="large" />
          </div>
        ) : (
          results &&
          ["MD5", "SHA1", "SHA256", "SHA512"].map((algo) => (
            <HashResult
              key={algo}
              label={algo}
              value={results[algo.toLowerCase()]}
            />
          ))
        )}
      </div>
    </div>
  );
};

const Generator = () => (
  <AppTabs
    defaultActiveKey="1"
    items={[
      { key: "1", label: "AsciiArt", children: <AsciiArtTab /> },
      { key: "2", label: "Hash", children: <HashTab /> },
    ]}
  />
);

export default Generator;
