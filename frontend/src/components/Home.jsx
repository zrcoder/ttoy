import React from "react";
import { Typography, Row, Col } from "antd";
import { ToolOutlined } from "@ant-design/icons";

const { Title, Paragraph } = Typography;

const Home = () => {
  return (
    <Row align="middle" justify="center" style={{ height: "100%" }}>
      <Col style={{ textAlign: "center" }}>
        <Row align="middle" justify="center" gutter={16}>
          <Col>
            <ToolOutlined style={{ fontSize: "32px", color: "#bbb" }} />
          </Col>
          <Col>
            <Title level={1} style={{ margin: 0, color: "#bbb" }}>
              TToy
            </Title>
          </Col>
        </Row>
        <Paragraph
          style={{
            marginTop: "16px",
            fontSize: "18px",
            lineHeight: "1.5",
            color: "#888",
          }}
        >
          My dev tools.
        </Paragraph>
      </Col>
    </Row>
  );
};

export default Home;
