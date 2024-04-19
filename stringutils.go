package stringutils

import (
	"database/sql"
	"strings"
)

// NullIfEmpty renvoie une valeur sql.NullString qui est null si la chaîne d'entrée est null ou ne contient que des espaces blancs,
// ou la chaîne d'entrée trimée si elle contient une valeur.
func NullIfEmpty(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	t := strings.TrimSpace(*s)

	if len(t) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{String: t, Valid: true}
}
