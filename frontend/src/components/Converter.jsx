import DualEditor from "./DualEditor";
import {
  ConvertJsonToToml,
  ConvertJsonToYaml,
  ConvertTomlToJson,
  ConvertTomlToYaml,
  ConvertYamlToJson,
  ConvertYamlToToml,
} from "../../bindings/github.com/zrcoder/ttoy/service/service";
import { Tabs } from 'antd';
import { useTransformer } from "./util";

const Converter = () => {
  const transform = useTransformer();

  const cvtJsonToYaml = (left, right, setL, setR) => transform(left, ConvertJsonToYaml, setR);
  const cvtYamlToJson = (left, right, setL, setR) => transform(right, ConvertYamlToJson, setL);
  const cvtJsonToToml = (left, right, setL, setR) => transform(left, ConvertJsonToToml, setR);
  const cvtTomlToJson = (left, right, setL, setR) => transform(right, ConvertTomlToJson, setL);
  const cvtYamlToToml = (left, right, setL, setR) => transform(left, ConvertYamlToToml, setR);
  const cvtTomlToYaml = (left, right, setL, setR) => transform(right, ConvertTomlToYaml, setL);

  const items = [
    {
      key: '1',
      label: 'JSON ↔ YAML',
      children: (
        <DualEditor
          leftLanguage={'json'}
          rightLanguage={'yaml'}
          buttonAction={cvtJsonToYaml}
          reverseButtonAction={cvtYamlToJson}
          rightReadOnly={false}
        />
      ),
    },
    {
      key: '2',
      label: 'JSON ↔ TOML',
      children: (
        <DualEditor
          leftLanguage={'json'}
          rightLanguage={'toml'}
          buttonAction={cvtJsonToToml}
          reverseButtonAction={cvtTomlToJson}
          rightReadOnly={false}
        />
      ),
    },
    {
      key: '3',
      label: 'YAML ↔ TOML',
      children: (
        <DualEditor
          leftLanguage={'yaml'}
          rightLanguage={'toml'}
          buttonAction={cvtYamlToToml}
          reverseButtonAction={cvtTomlToYaml}
          rightReadOnly={false}
        />
      ),
    },
  ];

  return (
    <Tabs defaultActiveKey="1" items={items} />
  );
};

export default Converter;
