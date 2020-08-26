package com.github.zeroberto.productsstore;

import com.github.zeroberto.productsstore.cmd.Server;

public class Application {

  public static void main(String[] args) {
    Server.start(args.length > 0 ? args[0] : null);
  }
}
