package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const defaultPort = ":3001"
const defaultTimeout = 30 * time.Second

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r.RemoteAddr)

		log.Printf("Received request from IP: %s", ip)
		w.Write([]byte(ip))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort // default port if not specified
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on :" + port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting")
}

// getIP extracts the IP address from a given address string
func getIP(addr string) string {
	// Split host and port
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		// If split fails, it could be a bare IP address (especially IPv6)
		if strings.Contains(addr, ":") && !strings.Contains(addr, "[") {
			// Validate and return the IPv6 address without port
			if ip := net.ParseIP(addr); ip != nil {
				return ip.String()
			}
		}
		return addr
	}
	return host
}
