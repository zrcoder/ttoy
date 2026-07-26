import { App as AntdApp } from "antd";

type Transformer = (input: string) => Promise<string>;
type ResultHandler = (result: string) => void;

export const useTransformer = () => {
  const { modal } = AntdApp.useApp();

  const fn = (
    input: string,
    transformer: Transformer,
    resultHandler: ResultHandler,
  ) => {
    transformer(input)
      .then((result) => {
        resultHandler(result);
      })
      .catch((error: Error) => {
        modal.error({ content: error.toString() });
      });
  };

  return fn;
};
