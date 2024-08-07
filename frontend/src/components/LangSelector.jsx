import React from "react";
import { Radio } from "antd";

const LangSelector = ({ value, onChange, languages }) => {
  return (
    <div style={{ padding: "8px 0", display: "flex", alignItems: "center" }}>
      <Radio.Group
        value={value}
        onChange={(e) => onChange(e.target.value)}
        style={{ display: "flex", alignItems: "center" }}
      >
        {languages.map((lang) => (
          <Radio.Button
            key={lang}
            value={lang}
            style={{
              borderRadius: 0,
              fontSize: "12px",
            }}
          >
            {lang}
          </Radio.Button>
        ))}
      </Radio.Group>
    </div>
  );
};

export default LangSelector;
