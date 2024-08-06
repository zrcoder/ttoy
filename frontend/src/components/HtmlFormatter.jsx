import DualEditor from "./DualEditor";
import { FormatHtml } from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const HtmlFormatter = () => {
  const transformAction = useTransformAction();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, FormatHtml, setRightValue);
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
