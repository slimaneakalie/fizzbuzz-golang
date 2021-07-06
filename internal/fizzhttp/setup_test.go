package fizzhttp

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFizzHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FizzHttp Suite")
}
