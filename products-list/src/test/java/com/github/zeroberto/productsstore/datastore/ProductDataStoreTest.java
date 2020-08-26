package com.github.zeroberto.productsstore.datastore;

import com.github.zeroberto.productsstore.datastore.database.ProductDataStoreMongoDB;
import com.github.zeroberto.productsstore.driver.ProductMongoClientBuilder;
import com.github.zeroberto.productsstore.model.Product;
import com.mongodb.MongoClient;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import org.bson.Document;
import org.bson.types.ObjectId;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import javax.annotation.Nonnull;
import java.util.Collection;
import java.util.List;
import java.util.Map;

import static java.util.Collections.singletonList;
import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

class ProductDataStoreTest {

  private MongoCollection<Document> mongoCollection;
  private ProductDataStore productDataStore;

  @BeforeEach
  void setUp() {
    final String databaseName = "test";
    final var mongoClient = mock(MongoClient.class);
    final var mongoDatabase = mock(MongoDatabase.class);
    final var productMongoClientBuilder = mock(ProductMongoClientBuilder.class);

    when(productMongoClientBuilder.build()).thenReturn(mongoClient);
    when(mongoClient.getDatabase(anyString())).thenReturn(mongoDatabase);

    mongoCollection = mock(TestMongoCollection.class);
    when(mongoDatabase.getCollection(anyString())).thenReturn(mongoCollection);

    productDataStore = new ProductDataStoreMongoDB(productMongoClientBuilder, databaseName);
  }

  @Test
  void findAll() {
    final List<Document> documents = singletonList(
      new Document(Map.of(
        "_id", new ObjectId("5e0c0b30bbb687466f3f433f"),
        "price_in_cents", 100,
        "title", "test",
        "description", "test"
      )));

    final FindIterable<Document> iterable = mock(TestFindIterable.class);
    when(mongoCollection.find()).thenReturn(iterable);
    when(iterable.into(any())).thenReturn(documents);

    final List<Product> products = productDataStore.findAll();

    assert !products.isEmpty();

    final Product product = products.get(0);

    assertAll("findAll",
      () -> assertEquals("5e0c0b30bbb687466f3f433f", product.getId()),
      () -> assertEquals(100, product.getPriceInCents()),
      () -> assertEquals("test", product.getTitle()),
      () -> assertEquals("test", product.getDescription()));
  }

  private interface TestMongoCollection extends MongoCollection<Document> {
  }

  private interface TestFindIterable extends FindIterable<Document> {

    @Override
    @Nonnull
    default <A extends Collection<? super Document>> A into(@Nonnull A target) {
      return target;
    }
  }
}
