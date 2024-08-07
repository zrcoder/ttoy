import React from "react";
import { Typography, Row } from "antd";

const { Title, Paragraph } = Typography;

const Home = () => {
  return (
    <Row
      align="middle" // 垂直居中对齐
      justify="center" // 水平居中对齐
      style={{ height: "100%" }} // 使 Row 占满父容器高度
    >
      <div style={{ textAlign: "center" }}>
        <Title level={1}>TToy</Title>
        <Paragraph>My dev tools.</Paragraph>
      </div>
    </Row>
  );
};

export default Home;
