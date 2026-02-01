package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"ezbt-mcp/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetContainerList = "get_container_list"
	GetContainerInfo = "get_container_info"
	GetImageList     = "get_image_list"
)

var GetContainerListTool = mcp.NewTool(
	GetContainerList,
	mcp.WithDescription("获取docker中所有容器的列表"),
)

var GetContainerInfoTool = mcp.NewTool(
	GetContainerInfo,
	mcp.WithDescription("获取指定容器的容器详情信息"),
	mcp.WithString("id",
		mcp.Required(),
		mcp.Description("容器ID，可以是短ID"),
	),
)

var GetImageListTool = mcp.NewTool(
	GetImageList,
	mcp.WithDescription("获取docker中本地镜像列表"),
)

func GetContainerListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("btdocker/container/get_list", map[string]string{})
}

func GetContainerInfoHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	containerID, ok := request.Params.Arguments["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id必须是字符串")
	}

	data := map[string]string{
		"id": containerID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("构建请求数据失败: %v", err)
	}

	return bt.Request("btdocker/container/get_container_info", map[string]string{
		"data": string(jsonData),
	})
}

func GetImageListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("btdocker/image/image_list", map[string]string{})
}
