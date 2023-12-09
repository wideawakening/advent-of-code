package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	It("calculates the next sequence up till all zeroes", func() {
		Expect(GenerateNextSequenceTree([][]int{{0, 3, 6, 9, 12, 15}})).To(Equal([][]int{
			{0, 3, 6, 9, 12, 15},
			{3, 3, 3, 3, 3},
			{0, 0, 0, 0}}))

		// TODO add more tests here if required
	})

	It("interpolates next value on first sequence", func() {
		Expect(InterpolateNextValues(GenerateNextSequenceTree([][]int{{0, 3, 6, 9, 12, 15}}))).To(Equal([][]int{
			{0, 3, 6, 9, 12, 15, 18},
			{3, 3, 3, 3, 3, 3},
			{0, 0, 0, 0, 0}}))
	})

	Context("Star1", func() {
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star("star_sample.txt", false)).To(Equal(114))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star("star_input.txt", false)).To(Equal(1898776583))
			})
		})
	})
	Context("Star2", func() {
		It("interpolates next value on first sequence", func() {
			Expect(InterpolatePreviousValues(GenerateNextSequenceTree([][]int{{10, 13, 16, 21, 30, 45}}))).To(Equal([][]int{
				{5, 10, 13, 16, 21, 30, 45},
				{5, 3, 3, 5, 9, 15},
				{-2, 0, 2, 4, 6},
				{2, 2, 2, 2},
				{0, 0, 0}}))
		})

		// optimized
		//It("interpolates next value on first sequence", func() {
		//	data := []int{10, 13, 16, 21, 30, 45}
		//	slices.Reverse(data)
		//	Expect(InterpolateNextValues(GenerateNextSequenceTree([][]int{data}))).To(Equal([][]int{}))
		//})

		When("given sample", func() {
			It("resolves", func() {
				Expect(Star("star_sample.txt", true)).To(Equal(2))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star("star_input.txt", true)).To(Equal(1100))
			})
		})
	})
})
