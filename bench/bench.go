package bench

import (
	"os"
)

// Измерение длины файла
func FileLen(f string, bufsize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var count int
	buf := make([]byte, bufsize)
	for {
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}

	return count, nil
}
