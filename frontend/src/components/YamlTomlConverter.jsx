import React from "react";
import DualEditor from "./DualEditor";
import {
  ConvertYamlToToml,
  ConvertTomlToYaml,
} from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const YamlTomlConverter = () => {
  const transform = useTransformer();
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(leftValue, ConvertYamlToToml, setRightValue);
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(rightValue, ConvertTomlToYaml, setLeftValue);
  };

  return (
    <DualEditor
      leftLabel="yaml"
      rightLabel="toml"
      leftLanguage="yaml"
      rightLanguage="toml"
      firstButtonLabel="yaml → toml"
      firstButtonAction={handleFirstButtonAction}
      secondButtonLabel="yaml ← toml"
      secondButtonAction={handleSecondButtonAction}
      singleButton={false}
    />
  );
};

export default YamlTomlConverter;
