let res = [
  db.products.drop(),
  db.products.insertMany([
    { _id: ObjectId("5f4962fb3ff6e3f16dca574e"), price_in_cents: NumberInt(120), title: "Product 1", description: "Test" },
    { _id: ObjectId("5f4962fb3ff6e3f16dca574f"), price_in_cents: NumberInt(1400), title: "Product 2", description: "Test" },
    { _id: ObjectId("5f4962fb3ff6e3f16dca5750"), price_in_cents: NumberInt(500), title: "Product 3", description: "Test" },
  ])
]

printjson(res)
