import DualEditor from "./DualEditor";
import { FormatToml } from "../../wailsjs/go/main/App";

const TomlFormatter = () => {
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    FormatToml(leftValue).then((res) => {
      setRightValue(res);
    });
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
