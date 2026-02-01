package files

import (
	"context"
	"fmt"
	"ezbt-mcp/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetFileContent     = "get_file_content"
	SaveFileContent    = "save_file_content"
	SetFilePermissions = "set_file_permissions"
	CreateFile         = "create_file"
	CreateDir          = "create_dir"
	DeleteFile         = "delete_file"
)

var GetFileContentTool = mcp.NewTool(
	GetFileContent,
	mcp.WithDescription("获取文件内容（下载文件）"),
	mcp.WithString("path",
		mcp.Required(),
		mcp.Description("文件的完整路径"),
	),
)

func GetFileContentHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, ok := request.Params.Arguments["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path必须是字符串")
	}

	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("files?action=GetFileBody", map[string]string{
		"path": path,
	})
}

var SaveFileContentTool = mcp.NewTool(
	SaveFileContent,
	mcp.WithDescription("保存文件内容（上传/修改文件）"),
	mcp.WithString("path",
		mcp.Required(),
		mcp.Description("文件的完整路径"),
	),
	mcp.WithString("content",
		mcp.Required(),
		mcp.Description("要保存的文件内容"),
	),
	mcp.WithString("encoding",
		mcp.Description("内容编码，默认为 utf-8"),
	),
)

func SaveFileContentHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, ok := request.Params.Arguments["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path必须是字符串")
	}
	content, ok := request.Params.Arguments["content"].(string)
	if !ok {
		return nil, fmt.Errorf("content必须是字符串")
	}
	encoding, ok := request.Params.Arguments["encoding"].(string)
	if !ok {
		encoding = "utf-8"
	}

	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	bt.Request("files?action=CreateFile", map[string]string{
		"path": path,
	})

	return bt.Request("files?action=SaveFileBody", map[string]string{
		"path":     path,
		"data":     content,
		"encoding": encoding,
	})
}

var SetFilePermissionsTool = mcp.NewTool(
	SetFilePermissions,
	mcp.WithDescription("修改文件或目录权限"),
	mcp.WithString("path",
		mcp.Required(),
		mcp.Description("文件或目录的完整路径"),
	),
	mcp.WithString("chmod",
		mcp.Required(),
		mcp.Description("权限值，如 644 或 755"),
	),
	mcp.WithString("chown",
		mcp.Description("所有者，默认为 www"),
	),
)

func SetFilePermissionsHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, ok := request.Params.Arguments["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path必须是字符串")
	}
	chmod, ok := request.Params.Arguments["chmod"].(string)
	if !ok {
		return nil, fmt.Errorf("chmod必须是字符串")
	}
	chown, ok := request.Params.Arguments["chown"].(string)
	if !ok {
		chown = "www"
	}

	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("files?action=SetFileAccept", map[string]string{
		"filename": path,
		"chmod":    chmod,
		"chown":    chown,
	})
}

var CreateFileTool = mcp.NewTool(
	CreateFile,
	mcp.WithDescription("创建新文件"),
	mcp.WithString("path",
		mcp.Required(),
		mcp.Description("新文件的完整路径"),
	),
)

func CreateFileHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, ok := request.Params.Arguments["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path必须是字符串")
	}
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("files?action=CreateFile", map[string]string{
		"path": path,
	})
}

var CreateDirTool = mcp.NewTool(
	CreateDir,
	mcp.WithDescription("创建新目录"),
	mcp.WithString("path",
		mcp.Required(),
		mcp.Description("新目录的完整路径"),
	),
)

func CreateDirHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, ok := request.Params.Arguments["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path必须是字符串")
	}
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("files?action=CreateDir", map[string]string{
		"path": path,
	})
}

var DeleteFileTool = mcp.NewTool(
	DeleteFile,
	mcp.WithDescription("删除文件或目录"),
	mcp.WithString("path",
		mcp.Required(),
		mcp.Description("要删除的文件或目录路径"),
	),
)

func DeleteFileHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	path, ok := request.Params.Arguments["path"].(string)
	if !ok {
		return nil, fmt.Errorf("path必须是字符串")
	}
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())
	return bt.Request("files?action=DeleteFile", map[string]string{
		"path": path,
	})
}
