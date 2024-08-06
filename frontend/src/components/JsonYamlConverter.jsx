import React from "react";
import DualEditor from "./DualEditor";
import {
  ConvertJsonToYaml,
  ConvertYamlToJson,
} from "../../wailsjs/go/main/App";

const JsonYamlConverter = () => {
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    ConvertJsonToYaml(leftValue).then((res) => {
      setRightValue(res);
    });
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    ConvertYamlToJson(rightValue).then((res) => {
      setLeftValue(res);
    });
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
