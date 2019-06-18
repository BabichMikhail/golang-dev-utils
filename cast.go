package utils

import (
	"database/sql"
)

func SqlNullInt64ToInt(value sql.NullInt64) int {
	CheckTrue(value.Valid, "Value is invalid")
	return int(value.Int64)
}
