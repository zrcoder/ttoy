import DualEditor from "./DualEditor";
import { FormatJson } from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const JsonFormatter = () => {
  const transformAction = useTransformAction();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, FormatJson, setRightValue);
  };
  return (
    <DualEditor
      leftLabel="input"
      leftLanguage="json"
      rightLabel="output"
      rightLanguage="json"
      leftReadOnly={false}
      rightReadOnly={true}
      firstButtonLabel="format"
      firstButtonAction={handleButtonAction}
      singleButton={true}
    />
  );
};

export default JsonFormatter;
