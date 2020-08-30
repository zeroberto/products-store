package com.github.zeroberto.productsstore.datastore;

import com.github.zeroberto.productsstore.cmd.grpc.client.DiscountGrpcClient;
import com.github.zeroberto.productsstore.datastore.network.DiscountDataStoreGrpc;
import com.github.zeroberto.productsstore.discountcalculator.DiscountCalculatorServiceGrpc;
import com.github.zeroberto.productsstore.discountcalculator.DiscountRequest;
import com.github.zeroberto.productsstore.discountcalculator.DiscountResponse;
import com.github.zeroberto.productsstore.driver.DiscountGrpcClientBuilder;
import com.github.zeroberto.productsstore.exceptions.DataStoreNetworkException;
import io.grpc.StatusRuntimeException;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class DiscountDataStoreTest {

  private DiscountCalculatorServiceGrpc.DiscountCalculatorServiceBlockingStub blockingStub;
  private DiscountDataStore discountDataStore;

  @BeforeEach
  void setUp() {
    blockingStub = mock(DiscountCalculatorServiceGrpc.DiscountCalculatorServiceBlockingStub.class);

    final var discountGrpcClientBuilder = mock(DiscountGrpcClientBuilder.class);
    final var discountGrpcClient = mock(DiscountGrpcClient.class);

    when(discountGrpcClient.getBlockingStub()).thenReturn(blockingStub);
    when(discountGrpcClientBuilder.build()).thenReturn(discountGrpcClient);

    discountDataStore = new DiscountDataStoreGrpc(discountGrpcClientBuilder);
  }

  @Test
  void getDiscountBy() {
    final String productId = "test";
    final long userId = 1L;
    final var discountRequest = DiscountResponse.newBuilder()
      .setPct(10f)
      .setValueInCents(10)
      .build();

    when(blockingStub.calculateDiscount(any(DiscountRequest.class)))
      .thenReturn(discountRequest);

    final var discount = discountDataStore.getDiscountBy(productId, userId);

    assert discount != null;

    assertAll("testGetDiscountBy",
      () -> assertEquals(discountRequest.getPct(), discount.getPct()),
      () -> assertEquals(discountRequest.getValueInCents(), discount.getValueInCents()));
  }

  @Test
  void getDiscountByWhenStatusRuntimeExceptionThenThrowsDataStoreNetworkException() {
    final String productId = "any";
    final long userId = 1L;

    when(blockingStub.calculateDiscount(any(DiscountRequest.class)))
      .thenThrow(mock(StatusRuntimeException.class));

    assertThrows(DataStoreNetworkException.class,
      () -> discountDataStore.getDiscountBy(productId, userId));
  }
}
