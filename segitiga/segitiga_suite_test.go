package segitiga_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSegitiga(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Segitiga Suite")
}
