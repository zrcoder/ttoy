import { App as AntdApp } from "antd";

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
