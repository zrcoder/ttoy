import DualEditor from "./common/DualEditor";
import { Encoder } from "../../bindings/github.com/zrcoder/ttoy/service/encode";
import AppTabs from "./common/AppTabs";
import { useTransformer } from "./common/useTransformer";

const Encode = () => {
  const t = useTransformer();
  const enc = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, _setL: (v: string) => void, setR: (v: string) => void) => t(_l, fn, setR);
  const dec = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, setL: (v: string) => void, _setR: (v: string) => void) => t(_r, fn, setL);
  const items = [
    { key: "1", label: "Base64", children: <DualEditor leftLanguage="text" rightLanguage="text" buttonAction={enc(Encoder.Base64Encode)} reverseButtonAction={dec(Encoder.Base64Decode)} rightReadOnly={false} /> },
    { key: "2", label: "HTML", children: <DualEditor leftLanguage="html" rightLanguage="html" buttonAction={enc(Encoder.HTMLEncode)} reverseButtonAction={dec(Encoder.HTMLDecode)} rightReadOnly={false} /> },
    { key: "3", label: "URL", children: <DualEditor leftLanguage="text" rightLanguage="text" buttonAction={enc(Encoder.URLEncode)} reverseButtonAction={dec(Encoder.URLDecode)} rightReadOnly={false} /> },
  ];
  return <AppTabs defaultActiveKey="1" items={items} />;
};

export default Encode;
