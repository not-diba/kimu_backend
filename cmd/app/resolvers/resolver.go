//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import "database/sql"

type Resolver struct {
	DB *sql.DB
}
