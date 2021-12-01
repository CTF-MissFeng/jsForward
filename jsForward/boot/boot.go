package boot

import (
	"jsForward/library/logger"

	"github.com/gogf/gf/frame/g"
)

// init 初始化Web
func init() {
	logger.InitLogs()
	server := g.Server()
	server.EnableHTTPS("ca.crt", "ca.key")
	server.SetHTTPSPort(443)
	server.SetPort(80)
}