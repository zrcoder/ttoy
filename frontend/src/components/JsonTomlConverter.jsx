import DualEditor from "./DualEditor";
import {
  ConvertJsonToToml,
  ConvertTomlToJson,
} from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const JsonTomlConverter = () => {
  const transform = useTransformer();
  const handleFirstButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(leftValue, ConvertJsonToToml, setRightValue);
  };

  const handleSecondButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(rightValue, ConvertTomlToJson, setLeftValue);
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
