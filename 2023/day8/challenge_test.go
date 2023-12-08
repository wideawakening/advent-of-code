package challenge

import (
	_ "embed"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//go:embed star_sample1.txt
var sample1 string

//go:embed star_sample2.txt
var sample2 string

//go:embed star_input.txt
var input string

//go:embed star2_sample.txt
var star2Sample string

var _ = Describe("challenge", func() {
	Context("Star1", func() {
		When("given sample", func() {
			It("extracts elfMap", func() {
				Expect(sample1).NotTo(BeEmpty())
				elfMap := GetElfMap(sample1)
				Expect(elfMap.Instructions).To(Equal([]Instruction{"R", "L"}))
				Expect(len(elfMap.Nodes)).To(Equal(7))
			})
		})
		When("given sample", func() {
			It("resolves", func() {
				Expect(GetStepsToExit(GetElfMap(sample1))).To(Equal(2))
				Expect(GetStepsToExit(GetElfMap(sample2))).To(Equal(6))
			})
		})

		When("given input", func() {
			It("resolves", func() {
				Expect(GetStepsToExit(GetElfMap(input))).To(Equal(11567))
			})
		})
	})
	Context("Star2", func() {
		It("fetches start points", func() {
			startNodes := GetStartNodes(GetElfMap(star2Sample))
			for _, nodeName := range startNodes {
				Expect(string(nodeName[2])).To(Equal("A"))
				Expect(len(nodeName)).To(Equal(3))
			}
		})

		It("computes end", func() {
			Expect(AllNodesEnded([]string{"AAA", "AAB"})).To(BeFalse())
			Expect(AllNodesEnded([]string{"AZA", "AAA"})).To(BeFalse())
			Expect(AllNodesEnded([]string{"AAA", "AAZ"})).To(BeFalse())
			Expect(AllNodesEnded([]string{"AAZ", "AAZ"})).To(BeTrue())
			Expect(AllNodesEnded([]string{"11B", "22Z"})).To(BeFalse())

		})

		When("given sample", func() {
			It("resolves", func() {
				Expect(InneficientGetGhostStepsToExit(GetElfMap(star2Sample))).To(Equal(6))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				steps := GetGhostStepsToExit(GetElfMap(input))
				Expect(steps).NotTo(Equal(11159))
				Expect(steps).To(Equal(9858474970153)) // manual lcm of 6 nodes steps
			})
		})
	})
})
