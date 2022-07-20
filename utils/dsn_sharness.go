//go:build sharness
// +build sharness

package utils

func dnsSharness(dsn string) string {
	return dsn + "_sharness"
}
