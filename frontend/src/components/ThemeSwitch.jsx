import React from "react";
import { Switch } from "antd";
import { MoonFilled, SunFilled } from "@ant-design/icons";
import { useTheme } from "../contexts/ThemeContext";

const ThemeSwitch = () => {
  const { isDark, toggleTheme } = useTheme();

  return (
    <Switch
      checked={isDark}
      onChange={toggleTheme}
      checkedChildren={<MoonFilled />}
      unCheckedChildren={<SunFilled />}
    />
  );
};

export default ThemeSwitch;
