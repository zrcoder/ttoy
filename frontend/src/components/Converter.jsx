import React, { useState } from "react";
import DualEditor from "./DualEditor";
import {
  ConvertJsonToToml,
  ConvertJsonToYaml,
  ConvertTomlToJson,
  ConvertTomlToYaml,
  ConvertYamlToJson,
  ConvertYamlToToml,
  FormatJson,
  FormatToml,
  FormatYaml,
} from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const Converter = () => {
  const [leftLang, setLeftLang] = useState("json");
  const [rightLang, setRightLang] = useState("yaml");
  const transform = useTransformer();
  const transformers = {
    "json>yaml": ConvertJsonToYaml,
    "json>toml": ConvertJsonToToml,
    "yaml>json": ConvertYamlToJson,
    "yaml>toml": ConvertYamlToToml,
    "toml>json": ConvertTomlToJson,
    "toml>yaml": ConvertTomlToYaml,
    "json>json": FormatJson,
    "yaml>yaml": FormatYaml,
    "toml>toml": FormatToml,
  };
  const buttonAction = (leftValue, rightValue, setLeftValue, setRightValue) => {
    transform(
      leftValue,
      transformers[leftLang + ">" + rightLang],
      setRightValue
    );
  };

  return (
    <DualEditor
      leftLanguage={leftLang}
      rightLanguage={rightLang}
      languages={["json", "yaml", "toml"]}
      buttonLabel="Convert→" // ←
      buttonAction={buttonAction}
      onLeftLangChange={(newVal) => {
        setLeftLang(newVal);
      }}
      onRightLangChange={(newVal) => {
        setRightLang(newVal);
      }}
    />
  );
};

export default Converter;
