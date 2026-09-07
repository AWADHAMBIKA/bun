# DB2 example

This example demonstrates using Bun with IBM DB2 via the
[github.com/ibmdb/go_ibm_db](https://github.com/ibmdb/go_ibm_db) driver and the
`db2dialect` package.

> **Note:** `go_ibm_db` uses cgo and links against IBM's proprietary DB2 CLI/ODBC
> client driver (`libdb2`). Building or running this example requires that
> driver to be installed locally, so it is excluded from this repo's CI and
> from `make test` / `make run-examples`.

Update the `dsn` connection string in `main.go` to point at your DB2 instance, then run:

```shell
go run .
```

To disable query logging:

```shell
BUNDEBUG=0 go run .
```

See [docs](https://bun.uptrace.dev/) for details.
