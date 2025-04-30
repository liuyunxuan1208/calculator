// @ts-ignore
import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { CalculatorService } from "../../../gen/src/gen/calculator/v1/calculator_connectweb";

interface CalculationRequest {
  first_number: number;
  second_number: number;
  operator: string;
}

interface CalculationResponse {
  result: number;
}

type CalculateResponse = {
  result: number;
};

/**
 * ConnectRPC客户端实现
 */
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const client = createPromiseClient(CalculatorService, transport);

/**
 * 执行计算操作
 * @param operand1 第一个数字
 * @param operand2 第二个数字
 * @param operator 运算符
 * @returns 计算结果
 */
export async function calculate(firstNumber: number, secondNumber: number, operator: string): Promise<number> {
  const response = await client.calculate({
    firstNumber,
    secondNumber,
    operator,
  });
  return (response as unknown as CalculateResponse).result;
}