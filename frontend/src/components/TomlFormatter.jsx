import DualEditor from "./DualEditor";
import { FormatToml } from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const TomlFormatter = () => {
  const transformAction = useTransformAction();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, FormatToml, setRightValue);
  };
  return (
    <DualEditor
      leftLabel="input"
      leftLanguage="toml"
      rightLanguage="toml"
      rightLabel="output"
      rightReadOnly={true}
      firstButtonLabel="format"
      firstButtonAction={handleButtonAction}
      singleButton={true}
    />
  );
};

export default TomlFormatter;
