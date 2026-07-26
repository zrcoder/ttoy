import { View as SvcView } from "../../bindings/github.com/zrcoder/ttoy/service";
import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";

const View = () => (
  <AppTabs
    defaultActiveKey="1"
    items={[
      {
        key: "1",
        label: "JSON",
        children: (
          <TextImager
            lang="json"
            imageGenerator={SvcView.Json}
            filename="view.json.svg"
          />
        ),
      },
      {
        key: "2",
        label: "YAML",
        children: (
          <TextImager
            lang="yaml"
            imageGenerator={SvcView.Yaml}
            filename="view.yaml.svg"
          />
        ),
      },
      {
        key: "3",
        label: "TOML",
        children: (
          <TextImager
            lang="toml"
            imageGenerator={SvcView.Toml}
            filename="view.toml.svg"
          />
        ),
      },
    ]}
  />
);

export default View;
