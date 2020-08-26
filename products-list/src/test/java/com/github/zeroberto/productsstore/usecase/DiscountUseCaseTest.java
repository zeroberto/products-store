package com.github.zeroberto.productsstore.usecase;

import com.github.zeroberto.productsstore.datastore.DiscountDataStore;
import com.github.zeroberto.productsstore.exceptions.DataStoreNetworkException;
import com.github.zeroberto.productsstore.model.Discount;
import com.github.zeroberto.productsstore.usecase.impl.DiscountUseCaseImpl;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class DiscountUseCaseTest {

  private DiscountDataStore discountDataStore;
  private DiscountUseCase discountUseCase;

  @BeforeEach
  void setUp() {
    discountDataStore = mock(DiscountDataStore.class);

    discountUseCase = new DiscountUseCaseImpl(discountDataStore);
  }

  @Test
  void calculateDiscount() {
    final String productId = "test";
    final long userId = 1L;
    final Discount expected = Discount.builder()
      .pct(10L)
      .valueInCents(10)
      .build();

    when(discountDataStore.getDiscountBy(anyString(), anyLong())).thenReturn(expected);

    final Discount got = discountUseCase.calculateDiscount(productId, userId);

    assertNotNull(got);
    assertEquals(expected, got);
  }

  @Test
  void calculateDiscountWhenDataStoreNetworkExceptionThenReturnNull() {
    final String productId = "test";
    final long userId = 1L;

    when(discountDataStore.getDiscountBy(anyString(), anyLong())).thenThrow(new DataStoreNetworkException(null));

    final Discount discount = discountUseCase.calculateDiscount(productId, userId);

    assertNull(discount);
  }
}
