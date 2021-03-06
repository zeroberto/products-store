package com.github.zeroberto.productsstore.datastore.network;

import com.github.zeroberto.productsstore.cmd.grpc.client.DiscountGrpcClient;
import com.github.zeroberto.productsstore.datastore.DiscountDataStore;
import com.github.zeroberto.productsstore.discountcalculator.DiscountRequest;
import com.github.zeroberto.productsstore.discountcalculator.DiscountResponse;
import com.github.zeroberto.productsstore.driver.DiscountGrpcClientBuilder;
import com.github.zeroberto.productsstore.exceptions.DataStoreNetworkException;
import com.github.zeroberto.productsstore.model.Discount;
import io.grpc.StatusRuntimeException;
import lombok.RequiredArgsConstructor;

import javax.annotation.Nonnull;
import java.util.concurrent.TimeUnit;

@RequiredArgsConstructor
public class DiscountDataStoreGrpc implements DiscountDataStore {

  private final DiscountGrpcClientBuilder discountGrpcClientBuilder;

  @Override
  public Discount getDiscountBy(String productId, long userId) {
    final DiscountRequest discountRequest = DiscountRequest.newBuilder()
      .setProductId(productId)
      .setUserId(userId)
      .build();

    final DiscountGrpcClient discountGrpcClient  = discountGrpcClientBuilder.build();

    try {
      return toDiscount(discountGrpcClient
        .getBlockingStub()
        .calculateDiscount(discountRequest));
    } catch (StatusRuntimeException e) {
      throw new DataStoreNetworkException(e);
    } finally {
      try {
        discountGrpcClient.getChannel().shutdownNow().awaitTermination(1, TimeUnit.SECONDS);
      } catch (InterruptedException e) {
        Thread.currentThread().interrupt();
        e.printStackTrace();
      }
    }
  }

  @Nonnull
  private Discount toDiscount(final DiscountResponse discountResponse) {
    return Discount.builder()
      .pct(discountResponse.getPct())
      .valueInCents(discountResponse.getValueInCents())
      .build();
  }
}
