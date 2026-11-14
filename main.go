package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	cfg := New()

	srv := run(cfg.Server.AppName, cfg.Server.AppVersion)

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_forecast",
		Description: "Get weather forecast for a location",
	}, getForecast)

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_alerts",
		Description: "Get weather alerts for a state",
	}, getAlerts)

	if err := srv.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
