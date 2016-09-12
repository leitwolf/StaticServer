package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// 监听的端口
var port = 5000

// startServer 开启web服务
func startServer(port int) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Listening on port: " + strconv.Itoa(port))
	log.Println("Open url: http://localhost:" + strconv.Itoa(port))
	// 当前程序所在目录
	f, err := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(f)
	path = filepath.Dir(path)
	// 静态文件
	http.Handle("/", http.FileServer(http.Dir(path)))
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func main() {
	port1 := flag.Int("port", port, "Web server port >1024")
	flag.Parse()
	if *port1 <= 1024 {
		log.Fatal("Web port must >1024")
	} else {
		port = *port1
	}

	startServer(port)
}
