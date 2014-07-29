package s3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func IsValidBucket(bucket string) bool {
	l := len(bucket)
	if l < 3 || l > 63 {
		return false
	}

	valid := false
	prev := byte('.')
	for i := 0; i < len(bucket); i++ {
		c := bucket[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z':
			valid = true
		case '0' <= c && c <= '9':
			// Is allowed, but bucketname can't be just numbers.
			// Therefore, don't set valid to true
		case c == '-':
			if prev == '.' {
				return false
			}
		case c == '.':
			if prev == '.' || prev == '-' {
				return false
			}
		}
		prev = c
	}

	if prev == '-' || prev == '.' {
		return false
	}
	return valid
}

func Init(accesskey string, secretKey string) *Client {
	return &Client{&Auth{accesskey, secretKey, ""}}
}

type Client struct {
	*Auth
}

// bucketURL returns the URL prefix of the bucket, with trailing slash
func (c *Client) bucketURL(bucket string) string {
	if IsValidBucket(bucket) && !strings.Contains(bucket, ".") {
		return fmt.Sprintf("https://%s.%s/", bucket, c.hostname())
	}
	return fmt.Sprintf("https://%s/%s/", c.hostname(), bucket)
}

func (c *Client) keyURL(bucket, key string) string {
	return c.bucketURL(bucket) + key
}

func (c *Client) Upload(key, bucket string, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	url := c.keyURL(bucket, key)
	req, _ := http.NewRequest("PUT", url, nil)
	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	req.ContentLength = int64(len(data))
	body := bytes.NewBuffer(data)
	req.Body = ioutil.NopCloser(body)
	c.Auth.SignRequest(req)
	httpClient := &http.Client{}
	res, _ := httpClient.Do(req)
	defer res.Body.Close()
	_, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return readErr
	}
	return nil
}
