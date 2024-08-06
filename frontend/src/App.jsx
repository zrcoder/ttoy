import React, { useState } from "react";
import { Layout, Menu, ConfigProvider, theme, App as AntdApp } from "antd";
import "antd/dist/reset.css";
import {
  ToolOutlined,
  FileOutlined,
  FileTextOutlined,
  CodeOutlined,
} from "@ant-design/icons";
import Home from "./components/Home";
import JsonYamlConverter from "./components/JsonYamlConverter";
import JsonTomlConverter from "./components/JsonTomlConverter";
import YamlTomlConverter from "./components/YamlTomlConverter";
import JsonFormatter from "./components/JsonFormatter";
import YamlFormatter from "./components/YamlFormatter";
import TomlFormatter from "./components/TomlFormatter";
import HtmlFormatter from "./components/HtmlFormatter";

const { Sider, Content } = Layout;

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
              <Menu.Item
                key="home"
                icon={<ToolOutlined />}
                style={{ textAlign: "left" }}
              >
                TToy
              </Menu.Item>

              <Menu.ItemGroup title="Converters" style={{ textAlign: "left" }}>
                <Menu.Item
                  key="cvt-jy"
                  icon={<FileOutlined />}
                  style={{ textAlign: "left" }}
                >
                  json - yaml
                </Menu.Item>
                <Menu.Item
                  key="cvt-jt"
                  icon={<FileTextOutlined />}
                  style={{ textAlign: "left" }}
                >
                  json - toml
                </Menu.Item>
                <Menu.Item
                  key="cvt-yt"
                  icon={<CodeOutlined />}
                  style={{ textAlign: "left" }}
                >
                  yaml - toml
                </Menu.Item>
              </Menu.ItemGroup>

              <Menu.ItemGroup
                title="Encoders/Decoders"
                style={{ textAlign: "left" }}
              >
                <Menu.Item
                  key="documents"
                  icon={<FileOutlined />}
                  style={{ textAlign: "left" }}
                >
                  Documents
                </Menu.Item>
              </Menu.ItemGroup>

              <Menu.ItemGroup title="Formatters" style={{ textAlign: "left" }}>
                <Menu.Item
                  key="fmt-j"
                  icon={<FileOutlined />}
                  style={{ textAlign: "left" }}
                >
                  json
                </Menu.Item>
                <Menu.Item
                  key="fmt-y"
                  icon={<FileTextOutlined />}
                  style={{ textAlign: "left" }}
                >
                  yaml
                </Menu.Item>
                <Menu.Item
                  key="fmt-t"
                  icon={<CodeOutlined />}
                  style={{ textAlign: "left" }}
                >
                  toml
                </Menu.Item>
                <Menu.Item key="fmt-h" icon={<FileTextOutlined />}>
                  html
                </Menu.Item>
              </Menu.ItemGroup>
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
            </Content>
          </Layout>
        </Layout>
      </AntdApp>
    </ConfigProvider>
  );
};

export default App;
