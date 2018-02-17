package cache

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"testing"
	"path/filepath"
)

func createRandomFile(dir string, size int64) (string, error) {
	data, filename := randomDataAndHash(size)
	filepath := dir + "/" + filename

	return filename, ioutil.WriteFile(filepath, data, 0744)
}

func randomDataAndHash(size int64) ([]byte, string) {
	data := make([]byte, size)
	rand.Read(data)
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])
	return data, hashStr
}

func createTmpCacheDirs(t *testing.T) string {
	path, err := ioutil.TempDir("", "ensurespacer")
	if err != nil {
		t.Error("Couldn't create tmp dir", err)
	}
	ensureDirExists(filepath.Join(path, "ac"))
	ensureDirExists(filepath.Join(path, "cas"))

	return path
}
