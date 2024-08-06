import React from "react";
import DualEditor from "./DualEditor";
import {
  ConvertYamlToToml,
  ConvertTomlToYaml,
} from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const YamlTomlConverter = () => {
  const transformAction = useTransformAction();
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, ConvertYamlToToml, setRightValue);
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(rightValue, ConvertTomlToYaml, setLeftValue);
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
