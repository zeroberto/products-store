package com.github.zeroberto.productsstore.helper;

import com.mongodb.MongoClient;
import com.mongodb.MongoClientOptions;
import com.mongodb.MongoCredential;
import com.mongodb.MongoException;
import com.mongodb.ServerAddress;
import lombok.NoArgsConstructor;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.ServerSocket;
import java.net.URISyntaxException;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.Map;

import static java.util.Objects.requireNonNull;
import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class InfraHelper {

  private static final int CONN_TIMEOUT = 10;

  public static void upInfra(final String dockerComposeFileName) {
    final String command = String.format("docker-compose -f %s up -d", dockerComposeFileName);
    executeCommand(command);
  }

  public static void upInfra(final String dockerComposeFileName, final String serviceName) {
    final String command = String.format("docker-compose -f %s up -d %s", dockerComposeFileName, serviceName);
    executeCommand(command);
  }

  public static void upInfra(final String dockerComposeFileName, final Map<String, String> envs) {
    final StringBuilder sb = new StringBuilder();
    envs.forEach((k, v) -> sb.append(k).append("=").append(v).append(" "));
    executeCommand(sb.toString() + "docker-compose -f " + dockerComposeFileName + " up -d");
  }

  public static void downInfra(final String dockerComposeFileName) {
    executeCommand("docker-compose -f " + dockerComposeFileName + " down");
  }

  public static boolean checkDBConnection(final String dsn, final String user, final String pass) {
    try (final Connection conn = DriverManager.getConnection(dsn, user, pass)) {
      return conn.isValid(CONN_TIMEOUT);
    } catch (SQLException throwables) {
      return false;
    }
  }

  public static boolean checkMongoDBConnection(
    final String host, final int port, final String user,
    final String pass, final String repo
  ) {
    try (final var mongoClient = new MongoClient(
      new ServerAddress(host, port),
      MongoCredential.createCredential(user, repo, pass.toCharArray()),
      MongoClientOptions.builder().build())
    ) {
      mongoClient.listDatabaseNames();
      return true;
    } catch (MongoException e) {
      return false;
    }
  }

  public static boolean checkPortIsAvailable(final int port) {
    try (final var ignored = new ServerSocket(port)) {
      return false;
    } catch (IOException e) {
      return true;
    }
  }

  public static String executeCommand(final String command) {
    final ProcessBuilder processBuilder = new ProcessBuilder();

    processBuilder.command("bash", "-c", command);

    try {
      processBuilder.directory(new File(requireNonNull(
        InfraHelper.class.getClassLoader().getResource("docker")).toURI()));

      final Process process = processBuilder.start();
      final var output = new StringBuilder();
      final var reader = new BufferedReader(
        new InputStreamReader(process.getInputStream()));

      String line;
      while ((line = reader.readLine()) != null) {
        output.append(line).append("\n");
      }

      int exitVal = process.waitFor();
      if (exitVal == 0) {
        return output.toString();
      } else {
        throw new InfraUncheckedException(
          String.format("Execution failure. Exit code %d, Error %s", exitVal, output.toString()));
      }
    } catch (InterruptedException e) {
      Thread.currentThread().interrupt();
      throw new InfraUncheckedException(e);
    } catch (IOException | URISyntaxException e) {
      throw new InfraUncheckedException(e);
    }
  }
}
