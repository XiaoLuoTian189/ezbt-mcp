package system

import (
	"context"
	"mcp_btez/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetPublicConfig = "get_public_config"
	GetNetWork      = "GetNetWork"
)

var GetPublicConfigTool = mcp.NewTool(
	GetPublicConfig,
	mcp.WithDescription("获取宝塔面板公共配置"),
)

var GetNetWorkTool = mcp.NewTool(
	GetNetWork,
	mcp.WithDescription("获取宝塔面板资源相关信息，比如CPU、内存、磁盘、网络等"),
)

func GetPublicConfigHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("panel/public/get_public_config", map[string]string{})
}

func GetNetWorkHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("system?action=GetNetWork", map[string]string{})
}
