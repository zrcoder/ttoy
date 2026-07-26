import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";
import { Plot as SvcPlot } from "../../bindings/github.com/zrcoder/ttoy/service";

const Plot = () => (
  <AppTabs
    defaultActiveKey="d2"
    items={[
      {
        key: "d2",
        label: "D2",
        children: <TextImager lang="python" imageGenerator={SvcPlot.D2} />,
      },
      {
        key: "ndor",
        label: "Ndor",
        children: (
          <TextImager lang="coffeescript" imageGenerator={SvcPlot.Ndor} />
        ),
      },
    ]}
  />
);

export default Plot;
