package main

import (
	"crypto/sha1"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"io/ioutil"
)

// Хэш содержимого файла
func getFileHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	h := adler32.New()
	h.Write(bs)

	return h.Sum32(), nil
}

func main() {
	fmt.Println(" \n[ КРИПТОГРАФИЯ ]\n ")

	msg := []byte("Secret message")
	fmt.Println("Исходный текст:", string(msg))

	fmt.Println()

	/* CRC32 */

	fmt.Println("CRC32:")
	crcHash := crc32.NewIEEE()
	crcHash.Write(msg)
	crcVal := crcHash.Sum32()
	fmt.Println(crcVal)

	fmt.Println()

	/* Adler32 */

	fmt.Println("Adler32:")
	adlerHash1, _ := getFileHash("hello.go")
	adlerHash2, _ := getFileHash("http.go")
	fmt.Println(adlerHash1, adlerHash2, adlerHash1 == adlerHash2)

	fmt.Println()

	/* SHA1 */

	fmt.Println("SHA1:")
	shaHash := sha1.New()
	shaHash.Write(msg)
	shaVal := shaHash.Sum([]byte{})
	fmt.Println(shaVal)
}
