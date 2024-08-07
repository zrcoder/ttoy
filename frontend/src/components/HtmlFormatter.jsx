import DualEditor from "./DualEditor";
import { FormatHtml } from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const HtmlFormatter = () => {
  const transform = useTransformer();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(leftValue, FormatHtml, setRightValue);
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
