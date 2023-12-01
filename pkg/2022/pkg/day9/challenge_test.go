package day9

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("given a steps", func() {
		DescribeTable("it returns the position", func(position Position, movement *Movement, expectedFinalPosition Position) {
			resultPosition := Move(position, []*Movement{movement})
			Expect(resultPosition.x).To(Equal(expectedFinalPosition.x))
			Expect(resultPosition.y).To(Equal(expectedFinalPosition.y))
		},
			Entry("_", Position{0, 0}, NewMovement("R", 4), Position{x: 4, y: 0}),
			Entry("_", Position{4, 0}, NewMovement("L", 4), Position{x: 0, y: 0}),
			Entry("_", Position{0, 0}, NewMovement("U", 4), Position{x: 0, y: 4}),
			Entry("_", Position{0, 4}, NewMovement("D", 4), Position{x: 0, y: 0}),
		)
	})

	When("given a sequence steps", func() {
		DescribeTable("it returns the HEAD position", func(position Position, movements []*Movement, expectedFinalPosition Position) {
			resultPosition := Move(position, movements)
			Expect(resultPosition).To(Equal(expectedFinalPosition))

		},
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4)}, Position{x: 4, y: 0}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4)}, Position{x: 4, y: 4}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4), NewMovement("L", 3)}, Position{x: 1, y: 4}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4), NewMovement("L", 3), NewMovement("D", 1)}, Position{x: 1, y: 3}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4), NewMovement("L", 3), NewMovement("D", 1), NewMovement("R", 4)}, Position{x: 5, y: 3}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4), NewMovement("L", 3), NewMovement("D", 1), NewMovement("R", 4), NewMovement("D", 1)}, Position{x: 5, y: 2}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4), NewMovement("L", 3), NewMovement("D", 1), NewMovement("R", 4), NewMovement("D", 1), NewMovement("L", 5)}, Position{x: 0, y: 2}),
			Entry("_", Position{0, 0}, []*Movement{NewMovement("R", 4), NewMovement("U", 4), NewMovement("L", 3), NewMovement("D", 1), NewMovement("R", 4), NewMovement("D", 1), NewMovement("L", 5), NewMovement("R", 2)}, Position{x: 2, y: 2}),
		)
	})

	When("given the head position and current tail position", func() {
		DescribeTable("it calculates new TAIL position", func(headPosition Position, tailPosition Position, newTailPositions []Position) {
			tailMovements := GetTailToHeadMovements(headPosition, tailPosition)

			Expect(tailMovements).To(Equal(newTailPositions))
		},
			Entry("_", Position{0, 0}, Position{0, 0}, []Position{{0, 0}}),
			Entry("_", Position{4, 0}, Position{0, 0}, []Position{{x: 1, y: 0}, {x: 2, y: 0}, {x: 3, y: 0}}),
			Entry("_", Position{4, 4}, Position{3, 0}, []Position{{x: 4, y: 1}, {x: 4, y: 2}, {x: 4, y: 3}}),
			Entry("_", Position{1, 4}, Position{4, 3}, []Position{{x: 3, y: 4}, {x: 2, y: 4}}),
			Entry("_", Position{1, 3}, Position{2, 4}, []Position{{2, 4}}),
			Entry("_", Position{5, 3}, Position{2, 4}, []Position{{x: 3, y: 3}, {x: 4, y: 3}}),
			Entry("_", Position{5, 2}, Position{4, 3}, []Position{{4, 3}}),
			Entry("_", Position{0, 2}, Position{4, 3}, []Position{{x: 3, y: 2}, {x: 2, y: 2}, {x: 1, y: 2}}),
			Entry("_", Position{2, 2}, Position{1, 2}, []Position{{1, 2}}),

			Entry("tail does not move - all angles", Position{0, 0}, Position{0, 0}, []Position{{0, 0}}),
			Entry("tail does not move - all angles", Position{0, 0}, Position{1, 0}, []Position{{1, 0}}),
			Entry("tail does not move - all angles", Position{0, 0}, Position{0, 1}, []Position{{0, 1}}),
			Entry("tail does not move - all angles", Position{0, 0}, Position{1, 1}, []Position{{1, 1}}),
			Entry("tail does not move - all angles", Position{0, 0}, Position{-1, 0}, []Position{{-1, 0}}),
			Entry("tail does not move - all angles", Position{0, 0}, Position{0, -1}, []Position{{0, -1}}),
			Entry("tail does not move - all angles", Position{0, 0}, Position{-1, -1}, []Position{{-1, -1}}),

			Entry("tail moves 1 - all angles", Position{0, 0}, Position{0, 2}, []Position{{0, 1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{1, 2}, []Position{{0, 1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{2, 2}, []Position{{1, 1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{2, 1}, []Position{{1, 0}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{2, 0}, []Position{{1, 0}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{2, -1}, []Position{{1, 0}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{2, -2}, []Position{{1, -1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-1, -2}, []Position{{0, -1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{0, -2}, []Position{{0, -1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-1, -2}, []Position{{0, -1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-2, -2}, []Position{{-1, -1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-2, -1}, []Position{{-1, 0}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-2, 0}, []Position{{-1, 0}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-2, 1}, []Position{{-1, 0}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-2, 2}, []Position{{-1, 1}}),
			Entry("tail moves 1 - all angles", Position{0, 0}, Position{-1, 2}, []Position{{0, 1}}),

			Entry("tail does not move - all angles", Position{0, 1}, Position{0, 2}, []Position{{0, 2}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{1, 2}, []Position{{1, 2}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{1, 1}, []Position{{1, 1}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{0, 1}, []Position{{0, 1}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{0, 0}, []Position{{0, 0}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{-1, 0}, []Position{{-1, 0}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{-1, 1}, []Position{{-1, 1}}),
			Entry("tail does not move - all angles", Position{0, 1}, Position{-1, 2}, []Position{{-1, 2}}),

			Entry("from test2", Position{-6, 2}, Position{-3, 2}, []Position{{-4, 2}, {-5, 2}}),
			Entry("from test2", Position{-3, 0}, Position{-4, 2}, []Position{{-3, 1}}),
			Entry("from input", Position{2, -21}, Position{3, -18}, []Position{{2, -19}, {2, -20}}),
			Entry("from input", Position{-6, -51}, Position{-24, -50}, []Position{{-23, -51}, {-22, -51}, {-21, -51}, {-20, -51}, {-19, -51}, {-18, -51}, {-17, -51}, {-16, -51}, {-15, -51}, {-14, -51}, {-13, -51}, {-12, -51}, {-11, -51}, {-10, -51}, {-9, -51}, {-8, -51}, {-7, -51}}),
			Entry("from input", Position{-87, 0}, Position{-87, 17}, []Position{{-87, 16}, {-87, 15}, {-87, 14}, {-87, 13}, {-87, 12}, {-87, 11}, {-87, 10}, {-87, 9}, {-87, 8}, {-87, 7}, {-87, 6}, {-87, 5}, {-87, 4}, {-87, 3}, {-87, 2}, {-87, 1}}),
		)
	})
	When("given the list of the elf positions", func() {
		It("returns the number of UNIQUE positions", func() {
			Expect(GetUniquePosition([]Position{{x: 0, y: 0}, {x: 4, y: 0}, {4, 4}, {1, 4}, {1, 3}, {5, 3}, {5, 2}, {2, 2}, {5, 0}})).To(Equal(9))
			Expect(GetUniquePosition([]Position{{0, 0}, {0, 0}, {5, 3}, {3, 5}, {3, 5}})).To(Equal(3))
			Expect(GetUniquePosition([]Position{{0, 0}, {-1, 0}, {0, -1}, {1, 1}, {0, 0}, {-1, 0}, {0, -1}, {1, 1}})).To(Equal(4))

		})
	})
	FWhen("adding tail positions", func() {
		It("increments number of tail positions", func() {
			ret := AddTailPositions([]Position{{0, 0}}, []Position{})
			Expect(len(ret)).To(Equal(1))

			ret = AddTailPositions([]Position{{0, 0}}, []Position{{0, 0}})
			Expect(len(ret)).To(Equal(2))

			ret = AddTailPositions([]Position{{0, 0}, {1, -1}}, []Position{{0, 0}})
			Expect(ret).To(Equal([]Position{{0, 0}, {1, -1}, {0, 0}}))
		})
	})
})

var _ = Describe("challenge", func() {
	When("given sample file", func() {
		It("resolve the number of position tail visited at least once", func() {
			Expect(Star1("sample.txt")).To(Equal(13))
		})
	})

	When("given sample file", func() {
		It("resolve the number of position tail visited at least once", func() {
			Expect(Star1("sample2.txt")).To(Equal(90))
		})
	})

	When("given test1 file", func() {
		It("resolve the number of position tail visited at least once", func() {
			Expect(Star1("test1.txt")).To(Equal(5))
		})
	})
	When("given test2 file", func() {
		It("resolve the number of position tail visited at least once", func() {
			Expect(Star1("test2.txt")).To(Equal(9))
		})
	})
	When("given test3 file", func() {
		It("resolve the number of position tail visited at least once", func() {
			Expect(Star1("test3.txt")).To(Equal(65))
		})
	})
	When("given input file", func() {
		It("resolve the number of position tail visited at least once", func() {
			Expect(Star1("input.txt")).NotTo(Equal(5922))
			Expect(Star1("input.txt")).NotTo(Equal(5925))
			Expect(Star1("input.txt")).NotTo(Equal(5924))
			Expect(Star1("input.txt")).NotTo(Equal(5923))
			Expect(Star1("input.txt")).To(Equal(0))
		})
	})
})
