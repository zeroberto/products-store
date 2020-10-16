package com.github.zeroberto.productsstore;

import com.github.zeroberto.productsstore.helper.ClientHelper;
import com.github.zeroberto.productsstore.helper.ServerHelper;
import com.github.zeroberto.productsstore.productslist.ProductsListRequest;
import com.github.zeroberto.productsstore.productslist.ProductsListResponse;
import lombok.SneakyThrows;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import static com.github.zeroberto.productsstore.helper.InfraHelper.checkDBConnection;
import static com.github.zeroberto.productsstore.helper.InfraHelper.checkMongoDBConnection;
import static com.github.zeroberto.productsstore.helper.InfraHelper.checkPortIsAvailable;
import static com.github.zeroberto.productsstore.helper.InfraHelper.downInfra;
import static com.github.zeroberto.productsstore.helper.InfraHelper.upInfra;
import static java.util.Arrays.asList;
import static org.junit.jupiter.api.Assertions.assertEquals;

class ApplicationTest {

  private static final String CLIENT_HOST = "localhost";
  private static final int CLIENT_PORT = 7773;
  private static final String DEFAULT_FILENAME = "docker-compose.yml";
  private static final int DISCOUNT_CALCULATOR_SERVER_PORT = 57772;
  private static final int USER_DATA_SERVER_PORT = 57771;
  private static final String PRODUCT_DB_HOST = "localhost";
  private static final int PRODUCT_DB_PORT = 27017;
  private static final String PRODUCT_DB_USER = "test";
  private static final String PRODUCT_DB_PASS = "test";
  private static final String PRODUCT_DB_REPO = "admin";
  private static final String USER_DB_DATASOURCE = "jdbc:postgresql://localhost:65432/user_info?sslmode=disable";
  private static final String USER_DB_USER = "test";
  private static final String USER_DB_PASS = "test";
  private static final String DEFAULT_LOCALTIME_ENV = "DEFAULT_LOCALTIME";
  private static final String DEFAULT_LOCALTIME_ENV_BLACKFRIDAY_VALUE = "2020-11-25T15:04:05Z";
  private static final String DEFAULT_LOCALTIME_ENV_BIRTHDAY_VALUE = "2020-01-01T15:04:05Z";

  @BeforeEach
  void setUp() {
    downInfra(DEFAULT_FILENAME);
  }

  @AfterEach
  void teardown() {
    downInfra(DEFAULT_FILENAME);
  }

  @Test
  void testCalculateDiscountWhenIsBlackFridayThenDiscountEqualToTenPercent() {
    downInfra(DEFAULT_FILENAME);
    upInfra(DEFAULT_FILENAME, Map.of(DEFAULT_LOCALTIME_ENV, DEFAULT_LOCALTIME_ENV_BLACKFRIDAY_VALUE));
    new Thread(ServerHelper::initServer).start();
    validateInfra(true);

    ClientHelper.createClient(CLIENT_HOST, CLIENT_PORT, stub -> {
      final List<ProductsListResponse> expected = asList(
        buildProduct1()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.1f)
            .setValueInCents(12)
            .build())
          .build(),
        buildProduct2()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.1f)
            .setValueInCents(140)
            .build())
          .build(),
        buildProduct3()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.1f)
            .setValueInCents(50)
            .build())
          .build()
      );

      final ProductsListRequest request = ProductsListRequest.newBuilder()
        .setUserId(1L)
        .build();

      final List<ProductsListResponse> got = new ArrayList<>();
      stub.listProducts(request).forEachRemaining(got::add);

      assertEquals(expected, got);
    });
  }

  @Test
  void testCalculateDiscountWhenIsBirthdayThenDiscountEqualToFivePercent() {
    downInfra(DEFAULT_FILENAME);
    upInfra(DEFAULT_FILENAME, Map.of(DEFAULT_LOCALTIME_ENV, DEFAULT_LOCALTIME_ENV_BIRTHDAY_VALUE));
    new Thread(ServerHelper::initServer).start();
    validateInfra(true);

    ClientHelper.createClient(CLIENT_HOST, CLIENT_PORT, stub -> {
      final List<ProductsListResponse> expected = asList(
        buildProduct1()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.05f)
            .setValueInCents(6)
            .build())
          .build(),
        buildProduct2()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.05f)
            .setValueInCents(70)
            .build())
          .build(),
        buildProduct3()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.05f)
            .setValueInCents(25)
            .build())
          .build()
      );

      final ProductsListRequest request = ProductsListRequest.newBuilder()
        .setUserId(1L)
        .build();

      final List<ProductsListResponse> got = new ArrayList<>();
      stub.listProducts(request).forEachRemaining(got::add);

      assertEquals(expected, got);
    });
  }

  @Test
  void testCalculateDiscountWhenIsBirthdayAndBlackFridayThenDiscountEqualToTenPercent() {
    downInfra(DEFAULT_FILENAME);
    upInfra(DEFAULT_FILENAME, Map.of(DEFAULT_LOCALTIME_ENV, DEFAULT_LOCALTIME_ENV_BLACKFRIDAY_VALUE));
    new Thread(ServerHelper::initServer).start();
    validateInfra(true);

    ClientHelper.createClient(CLIENT_HOST, CLIENT_PORT, stub -> {
      final List<ProductsListResponse> expected = asList(
        buildProduct1()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.1f)
            .setValueInCents(12)
            .build())
          .build(),
        buildProduct2()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.1f)
            .setValueInCents(140)
            .build())
          .build(),
        buildProduct3()
          .setDiscount(ProductsListResponse.Discount.newBuilder()
            .setPct(0.1f)
            .setValueInCents(50)
            .build())
          .build()
      );

      final ProductsListRequest request = ProductsListRequest.newBuilder()
        .setUserId(2L)
        .build();

      final List<ProductsListResponse> got = new ArrayList<>();
      stub.listProducts(request).forEachRemaining(got::add);

      assertEquals(expected, got);
    });
  }

  @Test
  void testListProductsWhenIsNotBirthdayNorBlackFridayThenDiscountIsEmpty() {
    downInfra(DEFAULT_FILENAME);
    upInfra(DEFAULT_FILENAME, Map.of(DEFAULT_LOCALTIME_ENV, "2020-03-20T00:01:02Z"));
    new Thread(ServerHelper::initServer).start();
    validateInfra(true);

    ClientHelper.createClient(CLIENT_HOST, CLIENT_PORT, stub -> {
      final List<ProductsListResponse> expected = asList(
        buildProduct1()
          .setDiscount(ProductsListResponse.Discount.newBuilder().build())
          .build(),
        buildProduct2()
          .setDiscount(ProductsListResponse.Discount.newBuilder().build())
          .build(),
        buildProduct3()
          .setDiscount(ProductsListResponse.Discount.newBuilder().build())
          .build()
      );

      final ProductsListRequest request = ProductsListRequest.newBuilder()
        .setUserId(1L)
        .build();

      final List<ProductsListResponse> got = new ArrayList<>();
      stub.listProducts(request).forEachRemaining(got::add);

      assertEquals(expected, got);
    });
  }

  @Test
  void testListProductsWhenUsersDataServiceIsDownThenDiscountIsEmpty() {
    downInfra(DEFAULT_FILENAME);
    upInfra(DEFAULT_FILENAME, "product-db-test");

    new Thread(ServerHelper::initServer).start();
    validateInfra(false);

    ClientHelper.createClient(CLIENT_HOST, CLIENT_PORT, stub -> {
      final List<ProductsListResponse> expected = asList(
        buildProduct1().build(),
        buildProduct2().build(),
        buildProduct3().build()
      );

      final ProductsListRequest request = ProductsListRequest.newBuilder()
        .setUserId(1L)
        .build();

      final List<ProductsListResponse> got = new ArrayList<>();
      stub.listProducts(request).forEachRemaining(got::add);

      assertEquals(expected, got);
    });
  }

  @SneakyThrows
  private static void validateInfra(final boolean considerDiscountCalculatorService) {
    boolean productDbUp = false;
    boolean userInfoDbUp = !considerDiscountCalculatorService;
    boolean userDataServerUp = !considerDiscountCalculatorService;
    boolean discountCalculatorServerUp = !considerDiscountCalculatorService;
    boolean serverUp = false;

    while (!productDbUp || !userInfoDbUp || !userDataServerUp || !discountCalculatorServerUp || !serverUp) {
      if (!productDbUp) {
        productDbUp = checkMongoDBConnection(PRODUCT_DB_HOST, PRODUCT_DB_PORT,
          PRODUCT_DB_USER, PRODUCT_DB_PASS, PRODUCT_DB_REPO);
      }
      if (!userInfoDbUp) {
        userInfoDbUp = checkDBConnection(USER_DB_DATASOURCE, USER_DB_USER, USER_DB_PASS);
      }
      if (!userDataServerUp) {
        userDataServerUp = checkPortIsAvailable(DISCOUNT_CALCULATOR_SERVER_PORT);
      }
      if (!discountCalculatorServerUp) {
        discountCalculatorServerUp = checkPortIsAvailable(USER_DATA_SERVER_PORT);
      }
      if (!serverUp) {
        serverUp = checkPortIsAvailable(CLIENT_PORT);
      }
      Thread.sleep(1000);
    }
  }

  private ProductsListResponse.Builder buildProduct1() {
    return ProductsListResponse.newBuilder()
      .setId("5f4962fb3ff6e3f16dca574e")
      .setPriceInCents(120)
      .setTilte("Product 1")
      .setDescription("Test");
  }

  private ProductsListResponse.Builder buildProduct2() {
    return ProductsListResponse.newBuilder()
      .setId("5f4962fb3ff6e3f16dca574f")
      .setPriceInCents(1400)
      .setTilte("Product 2")
      .setDescription("Test");
  }

  private ProductsListResponse.Builder buildProduct3() {
    return ProductsListResponse.newBuilder()
      .setId("5f4962fb3ff6e3f16dca5750")
      .setPriceInCents(500)
      .setTilte("Product 3")
      .setDescription("Test");
  }
}
