# Implements Hexagonal Architecture using TDD

- [x] Application
- [x] Adapters
    - [x] sqlite database
    - [x] cli using cobra
    - [x] webserver

# Install dependencies

```bash
$ go get ./...
```

# Run

```bash
$ docker compose up -d
$ docker exec -it appproduct bash
```

To run tests:

```bash
$ go test ./...
```

To start webserver(in port 9000):

Methods:

- `GET product/{id}` to get a product
- `POST product` to create a new product
- `GET product/{id}/enable` to enable a product
- `GET product/{id}/disable` to disable a product

```bash
$ go run main.go http
```

To run the cli:

flags:

- `--action` available: enable, disable, create or get
    - Example: `--action=enable`
- `--id`
    - Example: `--action=get --id=cce6c196-05a1-4b5c-9c3c-02dde12acf66`
- `--price`
    - Example: `--action=create --product="Product Name" --price=44.4`
- `--product`
    - Example: `--action=create --product="Product Name" --price=44.4`

```bash
$ go run main.go cli --help
```

or

```bash
$ go run main.go cli --action=create --product="Product Name" --price=44.4
```
