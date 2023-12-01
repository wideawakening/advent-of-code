package day5

import (
	"container/list"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("pushed items into fifo queue", func() {
		It("returns items in LIFO order", func() {
			stackA := list.New()
			stackA.PushFront(1)
			stackA.PushFront(2)
			stackA.PushFront(3)
			stackA.PushFront(4)
			Expect(stackA.Remove(stackA.Front())).To(Equal(4))
			Expect(stackA.Remove(stackA.Front())).To(Equal(3))

			stackB := list.New()
			for i := 0; i < stackA.Len(); i++ {
				stackB.PushFront(stackA.Remove(stackA.Front()))
			}

			Expect(stackB.Front().Value).To(Equal(2))

		})
	})
	When("given a container list", func() {
		It("returns containerList", func() {
			stacks := CreateStacks([][]string{{"N", "Z"}, {"A", "B"}, {}})
			Expect(len(stacks)).To(Equal(3))

			Expect(stacks[0].Front().Value).To(Equal("N"))
			Expect(stacks[1].Front().Value).To(Equal("A"))
			Expect(stacks[2].Front()).To(BeNil())
		})
	})

	When("asking for a movement in piles in LIFO mode", func() {
		DescribeTable("piles are moved around", func(crates [][]string, moves []int, expectedCrate [][]string) {
			listPiles := CreateStacks(crates)
			resultList := Move(false, listPiles, moves)
			expectedList := CreateStacks(expectedCrate)

			for i := 0; i < len(resultList); i++ {
				Expect(resultList[i].Len()).To(Equal(expectedList[i].Len()))
				if resultList[i].Len() > 0 {
					Expect(resultList[i].Remove(resultList[i].Front())).To(Equal(expectedList[i].Remove(expectedList[i].Front())))
				}
			}

		},
			Entry("_", [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}, []int{1, 2, 1}, [][]string{{"D", "N", "Z"}, {"C", "M"}, {"P"}}),
			Entry("_", [][]string{{"D", "N", "Z"}, {"C", "M"}, {"P"}}, []int{3, 1, 3}, [][]string{{}, {"C", "M"}, {"Z", "N", "D", "P"}}),
			Entry("_", [][]string{{}, {"C", "M"}, {"Z", "N", "D", "P"}}, []int{2, 2, 1}, [][]string{{"M", "C"}, {}, {"Z", "N", "D", "P"}}),
			Entry("_", [][]string{{"M", "C"}, {}, {"Z", "N", "D", "P"}}, []int{1, 1, 2}, [][]string{{"C"}, {"M"}, {"Z", "N", "D", "P"}}),
		)
	})

	When("asking for a movement in piles in FIFO mode", func() {
		DescribeTable("piles are moved around", func(crates [][]string, moves []int, expectedCrate [][]string) {
			listPiles := CreateStacks(crates)
			resultList := Move(true, listPiles, moves)
			expectedList := CreateStacks(expectedCrate)

			for i := 0; i < len(resultList); i++ {
				Expect(resultList[i].Len()).To(Equal(expectedList[i].Len()))
				if resultList[i].Len() > 0 {
					Expect(resultList[i].Remove(resultList[i].Front())).To(Equal(expectedList[i].Remove(expectedList[i].Front())))
				}
			}

		},
			Entry("_", [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}, []int{1, 2, 1}, [][]string{{"D", "N", "Z"}, {"C", "M"}, {"P"}}),
			Entry("_", [][]string{{"D", "N", "Z"}, {"C", "M"}, {"P"}}, []int{3, 1, 3}, [][]string{{}, {"C", "M"}, {"D", "N", "Z", "P"}}),
			Entry("_", [][]string{{}, {"C", "M"}, {"Z", "N", "D", "P"}}, []int{2, 2, 1}, [][]string{{"C", "M"}, {}, {"Z", "N", "D", "P"}}),
			Entry("_", [][]string{{"M", "C"}, {}, {"Z", "N", "D", "P"}}, []int{1, 1, 2}, [][]string{{"C"}, {"M"}, {"Z", "N", "D", "P"}}),
		)
	})
})

var _ = Describe("challenge - star1", func() {
	When("processed input file", func() {
		It("returns star1 puzzle value", func() {
			//XExpect(Star1(false, "mod_input.txt")).To(Equal("TPGVQPFDH"))
			Expect(Star1(true, "mod_input.txt")).To(Equal("TPGVQPFDH"))
		})
	})
})
