package com.github.zeroberto.productsstore.utils;

import com.github.zeroberto.productsstore.exceptions.FileReadingException;
import org.junit.jupiter.api.Test;

import java.io.File;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

class FileUtilsTest {

  @Test
  void readResourceFile() {
    final File file = FileUtils.readResourceFile("security/cert.pem");

    assertNotNull(file);
  }

  @Test
  void readResourceFileWhenFileNotExists() {
    assertThrows(FileReadingException.class, () -> FileUtils.readResourceFile("non_exists"));
  }
}
