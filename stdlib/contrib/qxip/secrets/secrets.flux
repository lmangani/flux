// Package secrets functions for working with sensitive secrets managed via ENV variables.
//
// ## Metadata
// introduced: NEXT
// tags: secrets,security
//
package secrets


// get retrieves a secret from the InfluxDB secret store.
//
// ## Parameters
// - key: Secret key to retrieve.
//
// ## Examples
//
// ### Retrive a key from ENV variables
// ```no_run
// import "contrib/qxip/secrets"
//
// secrets.get(key: "KEY_NAME")
// ```
//
// ### Populate sensitive credentials with ENV variables//
// ```no_run
// import "sql"
// import "contrib/qxip/secrets"
//
// username = secrets.get(key: "POSTGRES_USERNAME")
// password = secrets.get(key: "POSTGRES_PASSWORD")
//
// sql.from(
//     driverName: "postgres",
//     dataSourceName: "postgresql://${username}:${password}@localhost",
//     query: "SELECT * FROM example-table",
// )
// ```
//
builtin get : (key: string) => string
