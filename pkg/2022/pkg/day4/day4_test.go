package day4_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"2022-advent/pkg/day4"
)

var _ = Describe("utils", func() {
	When("given two ranges", func() {
		DescribeTable("gets if pair fully cover range", func(pair1 string, pair2 string, expectCoverage bool) {
			overange := day4.GetFullyCovers(pair1, pair2)
			Expect(overange).To(Equal(expectCoverage))
		},
			Entry("_", "2-4", "6-8", false),
			Entry("_", "2-3", "4-5", false),
			Entry("_", "5-7", "7-9", false),
			Entry("_", "2-8", "3-7", true),
			Entry("_", "6-6", "4-6", true),
			Entry("_", "2-6", "4-8", false),
		)

		DescribeTable("gets pair range overlaps ", func(pair1 string, pair2 string, expectCoverage bool) {
			covers := day4.GetCovers(pair1, pair2)
			Expect(covers).To(Equal(expectCoverage))
		},
			Entry("_", "2-4", "6-8", false),
			Entry("_", "2-3", "4-5", false),
			Entry("_", "5-7", "7-9", true),
			Entry("_", "2-8", "3-7", true),
			Entry("_", "6-6", "4-6", true),
			Entry("_", "2-6", "4-8", true),
			Entry("_", "47-96", "48-95", true),
			Entry("_", "40-40", "39-815", true),
			Entry("_", "24-91", "80-92", true),

			Entry("_", "9-43", "8-54", true),
			Entry("_", "6-51", "3-3", false),
			Entry("_", "9-9", "9-69", true),
			Entry("_", "29-79", "30-78", true),
			Entry("_", "38-70", "38-69", true),
			Entry("_", "2-99", "1-98", true),
			Entry("_", "2-3", "4-5", false),
		)
	})
})

var _ = Describe("Star1", func() {
	When("processed input file", func() {
		It("returns number of full overlaps", func() {
			Expect(day4.Star1("input.txt", true)).To(Equal(547))
		})
	})

	When("processed input file", func() {
		It("returns number of overlaps", func() {
			Expect(day4.Star1("input.txt", false)).To(Equal(843))
		})
	})
})
