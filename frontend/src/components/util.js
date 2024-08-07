import { App as AntdApp } from "antd";

/**
 * 自定义 Hook，用于处理 API 错误并显示 Modal
 * @returns {Function} 用于处理错误并显示 Modal 的函数
 */
export const useTransformer = () => {
  const { modal } = AntdApp.useApp();

  const fn = (input, transformer, resultHandler) => {
    transformer(input)
      .then((result) => {
        resultHandler(result);
      })
      .catch((error) => {
        modal.error({
          content: error.toString(),
        });
      });
  };

  return fn;
};
