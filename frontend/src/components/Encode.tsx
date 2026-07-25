import DualEditor from "./common/DualEditor";
import { DecodeBase64, DecodeHtml, DecodeUrl, EncodeBase64, EncodeHtml, EncodeUrl } from "../../bindings/github.com/zrcoder/ttoy/service/service";
import AppTabs from "./common/AppTabs";
import { useTransformer } from "./common/useTransformer";

const Encode = () => {
  const t = useTransformer();
  const enc = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, _setL: (v: string) => void, setR: (v: string) => void) => t(_l, fn, setR);
  const dec = (fn: (input: string) => Promise<string>) => (_l: string, _r: string, setL: (v: string) => void, _setR: (v: string) => void) => t(_r, fn, setL);
  const items = [
    { key: "1", label: "Base64", children: <DualEditor leftLanguage="text" rightLanguage="text" buttonAction={enc(EncodeBase64)} reverseButtonAction={dec(DecodeBase64)} rightReadOnly={false} /> },
    { key: "2", label: "HTML", children: <DualEditor leftLanguage="html" rightLanguage="html" buttonAction={enc(EncodeHtml)} reverseButtonAction={dec(DecodeHtml)} rightReadOnly={false} /> },
    { key: "3", label: "URL", children: <DualEditor leftLanguage="text" rightLanguage="text" buttonAction={enc(EncodeUrl)} reverseButtonAction={dec(DecodeUrl)} rightReadOnly={false} /> },
  ];
  return <AppTabs defaultActiveKey="1" items={items} />;
};

export default Encode;
