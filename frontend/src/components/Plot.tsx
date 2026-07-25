import AppTabs from "./common/AppTabs";
import TextImager from "./common/TextImager";
import { GenD2Svg, GenNdorPng } from "../../bindings/github.com/zrcoder/ttoy/service/service";

const Plot = () => (
  <AppTabs defaultActiveKey="d2" items={[
    { key: "d2", label: "D2", children: <TextImager lang="python" imageGenerator={GenD2Svg} /> },
    { key: "ndor", label: "Ndor", children: <TextImager lang="coffeescript" imageGenerator={GenNdorPng} /> },
  ]} />
);

export default Plot;