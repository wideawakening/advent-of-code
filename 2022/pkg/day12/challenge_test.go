package day12

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("loading data map", func() {
		It("loads nodes into map", func() {
			LoadMap("sample.txt")
			Expect(1).To(Equal(1))
		})
	})

	When("searching for start position", func() {
		It("return start position", func() {
			dataMap, err := LoadMap("sample.txt")
			Expect(err).To(Not(HaveOccurred()))

			position, err := GetPositionFor(dataMap, "S")
			Expect(err).To(Not(HaveOccurred()))
			Expect(position.position).To(Equal([2]int{0, 0}))
		})

		It("return start position", func() {
			dataMap, err := LoadMap("input.txt")
			Expect(err).To(Not(HaveOccurred()))

			startNode, err := GetPositionFor(dataMap, "S")
			Expect(err).To(Not(HaveOccurred()))
			Expect(startNode.position).To(Equal([2]int{20, 0}))
		})
	})

	When("testing if a destination node is jumpleable from a source node", func() {
		It("checks for position from S", func() {
			node := NewNode('S', [2]int{0, 0})
			Expect(node.IsJumpable(*NewNode('a', [2]int{5, 5}))).To(Equal(false))
			Expect(node.IsJumpable(*NewNode('a', [2]int{0, 5}))).To(Equal(false))
			Expect(node.IsJumpable(*NewNode('a', [2]int{5, 0}))).To(Equal(false))
		})

		It("check for signal from S", func() {
			node := NewNode('S', [2]int{0, 0})
			Expect(node.IsJumpable(*NewNode('a', [2]int{0, 1}))).To(Equal(true))
			Expect(node.IsJumpable(*NewNode('a', [2]int{1, 0}))).To(Equal(true))
			Expect(node.IsJumpable(*NewNode('a', [2]int{1, 1}))).To(Equal(true))
		})

		It("check for signal from any non S node", func() {
			node := NewNode('a', [2]int{0, 0})
			Expect(node.IsJumpable(*NewNode('b', [2]int{1, 1}))).To(Equal(true))
			Expect(node.IsJumpable(*NewNode('c', [2]int{1, 1}))).To(Equal(false))
		})

		It("sample data", func() {
			node := NewNode('b', [2]int{4, 1})
			Expect(node.IsJumpable(*NewNode('c', [2]int{3, 1}))).To(Equal(true))
			Expect(node.IsJumpable(*NewNode('d', [2]int{3, 2}))).To(Equal(false))
			Expect(node.IsJumpable(*NewNode('a', [2]int{3, 0}))).To(Equal(false))
		})

	})

	When("looking for jumpable nodes", func() {
		It("returns next jumpable nodes based on signal and position", func() {
			dataMap, err := LoadMap("sample.txt")
			Expect(err).NotTo(HaveOccurred())

			jumpableNodes := GetJumpableNodes(dataMap, [2]int{0, 0})

			Expect(len(jumpableNodes)).To(Equal(2))
			Expect(jumpableNodes[0].position).To(Equal([2]int{1, 0}))
			Expect(jumpableNodes[1].position).To(Equal([2]int{0, 1}))

		})
	})
})

var _ = Describe("challenge", func() {
	When("given sample", func() {
		It("resolve ...", func() {
			value := Star1("sample.txt")
			Expect(value).To(Equal(13))
		})
	})
	XWhen("given input file", func() {
		It("resolve ...", func() {
			Expect(Star1("input.txt")).To(Equal(0))
		})
	})
})
