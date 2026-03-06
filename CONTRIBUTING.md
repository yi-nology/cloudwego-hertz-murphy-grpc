# 贡献指南

感谢你考虑为 CloudWeGo 微服务脚手架做出贡献！🎉

## 📋 目录

- [行为准则](#行为准则)
- [如何贡献](#如何贡献)
- [开发流程](#开发流程)
- [代码规范](#代码规范)
- [提交规范](#提交规范)
- [问题反馈](#问题反馈)

---

## 行为准则

本项目采用贡献者公约作为行为准则。参与此项目即表示你同意遵守其条款。请阅读 [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 了解详情。

---

## 如何贡献

### 报告 Bug

在提交 Bug 报告之前，请先：
1. 检查 [Issues](https://github.com/yi-nology/cloudwego-template/issues) 中是否已有相同问题
2. 确认你使用的是最新版本
3. 阅读文档，确认不是使用方式问题

提交 Bug 报告时，请包含：
- **清晰的标题和描述**
- **复现步骤**（越详细越好）
- **预期行为** vs **实际行为**
- **环境信息**（Go 版本、操作系统等）
- **日志或截图**（如有）

[提交 Bug →](https://github.com/yi-nology/cloudwego-template/issues/new?template=bug_report.md)

### 建议新功能

我们非常欢迎新功能建议！请包含：
- **功能描述**
- **使用场景**
- **期望的实现方式**（可选）
- **参考示例**（如有）

[提交功能建议 →](https://github.com/yi-nology/cloudwego-template/issues/new?template=feature_request.md)

---

## 开发流程

### 1. Fork 和克隆

```bash
# Fork 后克隆你的仓库
git clone https://github.com/your-username/cloudwego-template.git
cd cloudwego-template

# 添加上游仓库
git remote add upstream https://github.com/yi-nology/cloudwego-template.git
```

### 2. 创建分支

```bash
# 创建特性分支
git checkout -b feature/amazing-feature

# 或修复分支
git checkout -b fix/bug-description
```

### 3. 开发和测试

```bash
# 安装依赖
go mod tidy

# 运行测试
make test

# 代码检查
make lint

# 运行服务
make run
```

### 4. 提交更改

```bash
# 添加更改
git add .

# 提交（遵循提交规范）
git commit -m "feat: 添加新功能描述"

# 推送到你的仓库
git push origin feature/amazing-feature
```

### 5. 创建 Pull Request

1. 访问你的 Fork 仓库
2. 点击 "New Pull Request"
3. 填写 PR 模板
4. 等待代码审查

---

## 代码规范

### Go 代码规范

遵循 [Effective Go](https://golang.org/doc/effective_go) 和 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)。

**关键点：**

```go
// ✅ 好的实践
func GetUser(ctx context.Context, id int64) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user id: %d", id)
    }
    
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}

// ❌ 避免
func get_user(ctx context.Context, ID int64) (*User, error) {
    // 缺少错误检查
    user, _ := s.repo.GetByID(ctx, ID)
    return user, nil
}
```

### 目录结构规范

```
internal/
├── app/          # 业务逻辑
├── repo/         # 数据访问
├── conf/         # 配置管理
└── pkg/          # 工具库
```

### 注释规范

```go
// UserService 提供用户相关的业务逻辑
type UserService struct {
    repo *dao.UserRepository
}

// GetUser 根据ID获取用户信息
// 
// ctx: 上下文
// id: 用户ID，必须大于0
// 
// 返回用户信息或错误
func (s *UserService) GetUser(ctx context.Context, id int64) (*User, error) {
    // 实现...
}
```

---

## 提交规范

我们使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

### 提交格式

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Type 类型

- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式（不影响功能）
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建/工具链相关
- `perf`: 性能优化

### 示例

```bash
# 新功能
git commit -m "feat: 添加用户注册功能"

# Bug 修复
git commit -m "fix: 修复用户登录时的空指针问题"

# 文档更新
git commit -m "docs: 更新 README 中的安装说明"

# 重构
git commit -m "refactor(database): 优化数据库连接池配置"

# 多行提交
git commit -m "feat: 添加用户权限管理

- 添加角色和权限模型
- 实现权限检查中间件
- 添加权限管理 API

Closes #123"
```

---

## 问题反馈

### 获取帮助

- 📖 查看 [文档](docs/)
- 💬 [GitHub Discussions](https://github.com/yi-nology/cloudwego-template/discussions)
- 🐛 [提交 Issue](https://github.com/yi-nology/cloudwego-template/issues)

### 安全问题

如果发现安全漏洞，请**不要**公开提交 Issue。

请发送邮件到：security@example.com

我们会尽快处理并致谢！

---

## 代码审查

所有提交都需要通过代码审查：

1. ✅ 代码符合规范
2. ✅ 测试通过
3. ✅ 文档更新
4. ✅ 提交信息清晰
5. ✅ 无冲突

---

## 许可证

通过贡献代码，你同意你的代码将以 MIT 许可证发布。

---

## 致谢

感谢所有贡献者！ 🎉

### 贡献者

<!-- ALL-CONTRIBUTORS-LIST:START -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

---

**再次感谢你的贡献！** ❤️

有问题随时在 [Discussions](https://github.com/yi-nology/cloudwego-template/discussions) 中提出。
