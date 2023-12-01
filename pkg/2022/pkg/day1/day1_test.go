package day1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("util testing", func() {
	When("testing function GetLowestValueIndex", func() {
		It("returns expected results", func() {
			index, value := GetLowestValueIndex([3]int{10, 6, 0})
			Expect(index).To(Equal(2))
			Expect(value).To(Equal(0))
		})
	})
})

var _ = Describe("given an elf calorie input file", func() {
	It("returns the number of calories with the elf with most calories for, star1", func() {
		ret := Star1("./input.txt")
		Expect(ret).To(Equal(67633))
	})

	It("returns the sum of top3 calories with the elf with most calories,  for star2", func() {
		ret := Star2("./input.txt")
		Expect(ret).To(Equal(199628))
	})

})
