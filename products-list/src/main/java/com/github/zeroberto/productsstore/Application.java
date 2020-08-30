package com.github.zeroberto.productsstore;

import com.github.zeroberto.productsstore.cmd.Server;

public class Application {

  public static void main(String[] args) {
    Server.newServer().start(System.getProperty("profile"));
  }
}
