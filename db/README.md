# products-list

This project is responsible for defining the data infrastructure of the services domain, surveying the databases used by the apps.

## Products Domain

The domain-oriented microservice architecture requires access to user and product data.

Although the application domain is related to products, it is necessary to obtain user information to, for example, calculate personalized discounts.

## Products Database

Using MongoDB, this NoSQL database is responsible for storing product data.

MongoDB was chosen because it is one of the most permormmatic, flexible and scalable document database technologies (NoSQL) today. In addition, more properties can be added later, as a well-defined product layout is not required, allowing for dynamic documents.

Below you can see the initial scheme for products.

```json
{
    "id": "string",
    "price_in_cents": "int",
    "title": "string",
    "description": "string",
    "discount": {
        "pct": "float",
        "value_in_cents": "int"
    }
}
```

## Users Database

Using PostgreSQL, this database is responsible for storing user data.

A structured database (SQL) serves this context in a simple way. With a well-defined scheme, the structure of this entity's data will hardly change. Although the number of users can surpass the number of products, a relational database manages to meet the needs analyzed.

Below you can see the schematic representation of the user.

| Column         | Type       | Description |
| -------------- |:---------- | :---------- |
| id             | BIGINT     | PRIMARY KEY |
| first_name     | VARCHAR    | Size 50     |
| last_name      | VARCHAR    | Size 50     |
| date_of_birth  | DATE       |             |
| created_at     | TIMESTAMP  |             |
| updated_at     | TIMESTAMP  |             |
| deactivated_at | TIMESTAMP  |             |

## Up Servers

### Prerequisites

To make the servers available, you must have:

* [Docker](https://docs.docker.com/get-docker/)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Execution

To up the servers independently of the applications, just execute the following command:

```bash
docker-compose up
```

Or

```bash
docker-compose up -d
```

to run the container in the background.

### Accessing the Database Administrator

#### Adminer

To connect to the Adminer, access the address below, after the container initialization:

> http://localhost:9000/

Then enter the desired base information, found in the [docker-compose.yml](docker-compose.yml) file.

#### Mongo Express

To connect to the Mongo Express, access the address below, after the container initialization:

> http://localhost:9001/

### Reset Databases

The following command is used to return the databases to their initial state. It will completely clean the folders of the base volumes. Then restart the container.

```bash
sudo rm -r /containers/
```

## License

[MIT](LICENSE) License
