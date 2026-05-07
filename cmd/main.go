package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/xcoulon/claw-device-pairing/internal/handlers"
	"github.com/xcoulon/claw-device-pairing/internal/k8s/client"
	"github.com/xcoulon/claw-device-pairing/internal/logger"
	"github.com/xcoulon/claw-device-pairing/internal/version"
)

var (
	port int

	rootCmd = &cobra.Command{
		Use:   "claw-device-pairing",
		Short: "Claw device pairing server",
		Long:  "Claw device pairing server with API endpoint for device pairing requests",
	}

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		Long:  "Start the HTTP server to accept device pairing requests",
		Run:   runServer,
	}
)

const uiBuildDir = "ui/dist"

func init() {
	// Add --port flag to serve command
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")

	// Add serve command to root
	rootCmd.AddCommand(serveCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	// Initialize logger with JSON output
	logger.Init()

	// Print version information
	slog.Info("claw-device-pairing starting", "commit", version.CommitHash, "build_time", version.BuildTime)

	// Validate required environment variables
	namespace := os.Getenv("NAMESPACE")
	clawInstance := os.Getenv("CLAW_INSTANCE")

	var missingVars []string
	if namespace == "" {
		missingVars = append(missingVars, "NAMESPACE")
	}
	if clawInstance == "" {
		missingVars = append(missingVars, "CLAW_INSTANCE")
	}

	if len(missingVars) > 0 {
		slog.Error("required environment variables not set", "missing", missingVars)
		panic(fmt.Sprintf("required environment variables not set: %v", missingVars))
	}

	// Validate port range
	if port < 1 || port > 65535 {
		slog.Error("invalid port number", "port", port, "valid_range", "1-65535")
		os.Exit(1)
	}

	// Validate UI build directory exists
	if _, err := os.Stat(uiBuildDir); os.IsNotExist(err) {
		slog.Error("UI build directory not found", "path", uiBuildDir)
		os.Exit(1)
	}

	// Initialize Kubernetes client manager
	k8sManager, err := client.NewManager()
	if err != nil {
		slog.Error("failed to initialize Kubernetes client", "error", err)
		os.Exit(1)
	}

	// Initialize handlers
	pairingHandler := handlers.NewPairingRequestsHandler(k8sManager)

	// Initialize Echo instance
	e := echo.New()

	// Configure CORS middleware (only needed in development)
	if os.Getenv("ENV") == "development" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:5173", "http://localhost:5174"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}

	// Add request logging middleware
	e.Use(middleware.RequestLogger())

	// Register API routes
	e.POST("/pairing-requests", pairingHandler.HandlePairDevice)
	e.GET("/pairing-requests/:id", pairingHandler.HandleGetPairingStatus)
	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// Serve UI from /pair-device path
	uiHandler := func(c *echo.Context) error {
		// Get the request path and strip the /pair-device prefix
		path := c.Request().URL.Path
		filePath := filepath.Join(uiBuildDir, path)

		// Check if file exists
		if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
			return c.File(filePath)
		}

		// For directories or missing files, serve index.html (SPA fallback)
		return c.File(filepath.Join(uiBuildDir, "index.html"))
	}

	// Register UI routes
	e.GET("/", uiHandler)
	e.GET("/*", uiHandler)

	// Create http.Server for graceful shutdown support
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	// Start server in a goroutine
	go func() {
		slog.Info("server started", "port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("error during shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("command execution failed", "error", err)
		os.Exit(1)
	}
}
