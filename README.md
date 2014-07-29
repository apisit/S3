go get github.com/apisit/s3


import (
"github.com/apisit/s3"
)

//using it
	accessKey := "key"
	secretKey := "key"
	keyFilename := "something.jpg"
	filename := "something.jpg"
	bucketName := "sps-photos"
	amazonS3 := s3.Init(accessKey, secretKey)
	amazonS3.Upload(keyFilename, bucketName, filename)