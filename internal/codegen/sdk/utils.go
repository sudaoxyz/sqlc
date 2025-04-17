package sdk

import (
	"fmt"
	"strings"
	"unicode"
)

func TrimPrefix(s, prefix string) string {
	return strings.ToLower(strings.TrimPrefix(s, prefix))
}

func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

func NoNullType(t string) string {
	switch t {
	case "sql.NullString":
		return "string"
	case "sql.NullInt64":
		return "int64"
	case "sql.NullTime":
		return "time.Time"
	default:
		return t
	}
}

func ToNullType(t, value string) string {
	switch t {
	case "sql.NullString":
		return fmt.Sprintf("sql.NullString{String:%s,Valid: %s != \"\"}", value)
	case "sql.NullInt64":
		return fmt.Sprintf("sql.NullInt64{Int64:%s,Valid: %s != 0}", value)
	case "sql.NullTime":
		return fmt.Sprintf("sql.NullTime{Time:%s,Valid: %s.IsZero()}", value)
	default:
		return value
	}
}

func LowerTitle(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func Title(s string) string {
	return strings.Title(s)
}

// Go string literals cannot contain backtick. If a string contains
// a backtick, replace it the following way:
//
// input:
//
//	SELECT `group` FROM foo
//
// output:
//
//	SELECT ` + "`" + `group` + "`" + ` FROM foo
//
// # The escaped string must be rendered inside an existing string literal
//
// A string cannot be escaped twice
func EscapeBacktick(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}

func DoubleSlashComment(s string) string {
	return "// " + strings.ReplaceAll(s, "\n", "\n// ")
}
