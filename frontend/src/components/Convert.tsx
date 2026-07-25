import DualEditor from "./common/DualEditor";
import { Converter } from "../../bindings/github.com/zrcoder/ttoy/service/convert";
import AppTabs from "./common/AppTabs";
import { useTransformer } from "./common/useTransformer";

const Convert = () => {
  const t = useTransformer();
  const cvt = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, _setL: (v: string) => void, setR: (v: string) => void) => t(_l, fn, setR);
  const rcv = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, setL: (v: string) => void, _setR: (v: string) => void) => t(_r, fn, setL);
  const items = [
    { key: "1", label: "JSON - YAML", children: <DualEditor leftLanguage="json" rightLanguage="yaml" buttonAction={cvt(Converter.Json2Yaml)} reverseButtonAction={rcv(Converter.Yaml2Json)} rightReadOnly={false} /> },
    { key: "2", label: "JSON - TOML", children: <DualEditor leftLanguage="json" rightLanguage="toml" buttonAction={cvt(Converter.Json2Toml)} reverseButtonAction={rcv(Converter.Toml2Json)} rightReadOnly={false} /> },
    { key: "3", label: "YAML - TOML", children: <DualEditor leftLanguage="yaml" rightLanguage="toml" buttonAction={cvt(Converter.Yaml2Toml)} reverseButtonAction={rcv(Converter.Toml2Yaml)} rightReadOnly={false} /> },
  ];
  return <AppTabs defaultActiveKey="1" items={items} />;
};

export default Convert;
