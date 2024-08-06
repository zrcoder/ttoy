import DualEditor from "./DualEditor";
import {
  ConvertJsonToToml,
  ConvertTomlToJson,
} from "../../wailsjs/go/main/App";

const JsonTomlConverter = () => {
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    ConvertJsonToToml(leftValue).then((res) => {
      setRightValue(res);
    });
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    ConvertTomlToJson(rightValue).then((res) => {
      setLeftValue(res);
    });
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
