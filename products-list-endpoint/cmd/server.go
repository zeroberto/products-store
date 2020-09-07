package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/zeroberto/products-store/products-list-endpoint/container/appcontainer"
	"grpc.go4.org"

	"github.com/zeroberto/products-store/products-list-endpoint/container"
	"github.com/zeroberto/products-store/products-list-endpoint/handler"
)

const (
	configFileNameFormat = "%sapplication%s.yml"
	defaultConfigPath    = "resources/"
	defaultProfile       = ""
	serverHostFormat     = ":%d"
)

// Server is responsible for starting application services
type Server struct {
	Container container.Container
}

// Start is responsible for initializing the server settings and routes
func (s *Server) Start() {
	log.Println("Starting server...")

	s.initContainer()
	s.initRoutes()

	port := s.Container.GetAppConfig().ServerConfig.Port

	log.Printf("Server will start on port %d", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(serverHostFormat, port), nil))
}

// Stop is responsible for shutting down the server
func (s *Server) Stop() {
	conn, ok := s.Container.Get(container.ProductsListServiceConnKey)
	if ok {
		if err := conn.(*grpc.ClientConn).Close(); err != nil {
			log.Fatalf("Could not end connection to the Product Listing gRPC server, error=%v", err)
		}
	}
	log.Println("The server was stopped")
}

func (s *Server) initContainer() {
	var configPath string
	var profile string

	flag.StringVar(&configPath, "fconfigPath", defaultConfigPath, "application configuration file path")
	flag.StringVar(&profile, "fprofile", defaultProfile, "application profile name")

	flag.Parse()

	log.Printf("Application using %s profile", getLogProfile(profile))

	s.Container = new(appcontainer.AppContainer)
	if err := s.Container.Initialize(getConfigFilePath(configPath, profile)); err != nil {
		log.Fatalf("Failed to initialize application container, err=%v", err)
	}
}

func (s *Server) initRoutes() {
	rhh := &handler.RestHTTPHandler{Container: s.Container}

	http.HandleFunc(handler.ProductPath, rhh.ProductRootHandle)
}

func getConfigFilePath(configPath string, profile string) string {
	formattedProfile := strings.Title(strings.ToLower(profile))
	return fmt.Sprintf(configFileNameFormat, configPath, formattedProfile)
}

func getLogProfile(profile string) string {
	if profile != "" {
		return profile
	}
	return "default"
}
