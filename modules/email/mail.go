package email

import (
	"context"
	"fmt"
	"math/rand"
	"mcp_btez/utils"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	GetMailsList = "get_mails_list"
	AddMailbox   = "add_mailbox"
)

var GetMailsListTool = mcp.NewTool(
	GetMailsList,
	mcp.WithDescription("获取指定邮箱的邮件列表"),
	mcp.WithString("username",
		mcp.Required(),
		mcp.Description("邮箱地址"),
	),
	mcp.WithString("p",
		mcp.Description("页码"),
	),
)

var AddMailboxTool = mcp.NewTool(
	AddMailbox,
	mcp.WithDescription("为指定域名添加邮箱"),
	mcp.WithString("username",
		mcp.Required(),
		mcp.Description("邮箱地址或域名"),
	),
	mcp.WithString("password",
		mcp.Description("邮箱密码，不填则随机生成"),
	),
	mcp.WithString("full_name",
		mcp.Description("全名，不填则随机生成"),
	),
	mcp.WithString("quota",
		mcp.Description("邮箱容量，格式：数字+GB，例如：5 GB"),
	),
	mcp.WithString("is_admin",
		mcp.Description("是否为管理员"),
	),
	mcp.WithString("active",
		mcp.Description("是否激活"),
	),
)

func GetMailsListHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, fmt.Errorf("username必须是字符串")
	}

	params := map[string]string{
		"username": username,
	}

	if p, ok := request.Params.Arguments["p"].(string); ok && p != "" {
		params["p"] = p
	} else {
		params["p"] = "1"
	}

	return bt.Request("mail/main/get_mails", params)
}

func generateRandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialChars := "!@#$%^&*"

	length := 10 + rand.Intn(6)
	password := make([]byte, length)

	password[0] = chars[26+rand.Intn(26)]
	password[1] = chars[rand.Intn(26)]
	password[2] = chars[52+rand.Intn(10)]
	password[3] = specialChars[rand.Intn(len(specialChars))]

	for i := 4; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	for i := range password {
		j := rand.Intn(i + 1)
		password[i], password[j] = password[j], password[i]
	}

	return string(password)
}

func generateRandomUsername(domain string) string {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"user", "mail", "info", "contact", "support", "admin", "service"}
	prefix := prefixes[rand.Intn(len(prefixes))] + fmt.Sprintf("%d", rand.Intn(1000))
	return prefix + "@" + domain
}

func generateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	prefixes := []string{"User", "Mail", "Info", "Contact", "Support", "Admin", "Service"}
	return prefixes[rand.Intn(len(prefixes))] + fmt.Sprintf("%d", rand.Intn(1000))
}

func formatQuota(quota string) string {
	if quota == "" {
		return "5 GB"
	}

	quota = strings.ReplaceAll(quota, " ", "")

	if strings.Contains(strings.ToLower(quota), "gb") {
		quota = strings.ToLower(quota)
		quota = strings.ReplaceAll(quota, "gb", "")
		return quota + " GB"
	}

	return quota + " GB"
}

func AddMailboxHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	bt := utils.NewBTPanel(utils.GetBaseURL(), utils.GetApiToken())

	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, fmt.Errorf("username必须是字符串")
	}

	if !strings.Contains(username, "@") {
		domain := username
		username = generateRandomUsername(domain)
	}

	password := ""
	if pwd, ok := request.Params.Arguments["password"].(string); ok && pwd != "" {
		password = pwd
	} else {
		password = generateRandomPassword()
	}

	fullName := ""
	if name, ok := request.Params.Arguments["full_name"].(string); ok && name != "" {
		fullName = name
	} else {
		parts := strings.Split(username, "@")
		if len(parts) > 0 && parts[0] != "" {
			fullName = parts[0]
		} else {
			fullName = generateRandomName()
		}
	}

	quota := "5 GB"
	if q, ok := request.Params.Arguments["quota"].(string); ok && q != "" {
		quota = formatQuota(q)
	}

	isAdmin := "0"
	if admin, ok := request.Params.Arguments["is_admin"].(string); ok && admin != "" {
		isAdmin = admin
	}

	active := "1"
	if a, ok := request.Params.Arguments["active"].(string); ok && a != "" {
		active = a
	}

	params := map[string]string{
		"username":  username,
		"password":  password,
		"full_name": fullName,
		"quota":     quota,
		"is_admin":  isAdmin,
		"active":    active,
	}

	return bt.Request("mail/main/add_mailbox", params)
}
