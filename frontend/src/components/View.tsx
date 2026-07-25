import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";
import { Viewer } from "../../bindings/github.com/zrcoder/ttoy/service/view";

const View = () => (
  <AppTabs defaultActiveKey="1" items={[
    { key: "1", label: "JSON", children: <TextImager lang="json" imageGenerator={Viewer.Json} /> },
    { key: "2", label: "YAML", children: <TextImager lang="yaml" imageGenerator={Viewer.Yaml} /> },
    { key: "3", label: "TOML", children: <TextImager lang="toml" imageGenerator={Viewer.Toml} /> },
  ]} />
);

export default View;
