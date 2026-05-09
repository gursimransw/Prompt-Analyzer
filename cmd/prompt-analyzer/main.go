package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/gursimransw/prompt-analyzer/internal"
	"github.com/gursimransw/prompt-analyzer/internal/http/handlers/prompts"
	loader "github.com/gursimransw/prompt-analyzer/internal/utils"
)

func main() {

	cfg := config.MustLoad()
	//Mustload function is a part of config package , this is responsible for loading the configuration variables into this code from the config.yaml file
	//present in - config/server/config.yaml

	// if err != nil {
	// 	log.Fatal(err)
	// }

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	PromptLibraryPath := cfg.PromptLibrary

	PatternConfig, err := loader.LoadPatterns(PromptLibraryPath)
	if err != nil {
		panic(err)
	}

	//Setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/prompt-analyzer/detect", prompts.PromptAnalyzer(PatternConfig).ServeHTTP)

	//Done Channel

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//Setup Server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("address", cfg.Addr))

	go func() {

		fmt.Printf("Students API Server Started %s", cfg.HTTPServer.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}

	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown gracefully !!!")

}
