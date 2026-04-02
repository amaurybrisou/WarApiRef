package main

import (
	"os"
	"path/filepath"
)

type Config struct {
	DocsRoot   string
	FeedingRoot string
	Transport  string
	ListenAddr string
}

func defaultDocsRoot() string {
	if fromEnv := os.Getenv("WAR_API_ROOT"); fromEnv != "" {
		return fromEnv
	}
	return filepath.Join("docs", "war-api")
}

func defaultFeedingRoot() string {
	if fromEnv := os.Getenv("WAR_API_FEEDING_ROOT"); fromEnv != "" {
		return fromEnv
	}
	return filepath.Join("docs", "platform", "feeding")
}
