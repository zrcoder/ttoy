import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";
import { Plotter } from "../../bindings/github.com/zrcoder/ttoy/service/plot";

const Plot = () => (
  <AppTabs defaultActiveKey="d2" items={[
    { key: "d2", label: "D2", children: <TextImager lang="python" imageGenerator={Plotter.D2} /> },
    { key: "ndor", label: "Ndor", children: <TextImager lang="coffeescript" imageGenerator={Plotter.Ndor} /> },
  ]} />
);

export default Plot;
