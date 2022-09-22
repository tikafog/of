package dbc

var err = merr.RegisterModule("DBC")

func Error(s string) error {
	return err.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return err.Errorf(format, args...)
}
