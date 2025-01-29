package api

type VulnsInfo struct {
	Name string
	Url  string
}

var vulns = map[string]VulnsInfo{
	"1": {Name: "反射型XSS漏洞", Url: "/api/v1/xss/reflected"},
	"2": {Name: "存储型XSS漏洞", Url: "/api/v1/xss/stored"},
	"3": {Name: "DOM型XSS漏洞", Url: "/api/v1/xss/dom"},
	"4": {Name: "XXE外部实体注入漏洞", Url: "/api/v1/xxe"},
	"5": {Name: "SQL注入漏洞", Url: "/api/v1/sqli"},
	"6": {Name: "SSRF服务端请求伪造漏洞", Url: "/api/v1/ssrf"},
	"7": {Name: "命令执行漏洞", Url: "/api/v1/rce"},
}
