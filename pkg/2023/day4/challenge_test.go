package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {
		It(" parses one scratchcard content", func() {
			Expect(EvaluateScratchStar1("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")).To(Equal(float64(8))) // 4
			Expect(EvaluateScratchStar1("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19")).To(Equal(float64(2))) // 2
			Expect(EvaluateScratchStar1("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1")).To(Equal(float64(2))) // 2
			Expect(EvaluateScratchStar1("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83")).To(Equal(float64(1))) // 1
			Expect(EvaluateScratchStar1("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36")).To(Equal(float64(0))) // 0

		})
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt")).To(Equal(float64(13)))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star1("star_input.txt")).To(Equal(float64(22674)))
			})
		})
	})

	Context("Star2", func() {
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star2("star_sample.txt")).To(Equal(30))
			})
		})
		When("given input", func() {
			It("resolves", func() {

				Expect(Star2("star_input.txt")).To(Not(Equal(5747444)))
				Expect(Star2("star_input.txt")).To(Not(Equal(5747456)))
				Expect(Star2("star_input.txt")).To(Equal(0))
			})
		})
	})
})
