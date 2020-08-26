package com.github.zeroberto.productsstore.usecase;

import com.github.zeroberto.productsstore.model.Discount;

public interface DiscountUseCase {

  Discount calculateDiscount(final String productId, final long userId);
}
