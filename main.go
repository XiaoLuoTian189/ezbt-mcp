package main

import (
	"flag"
	"fmt"
	"log"
	"ezbt-mcp/modules/databases"
	"ezbt-mcp/modules/docker"
	"ezbt-mcp/modules/email"
	"ezbt-mcp/modules/files"
	"ezbt-mcp/modules/sites"
	"ezbt-mcp/modules/system"

	"github.com/mark3labs/mcp-go/server"
)

func createServer() *server.MCPServer {
	return server.NewMCPServer(
		"mcp-ezbt 🚀",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
}

func registerTools(s *server.MCPServer) {
	s.AddTool(system.GetPublicConfigTool, system.GetPublicConfigHandle)
	s.AddTool(system.GetNetWorkTool, system.GetNetWorkHandle)
	s.AddTool(sites.GetSitesListTool, sites.GetSitesListHandle)
	s.AddTool(sites.AddSiteTool, sites.AddSiteHandle)
	s.AddTool(databases.GetMysqlListTool, databases.GetMysqlListHandle)
	s.AddTool(email.GetMailsListTool, email.GetMailsListHandle)
	s.AddTool(email.AddMailboxTool, email.AddMailboxHandle)
	s.AddTool(docker.GetContainerListTool, docker.GetContainerListHandle)
	s.AddTool(docker.GetContainerInfoTool, docker.GetContainerInfoHandle)
	s.AddTool(docker.GetImageListTool, docker.GetImageListHandle)
	s.AddTool(files.GetFileContentTool, files.GetFileContentHandle)
	s.AddTool(files.SaveFileContentTool, files.SaveFileContentHandle)
	s.AddTool(files.SetFilePermissionsTool, files.SetFilePermissionsHandle)
	s.AddTool(files.CreateFileTool, files.CreateFileHandle)
	s.AddTool(files.CreateDirTool, files.CreateDirHandle)
	s.AddTool(files.DeleteFileTool, files.DeleteFileHandle)
}

func startServer(s *server.MCPServer, useSSE bool) {
	if useSSE {
		var port = "8080"
		log.Panicf("SSE Server starting on port %s", port)
		sseServer := server.NewSSEServer(s, server.WithBaseURL(fmt.Sprintf("http://localhost:%s", port)))
		if err := sseServer.Start(fmt.Sprintf("http://localhost:%s", port)); err != nil {
			fmt.Printf("SSE Server error: %v\n", err)
		}
	} else {
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Stdio Server error: %v\n", err)
		}
	}
}

func main() {
	var (
		useSSE = flag.Bool("sse", false, "use SSE mode")
	)
	flag.Parse()
	s := createServer()
	registerTools(s)
	startServer(s, *useSSE)
}
