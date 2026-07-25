import { Tabs, TabsProps } from "antd";

const AppTabs = ({ defaultActiveKey, items }: Pick<TabsProps, "defaultActiveKey" | "items">) => {
  return <Tabs defaultActiveKey={defaultActiveKey} items={items} />;
};

export default AppTabs;
