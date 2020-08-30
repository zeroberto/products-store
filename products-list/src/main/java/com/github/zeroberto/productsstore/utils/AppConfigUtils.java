package com.github.zeroberto.productsstore.utils;

import com.github.zeroberto.productsstore.config.AppConfig;
import lombok.NoArgsConstructor;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class AppConfigUtils {

  private static final String ENV_TYPE = "env";

  public static String getAuthUser(final AppConfig.AuthConfig authConfig) {
    if (ENV_TYPE.equals(authConfig.getType())) {
      return System.getenv(authConfig.getUser());
    }
    return authConfig.getUser();
  }

  public static String getAuthPass(final AppConfig.AuthConfig authConfig) {
    if (ENV_TYPE.equals(authConfig.getType())) {
      return System.getenv(authConfig.getPass());
    }
    return authConfig.getPass();
  }
}
