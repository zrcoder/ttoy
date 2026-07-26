import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";
import { View as SvcView } from "../../bindings/github.com/zrcoder/ttoy/service";

const View = () => (
  <AppTabs
    defaultActiveKey="1"
    items={[
      {
        key: "1",
        label: "JSON",
        children: <TextImager lang="json" imageGenerator={SvcView.Json} />,
      },
      {
        key: "2",
        label: "YAML",
        children: <TextImager lang="yaml" imageGenerator={SvcView.Yaml} />,
      },
      {
        key: "3",
        label: "TOML",
        children: <TextImager lang="toml" imageGenerator={SvcView.Toml} />,
      },
    ]}
  />
);

export default View;
