package com.github.zeroberto.productsstore.usecase.impl;

import com.github.zeroberto.productsstore.datastore.DiscountDataStore;
import com.github.zeroberto.productsstore.model.Discount;
import com.github.zeroberto.productsstore.usecase.DiscountUseCase;
import lombok.RequiredArgsConstructor;
import lombok.extern.java.Log;

import static java.lang.String.format;
import static java.util.logging.Level.WARNING;

@Log
@RequiredArgsConstructor
public class DiscountUseCaseImpl implements DiscountUseCase {

  private final DiscountDataStore discountDataStore;

  @Override
  public Discount calculateDiscount(String productId, long userId) {
    try {
      return discountDataStore.getDiscountBy(productId, userId);
    } catch (Exception e) {
      log.log(WARNING, format("m=getDiscountBy, productId=%s, userId=%d, msg=%s",
        productId, userId, e.getMessage()), e);
      return null;
    }
  }
}
