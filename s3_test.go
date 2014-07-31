package s3

import (
	"fmt"
	"testing"
)

func TestUpload(t *testing.T) {
	s3 := Init("", "")
	file, err := s3.Upload("test.jpg", "sps-photos", "test.jpg")
	if err != nil {
		t.Error("upload error")
	}
	fmt.Printf("%s", file)
}
