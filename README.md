# 计算器项目

## 项目概述
这是一个基于Go和React的全栈计算器应用，提供基本的四则运算功能以及计算历史记录。

## 功能说明
- 支持加、减、乘、除运算
- 记录计算历史
- 前后端分离架构

## 安装指南

### 后端安装
1. 确保已安装Go 1.20+版本
2. 克隆项目
3. 进入backend目录运行:
```bash
cd backend

go mod download

go run cmd/server/main.go
```

### 前端安装
1. 确保已安装Node.js 18+版本
2. 进入frontend目录运行:
```bash
cd frontend

npm install

npm run dev
```

## API接口说明

### 计算接口
- 路径: /api/v1/calculate
- 方法: POST
- 请求体:
```json
{
  "operand1": 1,
  "operand2": 2,
  "operator": "+"
}
```

### 获取历史记录
- 路径: /api/v1/history
- 方法: GET

## 开发文档

### 项目结构
```
calculator/
├── backend/      # 后端代码
├── frontend/     # 前端代码
└── README.md     # 项目文档
```

### 测试方法

#### 后端测试
```bash
cd backend

go test ./...
```

#### 前端测试
```bash
cd frontend

npm test
```

## 部署指南

### 生产环境构建
```bash
# 前端构建
cd frontend
npm run build

# 后端构建
cd backend
go build -o calculator cmd/server/main.go
```