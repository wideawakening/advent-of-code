package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {
		XWhen("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star1_sample.txt")).To(Equal(0))
			})
		})
		XWhen("given input", func() {
			It("resolves", func() {
				Expect(Star1("star1_input.txt")).To(Equal(0))
			})
		})
	})
	Context("Star2", func() {
		XWhen("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star2_sample.txt")).To(Equal(0))
			})
		})
		XWhen("given input", func() {
			It("resolves", func() {
				Expect(Star1("star2_input.txt")).To(Equal(0))
			})
		})
	})
})
