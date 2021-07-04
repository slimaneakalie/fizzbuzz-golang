package api

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api package - Router file", func() {
	Context("Router testing", func() {
		It("returns `true`", func() {
			Expect(strings.Contains("Ginkgo is awesome", "is")).To(BeTrue())
		})
	})
})
