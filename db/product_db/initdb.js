let res = [
  db.products.drop(),
  db.products.insertMany([
    { price_in_cents: 119, title: "Chocolate", description: "Everybody likes!", discount: { pct: 10.0, value_in_cents: 12 } },
    { price_in_cents: 1000, title: "Popcorn", description: "The favorite of moviegoers!", discount: { pct: 10.0, value_in_cents: 100 } },
    { price_in_cents: 500, title: "Orange juice", description: "Give me a dollar!", discount: { pct: 10.0, value_in_cents: 50 } },
  ])
]

printjson(res)
