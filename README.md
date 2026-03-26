# AIOZ MCP Server

A [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server written in Go that integrates with [AIOZ Stream](https://aiozstream.network/) for media management and live streaming.

---

## Features

- **Media Statistics**: Count total videos and audios in an AIOZ Stream account.
- **Video Management**: List videos, search by name, and retrieve playback URLs.
- **Video Upload**: Upload video files to AIOZ Stream directly from a Google Drive link.
- **Live Streaming**: Create live stream keys for broadcasting.

---

## MCP Tools

---

### `count-total-media`

Get the total number of videos and audios in an AIOZ Stream account.

| Parameter   | Type   | Required | Description            |
| ----------- | ------ | -------- | ---------------------- |
| `publicKey` | string | Yes      | AIOZ Stream public key |
| `secretKey` | string | Yes      | AIOZ Stream secret key |

**Output example:**

```
AIOZ Stream Account Stats:
Videos: 42
Audios: 15
```

---

### `get-video-url`

Search for a video by name and retrieve all its playback and asset URLs.

| Parameter   | Type   | Required | Description                                   |
| ----------- | ------ | -------- | --------------------------------------------- |
| `publicKey` | string | Yes      | AIOZ Stream public key                        |
| `secretKey` | string | Yes      | AIOZ Stream secret key                        |
| `videoName` | string | Yes      | Name (or partial name) of the video to search |

**Output example:**

```json
{
  "EmbededURL": "https://...",
  "Mp4URL": "https://...",
  "Thumbnail": "https://...",
  "SourceURL": "https://..."
}
```

---

### `get-list-video`

Get a list of all videos in an AIOZ Stream account.

| Parameter   | Type   | Required | Description pen        |
| ----------- | ------ | -------- | ---------------------- |
| `publicKey` | string | Yes      | AIOZ Stream public key |
| `secretKey` | string | Yes      | AIOZ Stream secret key |

**Output example:**

```json
[
  {
    "MediaID": "abc123",
    "Name": "My Video",
    "Size": 104857600,
    "Duration": 120.5,
    "CreatedAt": "2024-01-15T10:00:00Z"
  }
]
```

---

### `upload-video`

Upload a video from a Google Drive link to an AIOZ Stream account. The server automatically:

1. Validates the Google Drive link format
2. Converts it to a direct download URL
3. Downloads the video file
4. Uploads it to AIOZ Stream with the specified title

| Parameter   | Type   | Required | Description                                      |
| ----------- | ------ | -------- | ------------------------------------------------ |
| `publicKey` | string | Yes      | AIOZ Stream public key                           |
| `secretKey` | string | Yes      | AIOZ Stream secret key                           |
| `videoLink` | string | Yes      | Google Drive shareable link to the video file    |
| `title`     | string | Yes      | Title/name for the uploaded video on AIOZ Stream |

**Supported Google Drive Link Format Or Link Can Download:**

```
https://drive.google.com/file/d/{FILE_ID}/view?usp=sharing
https://api.aiozstream.network/api/media/video/mp4
```

**Output example:**

```
Video uploaded successfully
```

---

### `create-key-live`

Create a live stream key in an AIOZ Stream account.

| Parameter   | Type   | Required | Description                   |
| ----------- | ------ | -------- | ----------------------------- |
| `publicKey` | string | Yes      | AIOZ Stream public key        |
| `secretKey` | string | Yes      | AIOZ Stream secret key        |
| `keyName`   | string | Yes      | Display name for the live key |

**Output example:**

```
Key created successfully to AIOZ Stream with name: my-stream-key
```

---

## HTTP Endpoints

### `GET /ping`

Health check endpoint.

**Response:** `pong`

---

## Installation

### Prerequisites

- [Go](https://go.dev/doc/install) 1.24.0 or later
- A valid [AIOZ Stream](https://stream.aioz.io/) account with API credentials

### Build

```bash
git clone <repository-url>
cd MCP-Server-Test
go build -o aioz-mcp-server .
```

### Environment Variables

Create a `.env` file in the project root (optional):

```env
SERVER_PORT=8087
```

If `SERVER_PORT` is not set, the server defaults to port `8087`.

---

## Running the Server

```bash
go run .
```

Or using the compiled binary:

```bash
./aioz-mcp-server
```

The server will start on `http://localhost:8087` by default. MCP clients connect via SSE at `http://localhost:8087/`.

---

## Configuration

---

### Cursor

Cursor is a code editor with native MCP support. Follow these steps:

**Step 1 — Open the MCP config file**

- Press `Ctrl + Shift + P` (Windows/Linux) or `Cmd + Shift + P` (macOS)
- Type **"Open MCP Settings"** → press Enter
- The `mcp.json` file will open (usually located at `~/.cursor/mcp.json`)

**Step 2 — Add the server to the config**

Paste the following into the file (if the file already has content, just add the `aioz-mcp-server` entry inside `mcpServers`):

```json
{
  "mcpServers": {
    "aioz-mcp-server": {
      "url": "http://localhost:8087/sse"
    }
  }
}
```

**Step 3 — Save and verify**

- Save the file (`Ctrl + S`)
- Open **Cursor Chat** (`Ctrl + L`) and try: _"Count the total number of videos in my AIOZ Stream account"_
- Cursor will automatically call the `count-total-media` tool

| Step 1                                          | Step 2                                          | Step 3                                          |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| <img src="docs/images/cursor1.png" width="300"> | <img src="docs/images/cursor2.png" width="300"> | <img src="docs/images/cursor3.png" width="300"> |

---

### Claude CLI

Claude CLI supports MCP through its own config file.

**Step 1 — Open the Claude CLI**

- Launch the Claude CLI from your terminal.

**Step 2 — Set up the MCP server**

- Use the following command to add your MCP server:
  `claude mcp add --transport sse <server-name> <server-url>`

  Parameters:
  - <server-name>: The name you want to assign to your MCP server (e.g., mcp-aioz-stream)
  - <server-url>: The endpoint URL of your MCP server (e.g., http://localhost:8087/sse)

  Example: claude mcp add --transport sse mcp-aioz-stream http://localhost:8087/sse

**Step 3 — Check your MCP server**

- Use the following command to list tool in your MCP server:
  `claude mcp list`

| Step 1                                              | Step 2                                              |
| --------------------------------------------------- | --------------------------------------------------- |
| <img src="docs/images/claudeCLI_1.png" width="350"> | <img src="docs/images/claudeCLI_2.png" width="350"> |

---

### ChatGPT

**Step 1 — Enable developer mode**

Make sure developer mode is enabled in your ChatGPT account settings before proceeding.

**Step 2 — Create a connector**

- In ChatGPT, navigate to **Settings → Apps → Advanced settings -> Create App**
- Fill in the connector metadata:
  - **Name** — a user-facing title such as `AIOZ MCP Server`
  - **Description** — explain what the connector does and when to use it. The model uses this text during discovery
  - **MCP Server URL** — the `/sse` endpoint of your server: `http://localhost:8087/sse`
  - **Authentication** — choose `No auth`
  - **Risk and policy** — checked box `I understand and want to continue`

**Step 3 — Verify**

- Click **Create**
- If the connection succeeds, ChatGPT will display a list of tools your server advertises
- If it fails, debug your app with [MCP Inspector](https://modelcontextprotocol.io/docs/tools/inspector) or the API Playground

**Step 4 — Use MCP server**

- Create a new chat
- Click the `+` button on the left side of the input box
- Select More, then choose your MCP server
- Enter a prompt to test a tool

| Step 1                                            | Step 2                                            | Step 3                                            |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| <img src="docs/images/chatgpt_1.png" width="300"> | <img src="docs/images/chatgpt_2.png" width="300"> | <img src="docs/images/chatgpt_3.png" width="300"> |

---

---

### Visual Studio Code (VS Code)

VS Code supports MCP through **GitHub Copilot** (Agent mode must be enabled).

**Step 1 — Open the config file**

- Press `Ctrl + Shift + P` → type `Open User Configuration` → `.vscode/mcp.json`

**Step 2 — Add the MCP config**

```json
{
  "servers": {
    "aioz-mcp-server": {
      "type": "sse",
      "url": "http://localhost:8087/sse"
    }
  }
}
```

**Step 3 — Enable Agent Mode in Copilot Chat**

- Open **Copilot Chat** (`Ctrl + Alt + I`)
- At the top of the chat, switch to **Agent** mode (instead of Ask/Edit)
- Try: _"List all videos in my AIOZ Stream account"_

| Step 1                                          | Step 2                                          |
| ----------------------------------------------- | ----------------------------------------------- |
| <img src="docs/images/vscode1.png" width="350"> | <img src="docs/images/vscode2.png" width="350"> |

---

## Docker

A `Dockerfile` and `docker-compose.local.yml` are provided for containerized deployment.

```bash
docker compose -f docker-compose.local.yml up --build
```

---

## Project Structure

```
.
├── main.go                        # Entry point, MCP server and tool registration
├── handler/
│   ├── aiozstream_handler.go      # MCP tool handlers for AIOZ Stream
│   ├── register_handler.go        # Register MCP tool handlers for AIOZ Stream
├── tool/
│   ├── aiozstream.go              # AIOZ Stream API client logic
├── pkg/
│   └── cache
│         └── cache                # Init caching variable
├── model/
│   └── client_upload_model.go     # Shared data models
├── util/
│   ├── drive.go                   # Drive func util
├── constant/
│   ├── file_constant.go           # Variable constant
├── go.mod
├── Dockerfile
└── docker-compose.local.yml
```
