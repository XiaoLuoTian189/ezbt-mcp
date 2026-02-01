package databases

import (
	"context"
	"ezbt-mcp/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetMysqlList = "get_mysql_list"
)

var GetMysqlListTool = mcp.NewTool(
	GetMysqlList,
	mcp.WithDescription("获取MySQL数据库列表"),
)

func GetMysqlListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	return bt.Request("datalist/data/get_data_list", map[string]string{
		"p":      "1",
		"limit":  "100000",
		"table":  "databases",
		"search": "",
	})
}
