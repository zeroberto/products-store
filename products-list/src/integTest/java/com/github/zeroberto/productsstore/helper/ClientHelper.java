package com.github.zeroberto.productsstore.helper;

import com.github.zeroberto.productsstore.productslist.ProductsListGrpc;
import io.grpc.ManagedChannel;
import lombok.NoArgsConstructor;

import java.util.concurrent.TimeUnit;

import static io.grpc.ManagedChannelBuilder.forTarget;
import static lombok.AccessLevel.PRIVATE;
import static org.junit.jupiter.api.Assertions.fail;

@NoArgsConstructor(access = PRIVATE)
public final class ClientHelper {

  public static void createClient(final String host, final int port, final ClientRunnable runnable) {
    final ManagedChannel channel = forTarget(host + ":" + port)
      .usePlaintext()
      .build();
    try {
      runnable.run(ProductsListGrpc.newBlockingStub(channel));
    } finally {
      try {
        channel.shutdownNow().awaitTermination(5, TimeUnit.SECONDS);
      } catch (InterruptedException e) {
        e.printStackTrace();
        Thread.currentThread().interrupt();
        fail();
      }
    }
  }

  public interface ClientRunnable {
    void run(ProductsListGrpc.ProductsListBlockingStub stub);
  }
}
