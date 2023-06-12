package s3_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ctx       = context.Background()
	err       error
	region    string
	accessKey string
	secretKey string
)

func TestS3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "S3 Suite")
}

var _ = BeforeSuite(func() {
	if v, ok := os.LookupEnv("AWS_REGION"); ok {
		region = v
	}
	if v, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); ok {
		accessKey = v
	}
	if v, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); ok {
		secretKey = v
	}
})
