import DualEditor from "./common/DualEditor";
import { Formatter } from "../../bindings/github.com/zrcoder/ttoy/service/format";
import AppTabs from "./common/AppTabs";
import { useTransformer } from "./common/useTransformer";

const Format = () => {
  const t = useTransformer();
  const fmt = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, _setL: (v: string) => void, setR: (v: string) => void) => t(_l, fn, setR);
  const items = [
    { key: "1", label: "JSON", children: <DualEditor leftLanguage="json" rightLanguage="json" buttonAction={fmt(Formatter.Json)} /> },
    { key: "2", label: "YAML", children: <DualEditor leftLanguage="yaml" rightLanguage="yaml" buttonAction={fmt(Formatter.Yaml)} /> },
    { key: "3", label: "TOML", children: <DualEditor leftLanguage="toml" rightLanguage="toml" buttonAction={fmt(Formatter.Toml)} /> },
    { key: "4", label: "HTML", children: <DualEditor leftLanguage="html" rightLanguage="html" buttonAction={fmt(Formatter.Html)} /> },
  ];
  return <AppTabs defaultActiveKey="1" items={items} />;
};

export default Format;
