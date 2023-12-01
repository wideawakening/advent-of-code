package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("foo", func() {
		It("bar", func() {
			Expect(1).To(Equal(1))
		})
	})
})

var _ = Describe("challenge", func() {
	XWhen("given sample", func() {
		It("resolve ...", func() {

		})
	})
	XWhen("given input file", func() {
		It("resolve ...", func() {
			Expect(Star1("input.txt")).To(Equal(0))
		})
	})
})
