import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";
import { GenJsonSvg, GenTomlSvg, GenYamlSvg } from "../../bindings/github.com/zrcoder/ttoy/service/service";

const View = () => (
  <AppTabs defaultActiveKey="1" items={[
    { key: "1", label: "JSON", children: <TextImager lang="json" imageGenerator={GenJsonSvg} /> },
    { key: "2", label: "YAML", children: <TextImager lang="yaml" imageGenerator={GenYamlSvg} /> },
    { key: "3", label: "TOML", children: <TextImager lang="toml" imageGenerator={GenTomlSvg} /> },
  ]} />
);

export default View;
