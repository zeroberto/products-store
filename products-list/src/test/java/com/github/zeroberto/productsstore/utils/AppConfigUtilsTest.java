package com.github.zeroberto.productsstore.utils;

import com.github.zeroberto.productsstore.config.AppConfig;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.junit.jupiter.MockitoExtension;

import static com.github.zeroberto.productsstore.utils.AppConfigUtils.getAuthPass;
import static com.github.zeroberto.productsstore.utils.AppConfigUtils.getAuthUser;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class AppConfigUtilsTest {

  private AppConfig.AuthConfig authConfig;

  @BeforeEach
  void setUp() {
    authConfig = mock(AppConfig.AuthConfig.class);
  }

  @Test
  void testGetAuthUserWhenAuthTypeEqualToPlain() {
    final String expected = "usr";

    when(authConfig.getType()).thenReturn("plain");
    when(authConfig.getUser()).thenReturn("usr");

    final String got = getAuthUser(authConfig);

    assertEquals(expected, got);
  }

  @Test
  void testGetAuthPassWhenAuthTypeEqualToPlain() {
    final String expected = "pass";

    when(authConfig.getType()).thenReturn("plain");
    when(authConfig.getPass()).thenReturn("pass");

    final String got = getAuthPass(authConfig);

    assertEquals(expected, got);
  }
}
