import { Format as SvcFormat } from "../../bindings/github.com/zrcoder/ttoy/service";
import AppTabs from "./common/AppTabs";
import DualEditor from "./common/DualEditor";
import { useTransformer } from "./common/useTransformer";

const Format = () => {
  const t = useTransformer();
  const fmt =
    (fn: (input: string) => Promise<string>) =>
    (
      _l: string,
      _r: string,
      _setL: (v: string) => void,
      setR: (v: string) => void,
    ) =>
      t(_l, fn, setR);
  const items = [
    {
      key: "1",
      label: "JSON",
      children: (
        <DualEditor
          leftLanguage="json"
          rightLanguage="json"
          buttonAction={fmt(SvcFormat.Json)}
        />
      ),
    },
    {
      key: "2",
      label: "YAML",
      children: (
        <DualEditor
          leftLanguage="yaml"
          rightLanguage="yaml"
          buttonAction={fmt(SvcFormat.Yaml)}
        />
      ),
    },
    {
      key: "3",
      label: "TOML",
      children: (
        <DualEditor
          leftLanguage="toml"
          rightLanguage="toml"
          buttonAction={fmt(SvcFormat.Toml)}
        />
      ),
    },
    {
      key: "4",
      label: "HTML",
      children: (
        <DualEditor
          leftLanguage="html"
          rightLanguage="html"
          buttonAction={fmt(SvcFormat.Html)}
        />
      ),
    },
  ];
  return <AppTabs defaultActiveKey="1" items={items} />;
};

export default Format;
