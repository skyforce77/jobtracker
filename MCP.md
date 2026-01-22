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

### Job Search Tools

#### list_providers
List all available job providers/companies.

#### search_jobs
Search for job offers with filters.

Parameters:
- `providers` (required): Comma-separated provider names or "all"
- `query` (optional): Filter job titles
- `location` (optional): Filter by location
- `limit` (optional): Max results (default: 50)

#### get_provider_jobs
Get all jobs from a specific provider.

Parameters:
- `provider` (required): Provider name (e.g., "greenhouse", "lever")

### Resume Tools (JSON Resume format)

#### load_resume
Load a JSON Resume from a file path.

Parameters:
- `path` (required): Path to the JSON Resume file

#### set_resume
Set resume data directly from JSON content.

Parameters:
- `json` (required): JSON Resume content as a string

#### get_resume
Get the currently loaded resume.

#### match_jobs
Find jobs that match the loaded resume skills.

Parameters:
- `providers` (required): Comma-separated provider names or "all"
- `limit` (optional): Max results (default: 20)
- `min_score` (optional): Minimum match score 0-1 (default: 0.1)

## JSON Resume Format

The MCP server uses the [JSON Resume](https://jsonresume.org/) open standard.
See `examples/resume.json` for a sample resume file.

## Example Usage in Claude

"Search for engineering jobs at Greenhouse"
"List all available job providers"
"Find remote jobs containing 'data' in the title"
"Load my resume from ~/resume.json"
"Find jobs that match my resume skills"
"What jobs at Greenhouse match my profile?"
