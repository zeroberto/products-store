let res = [
  db.products.drop(),
  db.products.insertMany([
    { price_in_cents: 119, title: "Chocolate", description: "Everybody likes!" },
    { price_in_cents: 1000, title: "Popcorn", description: "The favorite of moviegoers!" },
    { price_in_cents: 500, title: "Orange juice", description: "Give me a dollar!" },
  ])
]

printjson(res)
