package com.github.zeroberto.productsstore.datastore;

import com.github.zeroberto.productsstore.model.Discount;

public interface DiscountDataStore {

  Discount getDiscountBy(final String productId, final long userId);
}
