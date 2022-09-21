package errors

func ModuleError(i uint32, err error) error {
	return WrapIndex(err, MakeErrIndex(i, 0))
}
