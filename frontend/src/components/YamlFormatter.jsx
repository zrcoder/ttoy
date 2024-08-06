import DualEditor from "./DualEditor";
import { FormatYaml } from "../../wailsjs/go/main/App";

const YamlFormatter = () => {
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    FormatYaml(leftValue).then((res) => {
      setRightValue(res);
    });
  };
  return (
    <DualEditor
      leftLabel="input"
      leftLanguage="yaml"
      rightLabel="output"
      rightLanguage="yaml"
      firstButtonLabel="format"
      rightReadOnly={true}
      firstButtonAction={handleButtonAction}
      singleButton={true}
    />
  );
};

export default YamlFormatter;
