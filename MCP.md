# JobTracker MCP Server

This MCP server allows Claude Code and Claude Desktop to search for job offers.

## Build

```bash
go build -o jobtracker-mcp ./cmd/mcp
```

## Configuration

### Claude Code

Add to your `~/.claude/settings.json`:

```json
{
  "mcpServers": {
    "jobtracker": {
      "command": "/path/to/jobtracker-mcp"
    }
  }
}
```

### Claude Desktop

Add to your `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "jobtracker": {
      "command": "/path/to/jobtracker-mcp"
    }
  }
}
```

## Available Tools

### list_providers

List all available job providers/companies.

### search_jobs

Search for job offers with filters.

Parameters:
- `providers` (required): Comma-separated provider names or "all"
- `query` (optional): Filter job titles
- `location` (optional): Filter by location
- `limit` (optional): Max results (default: 50)

### get_provider_jobs

Get all jobs from a specific provider.

Parameters:
- `provider` (required): Provider name (e.g., "greenhouse", "lever")

## Example Usage in Claude

"Search for engineering jobs at Greenhouse"
"List all available job providers"
"Find remote jobs containing 'data' in the title"
