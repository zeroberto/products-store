package com.github.zeroberto.productsstore.container.configfactory;

import com.github.zeroberto.productsstore.config.AppConfig;
import lombok.NoArgsConstructor;
import org.yaml.snakeyaml.Yaml;
import org.yaml.snakeyaml.constructor.Constructor;
import org.yaml.snakeyaml.representer.Representer;

import javax.annotation.Nonnull;
import java.io.InputStream;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class AppConfigFactory {

  @Nonnull
  public static AppConfig makeAppConfig(final String filename) {
    InputStream inputStream = AppConfig.class
      .getClassLoader()
      .getResourceAsStream(filename);
    return getYaml().load(inputStream);
  }

  private static Yaml getYaml() {
    Representer representer = new Representer();
    representer.getPropertyUtils().setSkipMissingProperties(true);
    return new Yaml(new Constructor(AppConfig.class), representer);
  }
}
