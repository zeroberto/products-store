package com.github.zeroberto.productsstore.utils;

import com.github.zeroberto.productsstore.exceptions.FileReadingException;
import lombok.NoArgsConstructor;

import java.io.File;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Paths;
import java.util.Optional;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class FileUtils {

  public static File readResourceFile(final String fileName) {
    try {
      final Optional<URL> url = Optional.ofNullable(ClassLoader.getSystemResource(fileName));

      return Paths
        .get(url.orElseThrow(FileReadingException::new).toURI())
        .toFile();
    } catch (URISyntaxException | NullPointerException e) {
      throw new FileReadingException(e);
    }
  }
}
