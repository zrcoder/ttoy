import DualEditor from "./DualEditor";
import { FormatJson } from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const JsonFormatter = () => {
  const transform = useTransformer();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(leftValue, FormatJson, setRightValue);
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
