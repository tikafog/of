//go:build sharness
// +build sharness

package utils

func dsnSharness(dsn string) string {
	return dsn + "_sharness"
}
