package day6

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("given the transmission", func() {
		DescribeTable("it identifies the character after SOP", func(transmission string, expectedIndex int) {

			sopStartIndex, _ := FindSOPIndex(transmission, 4)
			Expect(sopStartIndex).To(Equal(expectedIndex))

		},
			Entry("_", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7),
			Entry("_", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5),
			Entry("_", "nppdvjthqldpwncqszvftbrmjlhg", 6),
			Entry("_", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10),
			Entry("_", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11),
		)
	})
})

var _ = Describe("challenge - star1/2", func() {
	When("processed input file", func() {
		It("returns star1 puzzle value", func() {
			index, _ := Star1("input.txt", 4)
			Expect(index).To(Equal(1109))
		})

		It("returns star1 puzzle value", func() {
			index, _ := Star1("input.txt", 14)
			Expect(index).To(Equal(3965))
		})
	})
})
