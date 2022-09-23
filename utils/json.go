package utils

func Must(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
}
