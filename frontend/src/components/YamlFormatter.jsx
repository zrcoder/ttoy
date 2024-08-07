import DualEditor from "./DualEditor";
import { FormatYaml } from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";

const YamlFormatter = () => {
  const transform = useTransformer();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transform(leftValue, FormatYaml, setRightValue);
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
