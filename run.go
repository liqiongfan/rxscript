package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"rxscript/ast"
	"time"
)

var (
	goNum       int
	debug       int
	author      bool
	showVersion bool
)

func NewOptimizedClient() *http.Client {
	// 1. 优化底层的 DialContext（控制 TCP 连接建立阶段）
	dialer := &net.Dialer{
		Timeout:   30 * time.Second, // 建立 TCP 连接的超时时间
		KeepAlive: 30 * time.Second, // TCP KeepAlive 探测间隔
	}

	// 2. 优化 Transport（控制连接池和传输层）
	transport := &http.Transport{
		Proxy:             http.ProxyFromEnvironment, // 支持系统代理
		DialContext:       dialer.DialContext,
		ForceAttemptHTTP2: true, // 强制尝试启用 HTTP/2（如果服务端支持）

		// 🚀 核心优化：连接池参数
		MaxIdleConns:        1000, // 全局最大空闲连接数（默认 100）
		MaxIdleConnsPerHost: 100,  // 每个 Host（域名/IP）最大空闲连接数（默认 2！高并发下必改）
		MaxConnsPerHost:     0,    // 每个 Host 总连接数限制，0 表示不限制（可根据业务限流调整）

		// ⏱️ 核心优化：超时与清理
		IdleConnTimeout:       90 * time.Second, // 空闲连接在连接池中的存活时间（默认 90s）
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue 响应超时
	}

	// 3. 构建 Client 并设置全局总超时
	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second, // 整个请求（含 Body 读取）的总超时
	}
}

func init() {
	http.DefaultClient = NewOptimizedClient()

	flag.IntVar(&goNum, "go", 10, "Go routines number limitation")
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.BoolVar(&author, "a", false, "show author")
	flag.BoolVar(&showVersion, "v", false, "show RXScript version")
}

func main() {
	flag.Parse()
	os.Args = flag.Args()
	if showVersion {
		log.Fatal("RXScript version 0.5.1")
	}
	if author {
		log.Fatal("RXScript Author: xeapplee@gmail.com(terry)")
	}
	if len(os.Args) < 1 {
		log.Fatal("No RXScript file specified")
	}
	ast.RunScript(os.Args[0], debug, goNum)
}
