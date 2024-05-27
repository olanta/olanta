package utils

import (
	"io/ioutil"
	"os"
)

// ReadFile reads the content of a file and returns it as a byte slice.
func ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

// WriteFile writes data to a file.
func WriteFile(filepath string, data []byte) error {
	return ioutil.WriteFile(filepath, data, os.ModePerm)
}
