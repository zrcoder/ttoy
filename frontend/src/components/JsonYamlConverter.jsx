import React from "react";
import DualEditor from "./DualEditor";
import {
  ConvertJsonToYaml,
  ConvertYamlToJson,
} from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const JsonYamlConverter = () => {
  const transformAction = useTransformAction();
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, ConvertJsonToYaml, setRightValue);
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(rightValue, ConvertYamlToJson, setLeftValue);
  };

  return (
    <DualEditor
      leftLabel="json"
      rightLabel="yaml"
      leftLanguage="json"
      rightLanguage="yaml"
      firstButtonLabel="json → yaml"
      firstButtonAction={handleFirstButtonAction}
      secondButtonLabel="json ← yaml"
      secondButtonAction={handleSecondButtonAction}
      singleButton={false}
    />
  );
};

export default JsonYamlConverter;
