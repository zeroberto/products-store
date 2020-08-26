package com.github.zeroberto.productsstore.datastore;

import com.github.zeroberto.productsstore.model.Product;

import java.util.List;

public interface ProductDataStore {

  List<Product> findAll();
}
