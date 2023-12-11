package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {

		/*
				5-9 = (1,7),(6,9)					= 9 OK
				1-7 = (0,3),(8,7) 				= 15 OK
				3-6 = (2,0),(1,7)					= 17 OK
				8-9 = (9,0) - (9,4) = 5 OK

			looking good but not matching
		*/

		FWhen("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt")).To(Equal(374))
			})
		})

		When("given input", func() {
			It("resolves", func() {
				resolution := Star1("star_input.txt")
				Expect(resolution).NotTo(Equal(9590997)) // too high
				Expect(resolution).To(Equal(0))
			})
		})
	})
	Context("Star2", func() {
		XWhen("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt")).To(Equal(0))
			})
		})
		XWhen("given input", func() {
			It("resolves", func() {
				Expect(Star1("star_input.txt")).To(Equal(0))
			})
		})
	})
})
