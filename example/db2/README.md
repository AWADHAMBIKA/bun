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

The example uses `db2dialect.New()`, which detects the target platform from the
connection. For an explicit platform, use `db2dialect.NewLUW()` for DB2 LUW,
`db2dialect.NewZOS()` for DB2 for z/OS, or `db2dialect.NewIBMi()` for DB2 for
IBM i.

To disable query logging:

```shell
BUNDEBUG=0 go run .
```

See [docs](https://bun.uptrace.dev/) for details.
