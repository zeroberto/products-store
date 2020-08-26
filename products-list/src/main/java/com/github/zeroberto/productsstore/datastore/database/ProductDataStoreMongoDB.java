package com.github.zeroberto.productsstore.datastore.database;

import com.github.zeroberto.productsstore.datastore.ProductDataStore;
import com.github.zeroberto.productsstore.driver.ProductMongoClientBuilder;
import com.github.zeroberto.productsstore.model.Product;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import lombok.RequiredArgsConstructor;
import org.bson.Document;

import javax.annotation.Nonnull;
import java.util.LinkedList;
import java.util.List;

import static java.util.stream.Collectors.toList;

@RequiredArgsConstructor
public class ProductDataStoreMongoDB implements ProductDataStore {

  private final ProductMongoClientBuilder productMongoClientBuilder;
  private final String databaseName;

  @Override
  public List<Product> findAll() {
    MongoDatabase mongoDatabase = productMongoClientBuilder.build().getDatabase(databaseName);
    MongoCollection<Document> products = mongoDatabase.getCollection("products");
    final List<Document> documents = products.find().into(new LinkedList<>());
    return documents.stream().map(this::toProduct).collect(toList());
  }

  @Nonnull
  private Product toProduct(final Document document) {
    return Product.builder()
      .id(document.getObjectId("_id").toString())
      .priceInCents(document.getInteger("price_in_cents"))
      .title(document.getString("title"))
      .description(document.getString("description"))
      .build();
  }
}
