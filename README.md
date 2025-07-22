# 通用 Go 后端服务框架

本框架总结自实际生产项目，适用于中大型 Go 后端服务，强调分层解耦、依赖注入、易扩展、易维护。

---

## 一、架构分层

- **入口层（cmd/）**：每个服务（如 api、admin、consumer、cron）有独立 main.go 和 wire.go，负责组装应用。
- **配置管理（config/）**：统一加载和管理配置，支持多环境。
- **日志与监控（log/、apm/）**：日志模块和 APM 集成，便于排查和监控。
- **数据库与缓存（db/）**：支持多种数据库（Postgres、ClickHouse）、缓存（Redis）、第三方服务。
- **中间件（middleware/）**：统一的认证、鉴权、限流、追踪等中间件。
- **领域模型（domain/）**：业务核心数据结构定义。
- **数据访问层（repo/）**：负责与数据库、外部服务的数据交互。
- **用例层（usecase/）**：业务逻辑实现，解耦 handler 与 repo。
- **接口适配层（handler/）**：HTTP handler、MQ handler、Admin handler 等，负责路由和请求响应。
- **服务启动与依赖注入（wire+ProviderSet）**：使用 Google Wire 进行依赖注入，自动组装各层组件。

---

## 二、目录结构

```text
cmd/
  api/
  admin/
  consumer/
  cron/
config/
db/
domain/
handler/
  v1/
  admin/
  mq/
log/
middleware/
mq/
repo/
server/
  http/
usecase/
utils/
```

---

## 三、依赖注入与模块解耦

- 每个模块（如 config、db、log、repo、usecase、handler）都通过 ProviderSet 注册，主入口通过 wire.Build 自动组装。
- 便于单元测试和模块替换。

---

## 四、启动流程

1. 加载配置（config）
2. 初始化日志、APM
3. 初始化数据库、缓存、MQ
4. 组装 repo、usecase、handler
5. 启动 HTTP 服务或 MQ 消费者或定时任务

---

## 五、适用场景

- 适合中大型 Go 后端项目
- 支持多服务（API、Admin、Consumer、Cron）
- 易于扩展和维护

---

## 六、推荐依赖

- [Google Wire](https://github.com/google/wire)：依赖注入
- [Echo](https://echo.labstack.com/)：高性能 Web 框架
- [GORM](https://gorm.io/)：ORM 框架
- [OpenTelemetry](https://opentelemetry.io/)：分布式追踪
- [Viper](https://github.com/spf13/viper)：配置管理
- [Redis](https://github.com/redis/go-redis)：缓存
- [Kafka/Sarama](https://github.com/Shopify/sarama)：消息队列

---

## 七、最佳实践

- 业务逻辑全部放在 usecase 层，handler 只做参数校验和响应封装。
- 所有依赖通过 ProviderSet 注入，便于 mock 和测试。
- 配置、日志、监控、数据库连接等均集中管理。
- 统一错误处理和响应格式。

---

如需详细代码模板或某一层实现示例，请补充说明。
