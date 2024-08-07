import React, { useState } from "react";
import TextImager from "./TextImager";
import { GenJsonSvg, GenTomlSvg, GenYamlSvg } from "../../wailsjs/go/main/App";

const DataViewer = () => {
  const [lang, setLang] = useState("json");
  const transformers = {
    "json>": GenJsonSvg,
    "yaml>": GenYamlSvg,
    "toml>": GenTomlSvg,
  };
  return (
    <TextImager
      buttonLabel="Generate→"
      langs={["json", "yaml", "toml"]}
      lang={lang}
      imageGenerator={transformers[lang + ">"]}
      onLangChange={(newValue) => {
        setLang(newValue);
      }}
    />
  );
};

export default DataViewer;
