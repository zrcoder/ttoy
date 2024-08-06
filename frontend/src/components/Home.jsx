import React from "react";
import { Layout, Typography, Row, Col } from "antd";

const { Title, Paragraph } = Typography;

const Home = () => {
  return (
    <Layout style={{ height: "100vh" }}>
      <Row
        align="middle" // 垂直居中对齐
        justify="center" // 水平居中对齐
        style={{ height: "100%" }} // 使 Row 占满父容器高度
      >
        <Col>
          <div style={{ textAlign: "center" }}>
            <Title level={1}>TToy</Title>
            <Paragraph>My dev tools.</Paragraph>
          </div>
        </Col>
      </Row>
    </Layout>
  );
};

export default Home;
