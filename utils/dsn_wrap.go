//go:build wrap
// +build wrap

package utils

func warpDSN(dsn string) string {
	return dsn + "_warp"
}
