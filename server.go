package main

import "github.com/modelcontextprotocol/go-sdk/mcp"

func run(appName, appVersion string) *mcp.Server {
	return mcp.NewServer(&mcp.Implementation{
		Name:    appName,
		Version: appVersion,
	}, nil)
}
