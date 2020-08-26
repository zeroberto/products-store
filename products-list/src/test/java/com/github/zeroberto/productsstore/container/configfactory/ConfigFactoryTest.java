package com.github.zeroberto.productsstore.container.configfactory;

import com.github.zeroberto.productsstore.config.AppConfig;
import org.junit.jupiter.api.Test;
import org.yaml.snakeyaml.error.YAMLException;

import static com.github.zeroberto.productsstore.container.configfactory.AppConfigFactory.makeAppConfig;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

class ConfigFactoryTest {

  @Test
  void givenMakeConfigFile_thenNotNull() {
    final AppConfig got = makeAppConfig("config.yml");

    assertNotNull(got);
  }

  @Test
  void givenReadConfigFile_whenFileNotExists_thenThrows() {
    assertThrows(YAMLException.class, () -> makeAppConfig("conf.yml"));
  }
}
