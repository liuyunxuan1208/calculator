import { CalculatorForm } from './components/CalculatorForm';

/**
 * 计算器功能页面
 * 整合所有计算器相关组件
 */
export function CalculatorPage() {
  return (
    <div>
      <h1>计算器</h1>
      <CalculatorForm />
    </div>
  );
}