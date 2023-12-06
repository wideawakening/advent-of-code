package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	/*
		calibration value
		per line - first and last digit combined -  form a number = calibration
		sum all those numbers
	*/
	Context("Star1", func() {
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star1_sample.txt")).To(Equal(142))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star1("star1_input.txt")).To(Equal(56042))
			})
		})
	})

	/*
		calibration value
		per line - first and last digit combined -  form a number = calibration
		sum all those numbers
	*/
	Context("Star2", func() {
		When("given sample", func() {
			It("resolve ...", func() {
				Expect(TryNumberWord("eightwothree")).To(Equal(8))
				Expect(TryNumberWord("twothree")).To(Equal(2))
				Expect(TryNumberWord("three")).To(Equal(3))

				Expect(Star2("star2_sample.txt")).To(Equal(281))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star2("star1_input.txt")).To(Equal(55358))
			})
		})
	})
})
