package stringListBuilder

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFizzbuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringListBuilder Suite")
}
