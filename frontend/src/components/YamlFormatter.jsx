import DualEditor from "./DualEditor";
import { FormatYaml } from "../../wailsjs/go/main/App";
import useTransformAction from "./useTransform";

const YamlFormatter = () => {
  const transformAction = useTransformAction();
  const handleButtonAction = (
    leftValue,
    rightValue,
    setLeftValue,
    setRightValue
  ) => {
    transformAction(leftValue, FormatYaml, setRightValue);
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
