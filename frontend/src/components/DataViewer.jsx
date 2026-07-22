import React from "react";
import { Tabs } from "antd";
import TextImager from "./TextImager";
import {
  GenJsonSvg,
  GenTomlSvg,
  GenYamlSvg,
} from "../../bindings/github.com/zrcoder/ttoy/service/service";

const DataViewer = () => {
  const items = [
    {
      key: "1",
      label: "JSON",
      children: <TextImager lang="json" imageGenerator={GenJsonSvg} />,
    },
    {
      key: "2",
      label: "YAML",
      children: <TextImager lang="yaml" imageGenerator={GenYamlSvg} />,
    },
    {
      key: "3",
      label: "TOML",
      children: <TextImager lang="toml" imageGenerator={GenTomlSvg} />,
    },
  ];

  return <Tabs defaultActiveKey="1" items={items} />;
};

export default DataViewer;
