import React from "react";
import DualEditor from "./DualEditor";
import {
  ConvertYamlToToml,
  ConvertTomlToYaml,
} from "../../wailsjs/go/main/App";

const YamlTomlConverter = () => {
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    ConvertYamlToToml(leftValue).then((res) => {
      setRightValue(res);
    });
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    ConvertTomlToYaml(rightValue).then((res) => {
      setLeftValue(res);
    });
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
