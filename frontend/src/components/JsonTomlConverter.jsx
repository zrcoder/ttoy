import DualEditor from "./DualEditor";
import {
  ConvertJsonToToml,
  ConvertTomlToJson,
} from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const JsonTomlConverter = () => {
  const transformAction = useTransformAction();
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, ConvertJsonToToml, setRightValue);
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(rightValue, ConvertTomlToJson, setLeftValue);
  };

  return (
    <DualEditor
      leftLabel="json"
      rightLabel="toml"
      leftLanguage="json"
      rightLanguage="toml"
      firstButtonLabel="json → toml"
      firstButtonAction={handleFirstButtonAction}
      secondButtonLabel="json ← toml"
      secondButtonAction={handleSecondButtonAction}
      singleButton={false}
    />
  );
};

export default JsonTomlConverter;
