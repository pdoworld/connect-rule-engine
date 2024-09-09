# connect-rule-engine

## 简介
connect-rule-engine 是一个基于 Benthos 的规则引擎系统。它提供了一个灵活的平台，用于配置和管理数据处理规则。主要特点包括：

1. 动态规则配置：通过 API 接口可以动态添加、修改和删除数据处理规则。

2. 基于 Benthos：利用 Benthos 强大的数据处理能力，支持多种输入输出格式和处理操作。

3. 持久化存储：使用数据库保存规则配置，确保系统重启后规则不会丢失。

4. 实时处理：支持实时启动新的 Benthos 实例来处理数据流。

5. RESTful API：提供简单易用的 API 接口，方便与其他系统集成。

6. 可扩展性：基于 Go 语言开发，易于扩展和定制新功能。

这个系统适用于需要灵活数据处理规则的场景，如日志处理、数据转换、流式处理等。

## 使用方法

### 安装

1. 克隆仓库到本地：
   ```
   git clone https://github.com/your-username/connect-rule-engine.git
   ```

2. 进入项目目录：
   ```
   cd connect-rule-engine
   ```

3. 安装依赖：
   ```
   go mod tidy

### 使用
1. 启动服务：
   ```
   go run main.go
   ```

2. 通过 API 接口进行规则管理：
   - 添加规则：
     ```
        curl -X POST http://localhost:8080/rules \
        -H "Content-Type: application/json" \
        -d '{
        "config_name": "example_config",
        "config": "# Config fields, showing default values\ninput:\n  label: \"\"\n  generate:\n    mapping: root = \"hello world\"\n    interval: 1s\n    count: 0\n    batch_size: 1\n    auto_replay_nacks: true\noutput:\n  type: stdout"
        }'
     ```

   - 删除规则：
     ```
     curl -X DELETE http://localhost:8080/rules/{config_name}
     ```
     
   - 获取所有规则：
     ```
     curl -X GET http://localhost:8080/rules
     ```
   
   - 获取特定规则：
     ```
     curl -X GET http://localhost:8080/rules/{config_name}
     ```