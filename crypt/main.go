package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"io/ioutil"
	"path/filepath"

	"go-learn/base"
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
	path := base.Dir("crypt/..")
	adlerHash1, _ := getFileHash(filepath.Join(path, "crypt", "main.go"))
	adlerHash2, _ := getFileHash(filepath.Join(path, "hello", "main.go"))
	fmt.Println(adlerHash1, adlerHash2, adlerHash1 == adlerHash2)
	fmt.Println()

	/* SHA1 */

	fmt.Println("SHA1:")
	sha1Hash := sha1.New()
	sha1Hash.Write(msg)
	sha1Val := sha1Hash.Sum(nil)
	fmt.Println(sha1Val)
	fmt.Println()

	/* SHA256 */

	fmt.Println("SHA256:")
	sha256Hash := sha256.New()
	sha256Hash.Write(msg)
	sha256Val := sha256Hash.Sum(nil)
	fmt.Println(sha256Val)
}
