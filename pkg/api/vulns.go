package api

type VulnsInfo struct {
	Name string
	Url  string
}

var vulns = map[string]VulnsInfo{
	"1": {Name: "xss漏洞", Url: "/xss/home"},
	"2": {Name: "文件包含漏洞", Url: "/include"},
	"3": {Name: "xxe外部实体注入漏洞", Url: "/xxe"},
	"4": {Name: "sql注入漏洞", Url: "/sql"},
	"5": {Name: "SSRF服务端请求伪造漏洞", Url: "/ssrf"},
	"6": {Name: "CSRF夸站请求伪造漏洞", Url: "/csrf"},
}
