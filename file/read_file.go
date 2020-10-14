package file

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

// Copy 读取文件
func Copy(path string) (fileMd5 string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return fileMd5, err
	}
	defer f.Close()

	md5hash := sha1.New()

	if _, err := io.Copy(md5hash, f); err != nil {
		return fileMd5, err
	}
	fileMd5 = hex.EncodeToString(md5hash.Sum(nil))
	return fileMd5, nil
}

// ReadAll 读取文件， 文件不能过大
func ReadAll(path string) (fileMd5 string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return fileMd5, err
	}
	defer f.Close()
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return fileMd5, err
	}
	hash := sha1.New()
	hash.Write(body)
	fileMd5 = hex.EncodeToString(hash.Sum(nil))
	return fileMd5, err
}

// ReadBuf 使用bufio.NewReader
func ReadBuf(path string) (fileMd5 string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return fileMd5, err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	reader := bufio.NewReader(f)
	md5hash := sha1.New()

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				goto stop
			}
			return fileMd5, err
		}
		md5hash.Write(buf[:n])
	}
stop:
	fileMd5 = hex.EncodeToString(md5hash.Sum(nil))
	return fileMd5, nil
}
