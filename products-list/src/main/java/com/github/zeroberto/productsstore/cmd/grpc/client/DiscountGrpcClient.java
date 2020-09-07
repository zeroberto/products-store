package com.github.zeroberto.productsstore.cmd.grpc.client;

import com.github.zeroberto.productsstore.discountcalculator.DiscountCalculatorGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import lombok.Getter;

import static io.grpc.ManagedChannelBuilder.forTarget;

@Getter
public final class DiscountGrpcClient {

  private final ManagedChannel channel;
  private final DiscountCalculatorGrpc.DiscountCalculatorBlockingStub blockingStub;
  private final DiscountCalculatorGrpc.DiscountCalculatorStub asyncStub;

  public DiscountGrpcClient(String host) {
    this(forTarget(host));
  }

  public DiscountGrpcClient(ManagedChannelBuilder<?> channelBuilder) {
    channel = channelBuilder.build();
    blockingStub = DiscountCalculatorGrpc.newBlockingStub(channel);
    asyncStub = DiscountCalculatorGrpc.newStub(channel);
  }
}
