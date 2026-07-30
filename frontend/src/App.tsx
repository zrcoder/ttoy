import {
  DiffOutlined,
  DotChartOutlined,
  FireOutlined,
  FormatPainterOutlined,
  LockOutlined,
  NodeIndexOutlined,
  SyncOutlined,
  ThunderboltOutlined,
  ToolOutlined,
} from "@ant-design/icons";
import { App as AntdApp, ConfigProvider, Layout, Menu, theme } from "antd";
import { useState } from "react";
// @ts-ignore
import "antd/dist/reset.css";
import ThemeSwitch from "./components/common/ThemeSwitch";
import Convert from "./components/Convert";
import Diff from "./components/Diff";
import Encode from "./components/Encode";
import Format from "./components/Format";
import Generate from "./components/Generate";
import Home from "./components/Home";
import Plot from "./components/Plot";
import View from "./components/View";
import { ThemeProvider, useTheme } from "./contexts/ThemeContext";
import IceMagic from "./games/IceMagic";

const { Sider, Content } = Layout;

type MenuKey =
  | "home"
  | "fmt"
  | "cvt"
  | "codec"
  | "gen"
  | "view"
  | "graph"
  | "diff"
  | "icemagic";

type MenuItem = { key: MenuKey; icon?: React.ReactNode; label: string };
type MenuGroup = { key: string; groupTitle: string; items: MenuItem[] };

const toolItems: MenuItem[] = [
  { key: "fmt", icon: <FormatPainterOutlined />, label: "Format" },
  { key: "cvt", icon: <SyncOutlined />, label: "Convert" },
  { key: "codec", icon: <LockOutlined />, label: "Encode" },
  { key: "gen", icon: <ThunderboltOutlined />, label: "Generate" },
  { key: "view", icon: <DotChartOutlined />, label: "View" },
  { key: "graph", icon: <NodeIndexOutlined />, label: "Plot" },
  { key: "diff", icon: <DiffOutlined />, label: "Diff" },
];

const gameItems: MenuItem[] = [
  { key: "icemagic", icon: <FireOutlined />, label: "IceMagic" },
];

const componentMap: Record<MenuKey, React.ReactNode> = {
  home: <Home />,
  fmt: <Format />,
  cvt: <Convert />,
  codec: <Encode />,
  gen: <Generate />,
  diff: <Diff />,
  view: <View />,
  graph: <Plot />,
  icemagic: <IceMagic />,
};

const menuItems: Array<MenuItem | MenuGroup> = [
  { key: "home", icon: <ToolOutlined />, label: "TToy" },
  { key: "tools", groupTitle: "Tools", items: toolItems },
  {
    key: "games",
    groupTitle: "Games",
    items: gameItems,
  },
];

const renderMenuItems = (items: typeof menuItems) =>
  items.map((item) =>
    "items" in item ? (
      <Menu.ItemGroup key={item.key} title={item.groupTitle}>
        {item.items.map((i) => (
          <Menu.Item key={i.key} icon={i.icon}>
            {i.label}
          </Menu.Item>
        ))}
      </Menu.ItemGroup>
    ) : (
      <Menu.Item key={item.key} icon={item.icon}>
        {item.label}
      </Menu.Item>
    ),
  );

const AppContent = () => {
  const { isDark } = useTheme();
  const [selectedKey, setSelectedKey] = useState<MenuKey>("home");

  return (
    <ConfigProvider
      theme={{
        algorithm: isDark ? theme.darkAlgorithm : theme.defaultAlgorithm,
      }}
    >
      <AntdApp>
        <Layout style={{ minHeight: "100vh", paddingTop: 46 }}>
          <Sider>
            <Menu
              selectedKeys={[selectedKey]}
              onClick={({ key }) => setSelectedKey(key as MenuKey)}
              style={{ height: "100%", borderRight: 0, textAlign: "left" }}
            >
              {renderMenuItems(menuItems)}
            </Menu>
            <div
              style={{
                position: "absolute",
                bottom: 16,
                left: 0,
                right: 0,
                padding: "0 16px",
              }}
            >
              <div style={{ width: "100%" }}>
                <ThemeSwitch />
              </div>
            </div>
          </Sider>
          <Layout>
            <Content style={{ padding: "0 15px 15px" }}>
              {componentMap[selectedKey]}
            </Content>
          </Layout>
        </Layout>
      </AntdApp>
    </ConfigProvider>
  );
};

const App = () => (
  <ThemeProvider>
    <AppContent />
  </ThemeProvider>
);

export default App;
