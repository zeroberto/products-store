package integration

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/container"

	"github.com/zeroberto/products-store/discount-calculator/chrono"
	"github.com/zeroberto/products-store/discount-calculator/chrono/provider"

	infra "github.com/zeroberto/integration-test-suite"
	"github.com/zeroberto/products-store/discount-calculator/cmd"
	"github.com/zeroberto/products-store/discount-calculator/pb/discountcalculator"
	"google.golang.org/grpc"
)

const (
	host                  string = "localhost"
	serverPort            string = "7777"
	productDBDSN          string = "mongodb://test:test@localhost:65017/products?authSource=admin"
	userDBType            string = "postgres"
	userDBDSN             string = "postgres://test:test@localhost:65432/user_info?sslmode=disable"
	userServiceHost       string = "localhost"
	userServicePort       string = "57773"
	dockerComposeFileName string = "docker-compose.yml"
)

var clock time.Time = time.Now()

func TestCalculateDiscount_WhenIsBlackFriday_ThenDiscountEqualToTenPercent(t *testing.T) {
	expected := &discountcalculator.DiscountResponse{
		ProductId:    "5f4962fb3ff6e3f16dca574e",
		UserId:       1,
		Pct:          0.10,
		ValueInCents: 12,
	}

	clock = time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC)

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dreq := &discountcalculator.DiscountRequest{}
	dreq.ProductId = "5f4962fb3ff6e3f16dca574e"
	dreq.UserId = 1

	got, err := client.CalculateDiscount(ctx, dreq)
	if err != nil {
		t.Fatalf("CalculateDiscount() failed, no errors was expected, but %v", err)
	}
	if !compareTo(expected, got) {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestCalculateDiscount_WhenIsBirthday_ThenDiscountEqualToFivePercent(t *testing.T) {
	expected := &discountcalculator.DiscountResponse{
		ProductId:    "5f4962fb3ff6e3f16dca574f",
		UserId:       1,
		Pct:          0.05,
		ValueInCents: 70,
	}

	clock = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dreq := &discountcalculator.DiscountRequest{}
	dreq.ProductId = "5f4962fb3ff6e3f16dca574f"
	dreq.UserId = 1

	got, err := client.CalculateDiscount(ctx, dreq)
	if err != nil {
		t.Fatalf("CalculateDiscount() failed, no errors was expected, but %v", err)
	}
	if !compareTo(expected, got) {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestCalculateDiscount_WhenIsBirthdayAndBlackFriday_ThenDiscountEqualToTenPercent(t *testing.T) {
	expected := &discountcalculator.DiscountResponse{
		ProductId:    "5f4962fb3ff6e3f16dca574e",
		UserId:       2,
		Pct:          0.10,
		ValueInCents: 12,
	}

	clock = time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC)

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dreq := &discountcalculator.DiscountRequest{}
	dreq.ProductId = "5f4962fb3ff6e3f16dca574e"
	dreq.UserId = 2

	got, err := client.CalculateDiscount(ctx, dreq)
	if err != nil {
		t.Fatalf("CalculateDiscount() failed, no errors was expected, but %v", err)
	}
	if !compareTo(expected, got) {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestCalculateDiscount_WhenIsNotBirthdayNorBlackFriday_ThenDiscountEqualToZeroPercent(t *testing.T) {
	expected := &discountcalculator.DiscountResponse{
		ProductId:    "5f4962fb3ff6e3f16dca5750",
		UserId:       1,
		Pct:          0,
		ValueInCents: 0,
	}

	clock = time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC)

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dreq := &discountcalculator.DiscountRequest{}
	dreq.ProductId = "5f4962fb3ff6e3f16dca5750"
	dreq.UserId = 1

	got, err := client.CalculateDiscount(ctx, dreq)
	if err != nil {
		t.Fatalf("CalculateDiscount() failed, no errors was expected, but %v", err)
	}
	if !compareTo(expected, got) {
		t.Errorf("CalculateDiscount() failed, expected %v, got %v", expected, got)
	}
}

func TestCalculateDiscount_WhenProductNotExists_ThenFailure(t *testing.T) {
	expected := "Could not retrieve product: mongo: no documents in result"

	clock = time.Now()

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dreq := &discountcalculator.DiscountRequest{}
	dreq.ProductId = "5f7bc4fd02c3d80006663ad1"
	dreq.UserId = 1

	d, got := client.CalculateDiscount(ctx, dreq)
	if d != nil {
		t.Fatalf("CalculateDiscount() failed, no discount was expected, but %v", d)
	}
	if d != nil && d.UserId != 0 {
		t.Fatalf("CalculateDiscount() failed, no product was expected, but %v", d.UserId)
	}
	if !strings.Contains(got.Error(), expected) {
		t.Fatalf("CalculateDiscount() failed, expected %s, but %v", expected, got)
	}
}

func TestCalculateDiscount_WhenUserInfoNotExists_ThenFailure(t *testing.T) {
	expected := "User not found for this id"

	clock = time.Now()

	client, closeConn := createClient()
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dreq := &discountcalculator.DiscountRequest{}
	dreq.ProductId = "5f4962fb3ff6e3f16dca574e"
	dreq.UserId = 3

	d, got := client.CalculateDiscount(ctx, dreq)
	if d != nil {
		t.Fatalf("CalculateDiscount() failed, no discount was expected, but %v", d)
	}
	if d != nil && d.UserId != 0 {
		t.Fatalf("CalculateDiscount() failed, no user info was expected, but %v", d.UserId)
	}
	if !strings.Contains(got.Error(), expected) {
		t.Fatalf("CalculateDiscount() failed, expected %s, but %v", expected, got)
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
	server := cmd.NewServer()
	server.Container.Put(container.TimeStampKey, &timeStampStub{})
	server.Start()
}

func validateStructure(dockerComposeFileName string) {
	var productDBUp, userDBUp, userServiceUp bool
	for {
		if !productDBUp {
			productDBUp = infra.CheckMongoDBConnection(productDBDSN) == nil
		}
		if !userDBUp {
			err := infra.CheckDBConnection(userDBType, userDBDSN)
			userDBUp = err == nil
		}
		if !userServiceUp {
			userServiceUp = infra.CheckPortIsOpen(userServiceHost, userServicePort)
		}
		if productDBUp && userServiceUp && userDBUp {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func createClient() (client discountcalculator.DiscountCalculatorClient, closeConn func()) {
	conn, err := grpc.Dial(host+":"+serverPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	closeConn = func() {
		conn.Close()
	}
	return discountcalculator.NewDiscountCalculatorClient(conn), closeConn
}

func compareTo(dc1 *discountcalculator.DiscountResponse, dc2 *discountcalculator.DiscountResponse) bool {
	switch {
	case dc1.ProductId != dc2.ProductId,
		dc1.UserId != dc2.UserId,
		dc1.Pct != dc2.Pct,
		dc1.ValueInCents != dc2.ValueInCents:
		return false
	}
	return true
}

var tsi chrono.TimeStamp = &provider.TimeStampImpl{}

type timeStampStub struct{}

func (tss *timeStampStub) GetCurrentTime() time.Time {
	return clock
}

func (tss *timeStampStub) GetBlackFridayDay() int {
	return tsi.GetBlackFridayDay()
}

func (tss *timeStampStub) GetBlackFridayMonth() time.Month {
	return tsi.GetBlackFridayMonth()
}

func (tss *timeStampStub) GetTimeByNanoSeconds(nanos int64) time.Time {
	return tsi.GetTimeByNanoSeconds(nanos)
}
