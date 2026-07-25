import { Tabs } from "antd";

const AppTabs = ({ defaultActiveKey, items }) => {
  return <Tabs defaultActiveKey={defaultActiveKey} items={items} centered/>;
};

export default AppTabs;
