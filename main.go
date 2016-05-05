package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

func GetFileMd5(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}
	return hash.Sum(result), nil
}

// GetFileInode accepts absolute or relative filepath string.
// Returns a uint64 of the file inode number from the Stat_t struct.
func GetFileInode(filepath string) (uint64, error) {
	var stat syscall.Stat_t
	if err := syscall.Stat(filepath, &stat); err != nil {
		return 0, err
	}
	return stat.Ino, nil
}

func main() {
	const filename string = "test.txt"
	// slice of filename strings (could create struct soon for this)
	//var files = []string{"test1.txt", "test2.txt", "test3.txt"}
	for {
		if b, err := GetFileMd5(filename); err != nil {
			fmt.Printf("Checksum Err: %v\n", err)
		} else {
			fmt.Printf("%s md5 checksum is: %x\n", filename, b)
		}
		if ino, err := GetFileInode(filename); err != nil {
			fmt.Println("Inode Err: %v\n", err)
		} else {
			fmt.Printf("%s inode is: %d\n", filename, ino)
		}
		time.Sleep(time.Second * 15)
	}
}
