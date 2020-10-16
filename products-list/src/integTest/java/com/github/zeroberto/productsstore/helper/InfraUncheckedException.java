package com.github.zeroberto.productsstore.helper;

public class InfraUncheckedException extends RuntimeException {

  public InfraUncheckedException(String message) {
    super(message);
  }

  public InfraUncheckedException(Throwable cause) {
    super(cause);
  }
}
