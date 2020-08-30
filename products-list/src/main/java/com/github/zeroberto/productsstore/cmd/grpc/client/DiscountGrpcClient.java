package com.github.zeroberto.productsstore.cmd.grpc.client;

import com.github.zeroberto.productsstore.discountcalculator.DiscountCalculatorServiceGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import lombok.Getter;

import static io.grpc.ManagedChannelBuilder.forTarget;

@Getter
public final class DiscountGrpcClient {

  private final ManagedChannel channel;
  private final DiscountCalculatorServiceGrpc.DiscountCalculatorServiceBlockingStub blockingStub;
  private final DiscountCalculatorServiceGrpc.DiscountCalculatorServiceStub asyncStub;

  public DiscountGrpcClient(String host) {
    this(forTarget(host));
  }

  public DiscountGrpcClient(ManagedChannelBuilder<?> channelBuilder) {
    channel = channelBuilder.build();
    blockingStub = DiscountCalculatorServiceGrpc.newBlockingStub(channel);
    asyncStub = DiscountCalculatorServiceGrpc.newStub(channel);
  }
}
