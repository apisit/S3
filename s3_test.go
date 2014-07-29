package s3

import "testing"

func TestUpload(t *testing.T) {
	var client Client
	client.Auth = &Auth{"", "", ""}
	client.upload("success_page.png", "sps-photos", "success_page.png")
}
