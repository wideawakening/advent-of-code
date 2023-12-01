package day8

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {

	treeHeights := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0}}

	When("given treeHeights and a tree position", func() {
		DescribeTable("returns whether that tree is visible", func(treeHeights [][]int, posX int, posY int, isVisible bool) {

			Expect(CheckVisibleTree(treeHeights, posX, posY)).To(Equal(isVisible))
		},
			Entry("_", treeHeights, 1, 1, true),
			Entry("_", treeHeights, 1, 2, true),
			Entry("_", treeHeights, 1, 3, false),

			Entry("_", treeHeights, 2, 1, true),
			Entry("_", treeHeights, 2, 2, false),
			Entry("_", treeHeights, 2, 3, true),

			Entry("_", treeHeights, 3, 1, false),
			Entry("_", treeHeights, 3, 2, true),
			Entry("_", treeHeights, 3, 3, false),

			Entry("_", treeHeights, 0, 0, true),
			Entry("_", treeHeights, 0, 1, true),
			Entry("_", treeHeights, 0, 2, true),
			Entry("_", treeHeights, 0, 3, true),
			Entry("_", treeHeights, 0, 4, true),

			Entry("_", treeHeights, 0, 0, true),
			Entry("_", treeHeights, 1, 0, true),
			Entry("_", treeHeights, 2, 0, true),
			Entry("_", treeHeights, 3, 0, true),
			Entry("_", treeHeights, 4, 0, true),

			Entry("_", treeHeights, 4, 0, true),
			Entry("_", treeHeights, 4, 1, true),
			Entry("_", treeHeights, 4, 2, true),
			Entry("_", treeHeights, 4, 3, true),
			Entry("_", treeHeights, 4, 4, true),

			Entry("_", treeHeights, 0, 4, true),
			Entry("_", treeHeights, 1, 4, true),
			Entry("_", treeHeights, 2, 4, true),
			Entry("_", treeHeights, 3, 4, true),
			Entry("_", treeHeights, 4, 4, true),
		)
	})

	When("given treeHeights and a tree position", func() {
		DescribeTable("returns the scenic score", func(treeHeights [][]int, posX int, posY int, scenicScore int) {
			Expect(DetermineScenicScore(treeHeights, posX, posY)).To(Equal(scenicScore))
		},

			Entry("_", treeHeights, 2, 1, 4),
			Entry("_", treeHeights, 2, 3, 8),
		)
	})

	When("given sample", func() {
		It("resolve how many trees are visible from the outside grid", func() {

			Expect(GetNumberVisibleTress(treeHeights)).To(Equal(21))
		})
	})
	When("processed input file", func() {
		It("returns star1 puzzle value", func() {
			Expect(EvaluateStar("input.txt", true)).To(Equal(1676))
			Expect(EvaluateStar("input.txt", false)).To(Equal(313200))
		})
	})
})
