import { App as AntdApp } from "antd";

/**
 * 自定义 Hook，用于处理 API 错误并显示 Modal
 * @returns {Function} 用于处理错误并显示 Modal 的函数
 */
const useTransformAction = () => {
  const { modal } = AntdApp.useApp();

  const fn = (input, transformer, user) => {
    transformer(input)
      .then((res) => {
        user(res);
      })
      .catch((err) => {
        modal.error({
          content: err.toString(),
        });
      });
  };

  return fn;
};

export default useTransformAction;
