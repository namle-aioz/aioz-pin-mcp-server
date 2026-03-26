package main

import (
	"aioz-pin-mcp-server/handler"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/server"
)

var (
	version   string
	buildTime string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	port := os.Getenv("SERVER_PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/.well-known/openai-apps-challenge", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Z5w8kz0a-_LAJr3zLuHnzDDwGGxojeUGeTObylWGnOQ"))
	})

	mcpServer := server.NewMCPServer(
		"AIOZ Pin MCP Server",
		version,
	)

	handler.RegisterHandlers(mcpServer)

	sseServer := server.NewSSEServer(mcpServer)
	mux.Handle("/", sseServer)
	fmt.Printf("Starting SSE server on port %s\n", port)

	displayVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()
	if *displayVersion {
		msg := fmt.Sprintf("Version:\t%s\nBuild time:\t%s\n", version, buildTime)
		fmt.Print(msg)
		os.Exit(0)
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
