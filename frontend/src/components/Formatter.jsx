import DualEditor from "./DualEditor";
import {
  FormatHtml,
  FormatJson,
  FormatToml,
  FormatYaml,
} from "../../bindings/github.com/zrcoder/ttoy/service/service";
import { Tabs } from "antd";
import { useTransformer } from "./util";

const Formatter = () => {
  const transform = useTransformer();

  const formatJson = (left, right, setL, setR) =>
    transform(left, FormatJson, setR);
  const formatYaml = (left, right, setL, setR) =>
    transform(left, FormatYaml, setR);
  const formatToml = (left, right, setL, setR) =>
    transform(left, FormatToml, setR);
  const formatHtml = (left, right, setL, setR) =>
    transform(left, FormatHtml, setR);

  const items = [
    {
      key: "1",
      label: "JSON",
      children: (
        <DualEditor
          leftLanguage={"json"}
          rightLanguage={"json"}
          buttonAction={formatJson}
        />
      ),
    },
    {
      key: "2",
      label: "YAML",
      children: (
        <DualEditor
          leftLanguage={"yaml"}
          rightLanguage={"yaml"}
          buttonAction={formatYaml}
        />
      ),
    },
    {
      key: "3",
      label: "TOML",
      children: (
        <DualEditor
          leftLanguage={"toml"}
          rightLanguage={"toml"}
          buttonAction={formatToml}
        />
      ),
    },
    {
      key: "4",
      label: "HTML",
      children: (
        <DualEditor
          leftLanguage={"html"}
          rightLanguage={"html"}
          buttonAction={formatHtml}
        />
      ),
    },
  ];

  return <Tabs defaultActiveKey="1" items={items} />;
};

export default Formatter;
