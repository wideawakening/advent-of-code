package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {

		/*
					1-2 = (3,0),(7,1) 			= 6
			        1-3 = (3,0),(0,2) 			= 5
					1-4 = (3,0),(6,4) 			= 9
					1-5 = (3,0),(1,5) 			= 9
					1-6= (3,0),(9,6) 			= 1
					1-7

					5-9 = (1,5),(4,9)			= 9
					1-7 = (3,0),(7,8) 			= 15
					3-6 = (0,2),(9,6)			= 17
					8-9 = (0,9) - (4,9) 		= 5

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

				Expect(resolution).NotTo(Equal(9152334)) // too low
				Expect(resolution).NotTo(Equal(9102603)) // too low
				Expect(resolution).To(Equal(82000210))
				Expect(resolution).NotTo(Equal(9590997)) // too high
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
