package com.github.zeroberto.productsstore.helper;

import com.github.zeroberto.productsstore.Application;
import lombok.NoArgsConstructor;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class ServerHelper {

  public static void initServer() {
    Application.main(new String[]{});
  }
}
