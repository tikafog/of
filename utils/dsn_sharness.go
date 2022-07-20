//go:build sharness
// +build sharness

package utils

func sharnessDSN(dsn string) string {
	return dsn + "_sharness"
}
