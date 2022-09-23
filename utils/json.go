package utils

func Must(data []byte, err error) []byte {
	if err != nil {
		return nil
	}
	return data
}
