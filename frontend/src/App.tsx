import { useState } from "react";
import { Layout, Menu, ConfigProvider, theme, App as AntdApp } from "antd";
// @ts-ignore
import "antd/dist/reset.css";
import {
  ToolOutlined,
  SyncOutlined,
  FormatPainterOutlined,
  DiffOutlined,
  DotChartOutlined,
  NodeIndexOutlined,
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

type MenuItemGroup = {
  key: string;
  icon?: React.ReactNode;
  label: string;
} | {
  key: string;
  groupTitle: string;
  items: Array<{
    key: string;
    icon?: React.ReactNode;
    label: string;
  }>;
};

const menuItems: MenuItemGroup[] = [
  {
    key: "home",
    icon: <ToolOutlined />,
    label: "TToy",
  },
  {
    key: "tools",
    groupTitle: "Tools",
    items: [
      { key: "fmt", icon: <FormatPainterOutlined />, label: "Format" },
      { key: "cvt", icon: <SyncOutlined />, label: "Convert" },
      { key: "codec", icon: <SyncOutlined />, label: "Encode" },
      { key: "generator", icon: <DotChartOutlined />, label: "Generate" },
      { key: "data-view", icon: <DotChartOutlined />, label: "View" },
      { key: "graph", icon: <NodeIndexOutlined />, label: "Plot" },
      { key: "diff", icon: <DiffOutlined />, label: "Diff" },
    ],
  },
  {
    key: "games",
    groupTitle: "Games",
    items: [{ key: "hanoi", label: "Hanoi" }],
  },
];

const AppContent = () => {
  const { isDark } = useTheme();
  const [selectedKey, setSelectedKey] = useState("home");

  const handleMenuClick = (e: { key: string }) => {
    setSelectedKey(e.key);
  };

  return (
    <ConfigProvider
      theme={{
        algorithm: isDark ? theme.darkAlgorithm : theme.defaultAlgorithm,
      }}
    >
      <AntdApp>
        <Layout style={{ minHeight: "100vh", paddingTop: "46px"}}>
          <Sider>
            <Menu
              selectedKeys={[selectedKey]}
              onClick={handleMenuClick}
              style={{ height: "100%", borderRight: 0, textAlign: "left" }}
            >
              {menuItems.map((group) =>
                "items" in group ? (
                  <Menu.ItemGroup key={group.key} title={group.groupTitle}>
                    {group.items.map((item) => (
                      <Menu.Item key={item.key} icon={item.icon}>
                        {item.label}
                      </Menu.Item>
                    ))}
                  </Menu.ItemGroup>
                ) : (
                  <Menu.Item key={group.key} icon={group.icon}>
                    {group.label}
                  </Menu.Item>
                ),
              )}
            </Menu>
          </Sider>
          <Layout>
            <Content style={{ padding: "0 15px 15px", position: "relative" }}>
              <div
                style={{ position: "absolute", top: 0, right: 15, zIndex: 1 }}
              >
                <ThemeSwitch />
              </div>
              {selectedKey === "home" && <Home />}
              {selectedKey === "fmt" && <Format />}
              {selectedKey === "cvt" && <Convert />}
              {selectedKey === "codec" && <Encode />}
              {selectedKey === "generator" && <Generate />}
              {selectedKey === "diff" && <Diff />}
              {selectedKey === "data-view" && <View />}
              {selectedKey === "graph" && <Plot />}
              {selectedKey === "hanoi" && "Hi"}
            </Content>
          </Layout>
        </Layout>
      </AntdApp>
    </ConfigProvider>
  );
};

const App = () => {
  return (
    <ThemeProvider>
      <AppContent />
    </ThemeProvider>
  );
};

export default App;
