package cmd

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/zeroberto/products-store/discount-calculator/cmd/grpc/grpcservice"
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/appcontainer"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/grpcfactory"
	pb "github.com/zeroberto/products-store/discount-calculator/pb/discountcalculator"
	"google.golang.org/grpc"
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
	s.initServer()
}

func (s *Server) initContainer() {
	var configPath string
	var profile string

	flag.StringVar(&configPath, "fconfigPath", defaultConfigPath, "application configuration file path")
	flag.StringVar(&profile, "fprofile", defaultProfile, "application profile name")

	flag.Parse()

	log.Printf("Application using %s profile", getLogProfile(profile))

	if err := s.Container.Initialize(getConfigFilePath(configPath, profile)); err != nil {
		log.Fatalf("Failed to initialize application container, err=%v", err)
	}
}

func (s *Server) initServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Container.GetAppConfig().ServerConfig.Port))
	if err != nil {
		log.Fatalf("Failed to init server, error=%v", err)
	}

	log.Printf("Server will start on port %d", s.Container.GetAppConfig().ServerConfig.Port)

	grpcServer := grpc.NewServer()
	pb.RegisterDiscountCalculatorService(grpcServer, pb.NewDiscountCalculatorService(s.initGrpcService()))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to init server, error=%v", err)
	}
}

func (s *Server) initGrpcService() *grpcservice.DiscountCalculatorGrpcService {
	gsf := &grpcfactory.GrpcServiceFactory{}
	return gsf.MakeDiscountCalculatorGrpcService(s.Container)
}

func getConfigFilePath(configPath string, profile string) string {
	formattedProfile := strings.Title(strings.ToLower(profile))
	return fmt.Sprintf(configFileNameFormat, configPath, formattedProfile)
}

// NewServer is responsible for creating a new instance of Server
func NewServer() *Server {
	s := new(Server)
	s.Container = appcontainer.NewContainer()
	return s
}

func getLogProfile(profile string) string {
	if profile != "" {
		return profile
	}
	return "default"
}
