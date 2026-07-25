import { useState } from "react";
import { Layout, Menu, ConfigProvider, theme, App as AntdApp } from "antd";
// @ts-ignore
import "antd/dist/reset.css";
import {
  ToolOutlined,
  SyncOutlined,
  FormatPainterOutlined,
  DiffOutlined,
  ThunderboltOutlined,
  DotChartOutlined,
  NodeIndexOutlined,
  LockOutlined,
} from "@ant-design/icons";
import Home from "./components/Home";
import Format from "./components/Format";
import Convert from "./components/Convert";
import Encode from "./components/Encode";
import Generate from "./components/Generate";
import View from "./components/View";
import Plot from "./components/Plot";
import Diff from "./components/Diff";
import { ThemeProvider, useTheme } from "./contexts/ThemeContext";
import ThemeSwitch from "./components/common/ThemeSwitch";

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
  | "hanoi";

type MenuItem = { key: MenuKey; icon?: React.ReactNode; label: string };
type MenuGroup = { key: string; groupTitle: string; items: MenuItem[] };

const toolsItems: MenuItem[] = [
  { key: "fmt", icon: <FormatPainterOutlined />, label: "Format" },
  { key: "cvt", icon: <SyncOutlined />, label: "Convert" },
  { key: "codec", icon: <LockOutlined />, label: "Encode" },
  { key: "gen", icon: <ThunderboltOutlined />, label: "Generate" },
  { key: "view", icon: <DotChartOutlined />, label: "View" },
  { key: "graph", icon: <NodeIndexOutlined />, label: "Plot" },
  { key: "diff", icon: <DiffOutlined />, label: "Diff" },
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
  hanoi: "Hi",
};

const menuItems: Array<MenuItem | MenuGroup> = [
  { key: "home", icon: <ToolOutlined />, label: "TToy" },
  { key: "tools", groupTitle: "Tools", items: toolsItems },
  {
    key: "games",
    groupTitle: "Games",
    items: [{ key: "hanoi", label: "Hanoi" }],
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
          </Sider>
          <Layout>
            <Content style={{ padding: "0 15px 15px", position: "relative" }}>
              <div
                style={{ position: "absolute", top: 0, right: 15, zIndex: 1 }}
              >
                <ThemeSwitch />
              </div>
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
