# ezbt 🚀

ezbt 是一个基于 [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) 协议开发的**宝塔面板 (BT Panel) 智能管理工具**。它允许 AI 编辑器（如 Trae, Cursor, Claude Desktop 等）通过自然语言直接管理和操作你的宝塔面板服务器。

## ✨ 核心特性

- **📂 文件系统管理**：支持文件的读取、保存、创建、删除以及权限修改（chmod/chown）。
- **🌐 网站自动化**：一键查询网站列表、快速创建新的 PHP 网站。
- **🗄️ 数据库操作**：实时获取 MySQL 数据库列表及信息。
- **🐳 Docker 集成**：管理容器生命周期，查看容器详情及本地镜像列表。
- **📧 邮件服务**：支持邮箱账户的创建与邮件列表查询。
- **🖥️ 系统监控**：实时获取服务器 CPU、内存、磁盘及网络状态。

## 🛠️ 安装与构建

### 环境要求
- Go 1.18 或更高版本
- 已开启 API 接口的宝塔面板

### 编译步骤
1. 克隆仓库：
   ```bash
   git clone https://github.com/你的用户名/ezbt.git
   cd ezbt
   ```
2. 编译可执行文件：
   ```bash
   go build -o build/mcp-ezbt.exe main.go
   ```

## ⚙️ 配置说明

在 AI 编辑器（以 Trae/Cursor 为例）的 MCP 设置中添加以下配置：

```json
{
  "mcpServers": {
    "ezbt": {
      "command": "C:\\你的路径\\ezbt\\build\\mcp-ezbt.exe",
      "env": {
        "BT_BASE_URL": "http://你的面板地址:8888",
        "BT_API_TOKEN": "你的宝塔API密钥"
      }
    }
  }
}
```

> **注意**：请确保在宝塔面板后台将你运行 AI 编辑器的 IP 加入到 API 接口的白名单中。

## 🚀 使用示例

你可以直接对 AI 说：
- "帮我列出服务器上所有的网站。"
- "使用域名 test.com 创建一个新网站。"
- "读取 /www/wwwroot/test.com/config.php 的内容。"
- "把 /www/wwwroot/ 目录下的 index.php 权限改为 755。"
- "查看当前服务器的内存占用情况。"

## 📄 开源协议
本项目采用 [MIT License](LICENSE) 开源。
