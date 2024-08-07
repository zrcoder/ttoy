import React from "react";
import TextImageGenerator from "./TextImageGenerator";
import { GenJsonSvg } from "../../wailsjs/go/main/App";

const JsonSvgViewer = () => {
  return (
    <TextImageGenerator
      editorLabel="Json"
      imageLabel="Graph"
      editorLanguage="json"
      generator={GenJsonSvg}
      buttonLabel="Generate →"
      imageType="svg"
    />
  );
};

export default JsonSvgViewer;
