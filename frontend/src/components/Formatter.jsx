import DualEditor from "./DualEditor";
import {
  FormatHtml,
  FormatJson,
  FormatToml,
  FormatYaml,
} from "../../wailsjs/go/main/App";
import { useTransformer } from "./util";
import { useState } from "react";

const Formatter = () => {
  const [lang, setLang] = useState("json");
  const transformers = {
    "json>": FormatJson,
    "yaml>": FormatYaml,
    "toml>": FormatToml,
    "html>": FormatHtml,
  };
  const transform = useTransformer();
  const buttonAction = (leftValue, rightValue, setLeftValue, setRightValue) => {
    transform(leftValue, transformers[lang + ">"], setRightValue);
  };
  return (
    <DualEditor
      leftLanguage={lang}
      rightLanguage={lang}
      buttonLabel="Format→"
      buttonAction={buttonAction}
      languages={["json", "yaml", "toml", "html"]}
      onLeftLangChange={(newValue) => setLang(newValue)}
      onRightLangChange={(newValue) => setLang(newValue)}
    />
  );
};

export default Formatter;
