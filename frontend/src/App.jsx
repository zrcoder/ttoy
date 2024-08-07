import React, { useState } from "react";
import { Layout, Menu, ConfigProvider, theme, App as AntdApp } from "antd";
import "antd/dist/reset.css";
import {
  ToolOutlined,
  SyncOutlined,
  FormatPainterOutlined,
  DiffOutlined,
  DotChartOutlined,
} from "@ant-design/icons";
import Home from "./components/Home";
import Converter from "./components/Converter";
import Formatter from "./components/Formatter";
import TextDiffer from "./components/TextDiffer";
import DataViewer from "./components/DataViewer";

const { Sider, Content } = Layout;

const menuItems = [
  {
    key: "home",
    icon: <ToolOutlined />,
    label: "TToy",
  },
  {
    key: "data", // Unique key for grouping
    groupTitle: "Data",
    items: [
      { key: "cvt", icon: <SyncOutlined />, label: "Convertor" },
      { key: "fmt", icon: <FormatPainterOutlined />, label: "Formater" },
    ],
  },
  {
    key: "generators", // Unique key for grouping
    groupTitle: "Generators",
    items: [
      { key: "data-view", icon: <DotChartOutlined />, label: "Data Viewer" },
    ],
  },
  {
    key: "text", // Unique key for grouping
    groupTitle: "Text",
    items: [{ key: "diff", icon: <DiffOutlined />, label: "Differ" }],
  },
];

const App = () => {
  const [selectedKey, setSelectedKey] = useState("home");

  const handleMenuClick = (e) => {
    setSelectedKey(e.key);
  };

  return (
    <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>
      <AntdApp>
        <Layout style={{ minHeight: "100vh" }}>
          <Sider className="site-layout-background">
            <Menu
              selectedKeys={[selectedKey]}
              onClick={handleMenuClick}
              style={{ height: "100%", borderRight: 0, textAlign: "left" }}
            >
              {menuItems.map((group) =>
                group.items ? (
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
                )
              )}
            </Menu>
          </Sider>
          <Layout>
            <Content style={{ padding: 15 }}>
              {selectedKey === "home" && <Home />}
              {selectedKey === "cvt" && <Converter />}
              {selectedKey === "fmt" && <Formatter />}
              {selectedKey === "diff" && <TextDiffer />}
              {selectedKey === "data-view" && <DataViewer />}
            </Content>
          </Layout>
        </Layout>
      </AntdApp>
    </ConfigProvider>
  );
};

export default App;
