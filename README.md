# connect-rule-engine

[English](README.md) | [中文](README_zh.md)

## Introduction
connect-rule-engine is a rule engine system based on Benthos. It provides a flexible platform for configuring and managing data processing rules. Key features include:

1. Dynamic Rule Configuration: Add, modify, and delete data processing rules dynamically through API interfaces.

2. Benthos-based: Leverages Benthos' powerful data processing capabilities, supporting various input/output formats and processing operations.

3. Persistent Storage: Uses a database to save rule configurations, ensuring rules are not lost after system restarts.

4. Real-time Processing: Supports real-time launching of new Benthos instances to process data streams.

5. RESTful API: Provides simple and easy-to-use API interfaces for integration with other systems.

6. Extensibility: Developed in Go, making it easy to extend and customize new features.

This system is suitable for scenarios requiring flexible data processing rules, such as log processing, data transformation, stream processing, etc.

## Usage

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/your-username/connect-rule-engine.git
   ```

2. Navigate to the project directory:
   ```
   cd connect-rule-engine
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

### Usage
1. Start the service:
   ```
   go run main.go
   ```

2. Manage rules through API interfaces:
   - Add a rule:
     ```
     curl -X POST http://localhost:8080/rules \
     -H "Content-Type: application/json" \
     -d '{
     "config_name": "example_config",
     "config": "# Config fields, showing default values\ninput:\n  label: \"\"\n  generate:\n    mapping: root = \"hello world\"\n    interval: 1s\n    count: 0\n    batch_size: 1\n    auto_replay_nacks: true\noutput:\n  type: stdout"
     }'
     ```

   - Delete a rule:
     ```
     curl -X DELETE http://localhost:8080/rules/{config_name}
     ```
     
   - Get all rules:
     ```
     curl -X GET http://localhost:8080/rules
     ```
   
   - Get a specific rule:
     ```
     curl -X GET http://localhost:8080/rules/{config_name}
     ```