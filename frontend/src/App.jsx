import React, { useState } from "react";
import { Layout, Menu, ConfigProvider, theme, App as AntdApp } from "antd";
import "antd/dist/reset.css";
import {
  ToolOutlined,
  FileOutlined,
  FileTextOutlined,
  CodeOutlined,
  DiffOutlined,
  SwapOutlined,
  AppstoreAddOutlined,
} from "@ant-design/icons";
import Home from "./components/Home";
import JsonYamlConverter from "./components/JsonYamlConverter";
import JsonTomlConverter from "./components/JsonTomlConverter";
import YamlTomlConverter from "./components/YamlTomlConverter";
import JsonFormatter from "./components/JsonFormatter";
import YamlFormatter from "./components/YamlFormatter";
import TomlFormatter from "./components/TomlFormatter";
import HtmlFormatter from "./components/HtmlFormatter";
import TextDiffer from "./components/TextDiffer";
import JsonSvgViewer from "./components/JsonSvgViewr";

const { Sider, Content } = Layout;

const menuItems = [
  {
    key: "home",
    icon: <ToolOutlined />,
    label: "TToy",
  },
  {
    groupTitle: "Converters",
    items: [
      { key: "cvt-jy", icon: <SwapOutlined />, label: "json-yaml" },
      { key: "cvt-jt", icon: <AppstoreAddOutlined />, label: "json-toml" },
      { key: "cvt-yt", icon: <CodeOutlined />, label: "yaml-toml" },
    ],
  },
  {
    groupTitle: "Formatters",
    items: [
      { key: "fmt-j", icon: <FileOutlined />, label: "json" },
      { key: "fmt-y", icon: <FileTextOutlined />, label: "yaml" },
      { key: "fmt-t", icon: <CodeOutlined />, label: "toml" },
      { key: "fmt-h", icon: <FileTextOutlined />, label: "html" },
    ],
  },
  {
    groupTitle: "Generators",
    items: [{ key: "gen-json-svg", icon: <DiffOutlined />, label: "json-svg" }],
  },
  {
    groupTitle: "Text",
    items: [{ key: "diff", icon: <DiffOutlined />, label: "diff" }],
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
              {menuItems.map((group, index) =>
                group.items ? (
                  <Menu.ItemGroup
                    key={index}
                    title={group.groupTitle}
                    style={{ textAlign: "left" }}
                  >
                    {group.items.map((item) => (
                      <Menu.Item
                        key={item.key}
                        icon={item.icon}
                        style={{ textAlign: "left" }}
                      >
                        {item.label}
                      </Menu.Item>
                    ))}
                  </Menu.ItemGroup>
                ) : (
                  <Menu.Item
                    key={group.key}
                    icon={group.icon}
                    style={{ textAlign: "left" }}
                  >
                    {group.label}
                  </Menu.Item>
                )
              )}
            </Menu>
          </Sider>
          <Layout>
            <Content style={{ paddingTop: 15 }}>
              {selectedKey === "home" && <Home />}
              {selectedKey === "cvt-jy" && <JsonYamlConverter />}
              {selectedKey === "cvt-jt" && <JsonTomlConverter />}
              {selectedKey === "cvt-yt" && <YamlTomlConverter />}
              {selectedKey === "fmt-j" && <JsonFormatter />}
              {selectedKey === "fmt-y" && <YamlFormatter />}
              {selectedKey === "fmt-t" && <TomlFormatter />}
              {selectedKey === "fmt-h" && <HtmlFormatter />}
              {selectedKey === "diff" && <TextDiffer />}
              {selectedKey === "gen-json-svg" && <JsonSvgViewer />}
            </Content>
          </Layout>
        </Layout>
      </AntdApp>
    </ConfigProvider>
  );
};

export default App;
