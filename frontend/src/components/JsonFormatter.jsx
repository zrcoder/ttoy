import DualEditor from "./DualEditor";
import { FormatJson } from "../../wailsjs/go/main/App";

const JsonFormatter = () => {
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    FormatJson(leftValue).then((res) => {
      setRightValue(res);
    });
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
