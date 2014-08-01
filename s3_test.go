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

func TestUpload(t *testing.T) {
	s3 := Init("", "")
	bucketResult, err := s3.ListBucket("sps-photos")
	if err != nil {
		t.Error("upload error")
	}

	for k, v := range bucketResult.Contents {
		fmt.Printf("%s %s\n", k, v.Key)
	}
}
