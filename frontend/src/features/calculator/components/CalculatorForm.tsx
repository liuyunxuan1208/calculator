'use client';
import { useState } from 'react';
import { useCalculate } from '../hooks/useCalculate';

/**
 * 计算器表单组件
 * 包含数字输入、操作符选择和计算按钮
 */
export function CalculatorForm() {
  const [num1, setNum1] = useState('');
  const [num2, setNum2] = useState('');
  const [operator, setOperator] = useState<string>('+'); // '+'表示加法，'-'表示减法，'*'表示乘法，'/'表示除法
  const { calculate, result, error } = useCalculate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    calculate({ first_number: Number(num1), second_number: Number(num2), operator: operator });
  };

  return (
    <form onSubmit={handleSubmit}>
      <input 
        type="number" 
        value={num1} 
        onChange={(e) => setNum1(e.target.value)} 
        placeholder="第一个数字"
      />
      <select value={operator} onChange={(e) => setOperator(e.target.value)}>
        <option value="+">+</option>
        <option value="-">-</option>
        <option value="*">*</option>
        <option value="/">÷</option>
      </select>
      <input 
        type="number" 
        value={num2} 
        onChange={(e) => setNum2(e.target.value)} 
        placeholder="第二个数字"
      />
      <button type="submit">计算</button>
      
      {result !== null && <div>结果: {result}</div>}
      {error && <div style={{ color: 'red' }}>{error}</div>}
    </form>
  );
}