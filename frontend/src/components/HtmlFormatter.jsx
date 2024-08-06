import DualEditor from "./DualEditor";
import { FormatHtml } from "../../wailsjs/go/main/App";

const HtmlFormatter = () => {
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    FormatHtml(leftValue).then((res) => {
      setRightValue(res);
    });
  };
  return (
    <DualEditor
      leftLabel="input"
      leftLanguage="html"
      rightLabel="output"
      rightLanguage="html"
      rightReadOnly={true}
      firstButtonLabel="format"
      firstButtonAction={handleButtonAction}
      singleButton={true}
    />
  );
};

export default HtmlFormatter;
