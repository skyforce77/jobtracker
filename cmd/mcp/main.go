package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/skyforce77/jobtracker/providers"
)

func main() {
	s := server.NewMCPServer(
		"jobtracker",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Tool: list_providers
	listProvidersTool := mcp.NewTool("list_providers",
		mcp.WithDescription("List all available job providers/companies that can be searched"),
	)

	s.AddTool(listProvidersTool, listProvidersHandler)

	// Tool: search_jobs
	searchJobsTool := mcp.NewTool("search_jobs",
		mcp.WithDescription("Search for job offers from specified providers"),
		mcp.WithString("providers",
			mcp.Required(),
			mcp.Description("Comma-separated list of provider names (e.g., 'greenhouse,lever') or 'all' for all providers"),
		),
		mcp.WithString("query",
			mcp.Description("Optional search query to filter job titles (case-insensitive)"),
		),
		mcp.WithString("location",
			mcp.Description("Optional location filter (case-insensitive)"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of results to return (default: 50)"),
		),
	)

	s.AddTool(searchJobsTool, searchJobsHandler)

	// Tool: get_provider_jobs
	getProviderJobsTool := mcp.NewTool("get_provider_jobs",
		mcp.WithDescription("Get all jobs from a specific provider"),
		mcp.WithString("provider",
			mcp.Required(),
			mcp.Description("Provider name (e.g., 'greenhouse', 'lever', 'netflix')"),
		),
	)

	s.AddTool(getProviderJobsTool, getProviderJobsHandler)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func listProvidersHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	providerList := providers.GetProviders()

	type ProviderInfo struct {
		Name string `json:"name"`
	}

	var result []ProviderInfo
	seen := make(map[string]bool)

	for _, p := range providerList {
		name := getProviderName(p)
		if name != "" && !seen[name] {
			seen[name] = true
			result = append(result, ProviderInfo{Name: name})
		}
	}

	jsonResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(string(jsonResult)), nil
}

func searchJobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	providersArg := request.GetString("providers", "")
	query := request.GetString("query", "")
	location := request.GetString("location", "")
	limit := request.GetInt("limit", 50)

	query = strings.ToLower(query)
	location = strings.ToLower(location)

	var selectedProviders []providers.Provider

	if providersArg == "all" {
		selectedProviders = providers.GetProviders()
	} else {
		names := strings.Split(providersArg, ",")
		for _, name := range names {
			name = strings.TrimSpace(name)
			p := providers.ProviderFromName(name)
			if p != nil {
				selectedProviders = append(selectedProviders, p)
			}
		}
	}

	if len(selectedProviders) == 0 {
		return mcp.NewToolResultError("No valid providers found"), nil
	}

	var jobs []*providers.Job

	for _, p := range selectedProviders {
		if len(jobs) >= limit {
			break
		}

		p.RetrieveJobs(func(job *providers.Job) {
			if len(jobs) >= limit {
				return
			}

			// Apply filters
			if query != "" && !strings.Contains(strings.ToLower(job.Title), query) {
				return
			}
			if location != "" && !strings.Contains(strings.ToLower(job.Location), location) {
				return
			}

			jobs = append(jobs, job)
		})
	}

	jsonResult, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Found %d jobs:\n%s", len(jobs), string(jsonResult))), nil
}

func getProviderJobsHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	providerName := request.GetString("provider", "")

	p := providers.ProviderFromName(providerName)
	if p == nil {
		return mcp.NewToolResultError(fmt.Sprintf("Provider '%s' not found", providerName)), nil
	}

	var jobs []*providers.Job
	err := p.RetrieveJobs(func(job *providers.Job) {
		jobs = append(jobs, job)
	})

	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error retrieving jobs: %v", err)), nil
	}

	jsonResult, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Found %d jobs from %s:\n%s", len(jobs), providerName, string(jsonResult))), nil
}

func getProviderName(p providers.Provider) string {
	// Use reflection to get the provider name
	name := fmt.Sprintf("%T", p)
	// Remove package prefix "*providers."
	if idx := strings.LastIndex(name, "."); idx != -1 {
		name = name[idx+1:]
	}
	// Remove pointer prefix
	name = strings.TrimPrefix(name, "*")
	return name
}
