package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"roraddons/tools/mcp_server/server"
)

func main() {
	var cfg Config
	flag.StringVar(&cfg.DocsRoot, "docs-root", defaultDocsRoot(), "Path to docs/war-api root")
	flag.StringVar(&cfg.FeedingRoot, "feeding-root", defaultFeedingRoot(), "Path to docs/platform/feeding root")
	flag.StringVar(&cfg.Transport, "transport", "stdio", "Transport: stdio or http")
	flag.StringVar(&cfg.ListenAddr, "listen", ":8091", "HTTP listen address when transport=http")
	flag.Parse()

	app, err := server.NewApp(cfg.DocsRoot)
	if err != nil {
		log.Fatalf("load docs store: %v", err)
	}
	app.SetFeedingRoot(cfg.FeedingRoot)

	switch cfg.Transport {
	case "stdio":
		if err := app.ServeStdio(context.Background(), os.Stdin, os.Stdout); err != nil && !errors.Is(err, context.Canceled) {
			log.Fatalf("stdio server: %v", err)
		}
	case "http":
		h := http.NewServeMux()
		h.HandleFunc("/mcp", app.HandleHTTP)
		srv := &http.Server{
			Addr:              cfg.ListenAddr,
			Handler:           h,
			ReadHeaderTimeout: 10 * time.Second,
		}
		log.Printf("mcp server listening on %s", cfg.ListenAddr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server: %v", err)
		}
	default:
		fmt.Fprintf(os.Stderr, "unsupported transport %q\n", cfg.Transport)
		os.Exit(2)
	}
}
