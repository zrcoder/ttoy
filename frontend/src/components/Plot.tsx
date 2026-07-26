import { Plot as SvcPlot } from "../../bindings/github.com/zrcoder/ttoy/service";
import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";

const Plot = () => (
  <AppTabs
    defaultActiveKey="d2"
    items={[
      {
        key: "d2",
        label: "D2",
        children: (
          <TextImager
            lang="python"
            imageGenerator={SvcPlot.D2}
            filename="d2.svg"
          />
        ),
      },
      {
        key: "ndor",
        label: "Ndor",
        children: (
          <TextImager
            lang="coffeescript"
            imageGenerator={SvcPlot.Ndor}
            filename="ndor.png"
          />
        ),
      },
    ]}
  />
);

export default Plot;
