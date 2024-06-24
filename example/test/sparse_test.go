package test

import (
	"github.com/svenwiltink/sparsecat"
	"github.com/svenwiltink/sparsecat/format"
	"io"
	"os"
	"testing"
)

func TestSparseCat(t *testing.T) {
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := sparsecat.NewEncoder(file)
	encoder.Format = format.RbdDiffv2
	encoder.MaxSectionSize = 16_000_000

	//file.Seek(10, io.SeekStart)
	_, err = io.Copy(os.Stdout, encoder)
	if err != nil {
		panic(err)
	}
}
