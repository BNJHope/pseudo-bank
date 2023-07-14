# pseudo-bank

A simulation digital wallet app, using a Go REST API with a PostgreSQL based database storage.

## Run

The whole app ecosystem can be brought up through `docker compose`. To do this, run:

```shell
docker compose --profile app up
```

It can then be queried from your host machine:

```shell
➜  pseudo-bank git:(main) ✗ curl -s localhost:3333/transaction | jq "."
[
  {
    "id": 1,
    "amount": 20.5,
    "from": "d2e19190-59c8-4a43-8bb7-a729ea2b5173",
    "to": "1a8580b6-fb6c-4f3a-8254-3c19e638f385"
  }
]
```
