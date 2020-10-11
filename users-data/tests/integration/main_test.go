package integration

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	infra "github.com/zeroberto/integration-test-suite"
	"github.com/zeroberto/products-store/users-data/cmd"
	"github.com/zeroberto/products-store/users-data/pb/userinfo"
	"google.golang.org/grpc"
)

const (
	host                  string = "localhost"
	serverPort            string = "7777"
	userDBType            string = "postgres"
	userDBDSN             string = "postgres://test:test@localhost:65432/user_info?sslmode=disable"
	dockerComposeFileName string = "docker-compose.yml"
)

func TestFindUserInfo(t *testing.T) {
	expected := &userinfo.UserInfoResponse{
		Id:          1,
		FirstName:   "User",
		LastName:    "Test",
		DateOfBirth: 315532800000000000,
		CreatedAt:   1598918400000000000,
	}

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uireq := &userinfo.UserInfoRequest{}
	uireq.Id = 1

	got, err := client.GetUserInfo(ctx, uireq)
	if err != nil {
		t.Fatalf("GetUserInfo() failed, no errors was expected, but %v", err)
	}
	if !compareTo(expected, got) {
		t.Errorf("GetUserInfo() failed, expected %v, got %v", expected, got)
	}
}

func TestFindUserInfo_WhenUserInfoNotExists_ThenFailure(t *testing.T) {
	expected := "User not found for this id"

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uireq := &userinfo.UserInfoRequest{}
	uireq.Id = 2

	ui, got := client.GetUserInfo(ctx, uireq)
	if ui != nil {
		t.Fatalf("GetUserInfo() failed, no user info was expected, but %v", ui)
	}
	if !strings.Contains(got.Error(), expected) {
		t.Fatalf("GetUserInfo() failed, expected %s, but %v", expected, got)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	infra.DownInfra(dockerComposeFileName)
	infra.UpInfra(dockerComposeFileName)
	go initServer()
	validateStructure(dockerComposeFileName)
}

func teardown() {
	infra.DownInfra(dockerComposeFileName)
}

func initServer() {
	cmd.NewServer().Start()
}

func validateStructure(dockerComposeFileName string) {
	for {
		if infra.CheckPortIsOpen(host, serverPort) && infra.CheckDBConnection(userDBType, userDBDSN) == nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func createClient() (client userinfo.UserInfoClient, closeConn func()) {
	conn, err := grpc.Dial(host+":"+serverPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	closeConn = func() {
		conn.Close()
	}
	return userinfo.NewUserInfoClient(conn), closeConn
}

func compareTo(ui1 *userinfo.UserInfoResponse, ui2 *userinfo.UserInfoResponse) bool {
	switch {
	case ui1.Id != ui2.Id,
		ui1.FirstName != ui2.FirstName,
		ui1.LastName != ui2.LastName,
		ui1.DateOfBirth != ui2.DateOfBirth,
		ui1.CreatedAt != ui2.CreatedAt,
		ui1.UpdatedAt != ui2.UpdatedAt,
		ui1.DeactivatedAt != ui2.DeactivatedAt:
		return false
	}
	return true
}
