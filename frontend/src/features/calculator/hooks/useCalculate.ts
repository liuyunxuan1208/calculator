import { useState } from 'react';
import { calculate } from '../api/calculatorService';

/**
 * 自定义hook封装计算逻辑
 * @returns 计算函数和结果状态
 */
export function useCalculate() {
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);

  const calculateHandler = async (params: {
    first_number: number;
    second_number: number;
    operator: string;
  }) => {
    try {
      setError(null);
      const result = await calculate(
        params.first_number,
        params.second_number,
        params.operator
      );
      setResult(result);
    } catch (err) {
      setError('计算过程中发生错误');
    }
  };

  return { calculate: calculateHandler, result, error };
}