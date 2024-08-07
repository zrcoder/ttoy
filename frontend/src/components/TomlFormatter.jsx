import DualEditor from "./DualEditor";
import { FormatToml } from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const TomlFormatter = () => {
  const transform = useTransformer();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(leftValue, FormatToml, setRightValue);
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
