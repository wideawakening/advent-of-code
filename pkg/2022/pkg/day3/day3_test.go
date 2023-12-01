package day3

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day3 - Star1 - Utils", func() {
	When("spliting each ruckack each rucksack", func() {
		DescribeTable("test rucksack compartment", func(rucksackContent string, compartmentA string, compartmentB string) {
			retCompA, retCompB := GetSplitRuckSac(rucksackContent)
			Expect(retCompA).To(Equal(compartmentA))
			Expect(retCompB).To(Equal(compartmentB))
		},
			Entry("_", "vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp"),
			Entry("_", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"),
			Entry("_", "PmmdzqPrVvPwwTWBwg", "PmmdzqPrV", "vPwwTWBwg"),
		)
	})

	When("searching for matching character", func() {
		DescribeTable("returns matching character", func(text1 string, text2 string, expectedCharacter string) {
			matching := GetMatchingCharacter(text1, text2, "")
			Expect(matching).To(Equal(expectedCharacter))
		},
			Entry("_", "vJrwpWtwJgWr", "hcsFMMfFFhFp", "p"),
			Entry("_", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL", "L"),
			Entry("_", "PmmdzqPrV", "vPwwTWBwg", "P"),
		)
	})

	When("getting priority of item", func() {
		DescribeTable("returns expected value", func(text string, expectedValue int) {
			priorityValue := GetItemPriorityValue(text)
			Expect(priorityValue).To(Equal(expectedValue))
		},

			Entry("_", "a", 1),  //65
			Entry("_", "b", 2),  //66
			Entry("_", "c", 3),  //67
			Entry("_", "y", 25), //89
			Entry("_", "z", 26), //90

			Entry("_", "A", 27),
			Entry("_", "B", 28),
			Entry("_", "C", 29),
			Entry("_", "D", 30),
			Entry("_", "Z", 52),
			Entry("_", "p", 16),
			Entry("_", "L", 38),
			Entry("_", "P", 42),
			Entry("_", "v", 22),
			Entry("_", "t", 20),
			Entry("_", "s", 19),
		)
	})
})

var _ = Describe("Day3 - Star1", func() {
	When("inspecting each rucksack", func() {
		It("returns expected duplicated item priority value", func() {
			Expect(GetDuplicatedItemValue("vJrwpWtwJgWrhcsFMMfFFhFp")).To(Equal(16))
			Expect(GetDuplicatedItemValue("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")).To(Equal(38))
			Expect(GetDuplicatedItemValue("PmmdzqPrVvPwwTWBwg")).To(Equal(42))
			Expect(GetDuplicatedItemValue("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")).To(Equal(22))
			Expect(GetDuplicatedItemValue("ttgJtRGJQctTZtZT")).To(Equal(20))
			Expect(GetDuplicatedItemValue("CrZsJsPPZsGzwwsLwLmpwMDw")).To(Equal(19))
		})
	})

	When("processed input file", func() {
		It("returns star1 puzzle value", func() {
			Expect(Star1("input.txt")).To(Equal(7793))
		})
	})
})

var _ = Describe("Day3 - Star2 - Utils", func() {
	When("finding common  item", func() {
		DescribeTable("returns common intem", func(content1 string, content2 string, content3 string, expectedItem string) {
			matchingItem := GetMatchingCharacter(content1, content2, content3)
			Expect(matchingItem).To(Equal(expectedItem))

		},
			Entry("_", "vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "r"),
			Entry("_", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw", "Z"),
		)
	})
})

var _ = Describe("Day3 - Star2", func() {
	When("processed input file", func() {
		It("returns star2 puzzle value", func() {
			Expect(Star2("input.txt")).To(Equal(2499))
		})
	})
})
