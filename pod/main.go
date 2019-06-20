package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var version = "PLACEHOLDER"
var podName string
var startTime string

func init() {
	podName = os.Getenv("HOSTNAME")
	startTime = time.Now().Format(time.RFC3339)
}

func dump(r *http.Request) string {
	host, port, _ := net.SplitHostPort(r.RemoteAddr)
	m := map[string]string{
		"method":          r.Method,
		"path":            r.URL.String(),
		"user_agent":      r.Header.Get("User-Agent"),
		"forwarded_for":   r.Header.Get("X-Forwarded-For"),
		"forwarded_proto": r.Header.Get("X-Forwarded-Proto"),
		"forwarded_port":  r.Header.Get("X-Forwarded-Port"),
		"remote_host":     host,
		"remote_port":     port,
		"pod_name":        podName,
		"start_time":      startTime,
		"current_time":    time.Now().Format(time.RFC3339),
		"pod_version":     version,
	}
	b, _ := json.Marshal(m)
	return string(b)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d := dump(r)
		fmt.Println(d)
		fmt.Fprintln(w, d)
	})

	addr := ":8080"
	fmt.Println("Start server with addr", addr)
	http.ListenAndServe(addr, nil)
}
