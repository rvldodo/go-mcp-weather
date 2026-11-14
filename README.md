# mcp-weather

A [Model Context Protocol](https://modelcontextprotocol.io) (MCP) server, written in Go, that serves US weather data from the [National Weather Service API](https://www.weather.gov/documentation/services-web-api) (`api.weather.gov`). The server talks over **stdio**, so an MCP client such as Claude Desktop or Claude Code starts it as a subprocess.

## Tools

| Tool           | Input                                           | Description                                        |
| -------------- | ----------------------------------------------- | -------------------------------------------------- |
| `get_forecast` | `latitude` (number), `longitude` (number)       | Returns the next 5 forecast periods for a location |
| `get_alerts`   | `state` (string, two-letter US code, e.g. `CA`) | Returns the active weather alerts for a US state   |

> The NWS API only covers the United States and its territories. Coordinates outside it return "Unable to fetch forecast data for this location."

## Project layout

| File         | Purpose                                                        |
| ------------ | -------------------------------------------------------------- |
| `main.go`    | Entry point: builds the server, registers tools, runs on stdio |
| `server.go`  | Creates the `mcp.Server` with the app name and version         |
| `config.go`  | Loads configuration from environment variables                 |
| `const.go`   | NWS base URL and User-Agent                                    |
| `usecase.go` | Tool handlers (`getForecast`, `getAlerts`)                     |
| `helper.go`  | Generic NWS HTTP client and output formatting                  |
| `dto.go`     | Request and response types for NWS and the tool inputs         |
| `utils.go`   | Environment variable helper                                    |

## Requirements

- **Go 1.27.3 or newer** (as set in `go.mod`). Check with `go version`.
- Internet access to `https://api.weather.gov` (no API key needed).

## Build

```bash
# 1. Clone and enter the project
git clone https://github.com/rvldodo/mcp-weather.git
cd mcp-weather

# 2. Download dependencies
go mod download

# 3. Build the binary (writes ./weather)
go build -o weather .
```

To check the code before building:

```bash
go vet ./...
```

### Cross-compiling

Go can build for another platform by setting `GOOS` and `GOARCH`:

```bash
GOOS=linux   GOARCH=amd64 go build -o weather-linux-amd64 .
GOOS=darwin  GOARCH=arm64 go build -o weather-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -o weather.exe .
```

## Configuration

The server reads these environment variables (see `.env.example`):

| Variable      | Default   | Description                            |
| ------------- | --------- | -------------------------------------- |
| `APP_NAME`    | `weather` | Server name reported to MCP clients    |
| `APP_VERSION` | `1.0`     | Server version reported to MCP clients |

The program does **not** load `.env` files by itself. Set the variables in your shell, or in the `env` block of your MCP client config (shown below). Both are optional.

## Run

The server communicates over stdin/stdout, so running it directly just waits for MCP messages:

```bash
./weather
```

You can test it interactively with the [MCP Inspector](https://github.com/modelcontextprotocol/inspector):

```bash
npx @modelcontextprotocol/inspector ./weather
```

### Use with Claude Code

```bash
claude mcp add weather -- /absolute/path/to/mcp-weather/weather
```

### Use with Claude Desktop

Add this to `claude_desktop_config.json` (on macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "weather": {
      "command": "/absolute/path/to/mcp-weather/weather",
      "env": {
        "APP_NAME": "mcp-weather",
        "APP_VERSION": "1.0"
      }
    }
  }
}
```

Restart Claude Desktop, then ask something like _"What's the forecast for Sacramento?"_ or _"Are there any weather alerts in TX?"_
