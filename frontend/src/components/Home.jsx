import React from "react";
import { Typography, Row } from "antd";

const { Title, Paragraph } = Typography;

const Home = () => {
  return (
    <Row align="middle" justify="center" style={{ height: "100%" }}>
      <div style={{ textAlign: "center" }}>
        <Title level={1}>TToy</Title>
        <Paragraph>My dev tools.</Paragraph>
      </div>
    </Row>
  );
};

export default Home;
