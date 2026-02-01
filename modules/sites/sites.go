package sites

import (
	"context"
	"encoding/json"
	"fmt"
	"ezbt-mcp/utils"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetSitesList = "get_sites_list"
	AddSite      = "add_site"
)

var GetSitesListTool = mcp.NewTool(
	GetSitesList,
	mcp.WithDescription("获取PHP网站项目列表"),
)

func GetSitesListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	return bt.Request("datalist/data/get_data_list", map[string]string{
		"type":   "-1",
		"search": "",
		"p":      "1",
		"limit":  "100000",
		"table":  "sites",
		"order":  "",
	})
}

type WebNameStruct struct {
	Domain     string   `json:"domain"`
	DomainList []string `json:"domainlist"`
	Count      int      `json:"count"`
}

var AddSiteTool = mcp.NewTool(
	AddSite,
	mcp.WithDescription("创建新的网站"),
	mcp.WithString("domains",
		mcp.Required(),
		mcp.Description("要添加的域名，多个域名用逗号分隔"),
	),
)

func AddSiteHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	domainsStr, ok := request.Params.Arguments["domains"].(string)
	if !ok {
		return nil, fmt.Errorf("domains必须是字符串")
	}

	domains := strings.Split(domainsStr, ",")
	if len(domains) == 0 {
		return nil, fmt.Errorf("至少需要一个域名")
	}

	for i, domain := range domains {
		domains[i] = strings.TrimSpace(domain)
	}

	mainDomain := domains[0]

	var domainList []string
	if len(domains) > 1 {
		domainList = domains[1:]
	} else {
		domainList = []string{}
	}

	webName := WebNameStruct{
		Domain:     mainDomain,
		DomainList: domainList,
		Count:      len(domainList),
	}

	webNameJSON, err := json.Marshal(webName)
	if err != nil {
		return nil, fmt.Errorf("构建webname失败: %v", err)
	}

	path := fmt.Sprintf("/www/wwwroot/%s", mainDomain)

	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	return bt.Request("site?action=AddSite", map[string]string{
		"path":           path,
		"ftp":            "false",
		"type":           "PHP",
		"type_id":        "0",
		"ps":             mainDomain,
		"port":           "80",
		"version":        "00",
		"need_index":     "0",
		"need_404":       "0",
		"sql":            "false",
		"codeing":        "utf8mb4",
		"webname":        string(webNameJSON),
		"add_dns_record": "false",
	})
}
