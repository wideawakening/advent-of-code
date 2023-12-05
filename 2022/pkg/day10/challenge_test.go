package day10

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("given a sequence", func() {
		DescribeTable("calculates the regX value", func(inputFile string, expectedValue int) {
			_, _, regx := Star1(inputFile)
			Expect(regx).To(Equal(expectedValue))
		},
			Entry("_", "test1.txt", -1),
			Entry("_", "test3.txt", 32),
			Entry("_", "test3.txt", 18),
			Entry("_", "sample.txt", 17),
		)
	})

	When("given a pre and post cycle value", func() {
		DescribeTable("determines if is time to calculate signal strength", func(cyclePrevValue int, cyclePostValue int, expectedDetermineStrength bool, expectedCycle int) {
			determineStrength, cycle := IsCalculateSignalStrength(cyclePrevValue, cyclePostValue)
			Expect(determineStrength).To(Equal(expectedDetermineStrength))
			Expect(cycle).To(Equal(expectedCycle))
		},
			Entry("_", 1, 2, false, 0),
			Entry("_", 19, 20, true, 20),
			Entry("_", 20, 21, false, 0),

			Entry("_", 40, 41, false, 0),
			Entry("_", 58, 60, true, 60),
			Entry("_", 59, 60, true, 60),
			Entry("_", 60, 61, false, 0),
			Entry("_", 60, 62, false, 0),
			Entry("_", 61, 62, false, 0),
		)
	})
})

var _ = Describe("challenge - star2 - utils", func() {
	When("given a regx and cycle values", func() {
		DescribeTable("determines whether to lit pixel or not, and if to breakline", func(cycle int, regx int, _litPixel bool, _breakLine bool) {
			litPixel, breakLine := ProcessCRT(cycle, regx)
			Expect(litPixel).To(Equal(_litPixel))
			Expect(breakLine).To(Equal(_breakLine))
		},
			Entry("_", 1, 1, true, false), // add has 2 cycles
			Entry("_", 2, 1, true, false),
			Entry("_", 3, 16, false, false),
			Entry("_", 4, 16, false, false), // add 2 cycles
			Entry("_", 5, 16, false, false),
			Entry("_", 6, 5, true, false),
		)
	})
	When("calculating crt position", func() {
		DescribeTable("converts cycle value in a 40positions per line/cycle display", func(cycle int, crtPosition int) {
			Expect(ConvertCycleToCRTPosition(cycle)).To(Equal(crtPosition))
		},
			Entry("_", 1, 1),
			Entry("_", 2, 2),
			Entry("_", 40, 40),
			Entry("_", 41, 1),
			Entry("_", 80, 40),
			Entry("_", 81, 1),
		)
	})
})

var _ = Describe("challenge", func() {
	When("given sample", func() {
		It("returns signal strength sum", func() {
			signalStrengthSum, _, _ := Star1("sample.txt")
			Expect(signalStrengthSum).To(Equal(13140))
		})
	})
	When("given input file for Star1", func() {
		It("resolve ...", func() {
			signalStrengthSum, _, _ := Star1("input.txt")
			Expect(signalStrengthSum).To(Equal(12740))
		})
	})

	When("given sample file for Star2", func() {
		It("it prints the result", func() {
			fmt.Println()
			Star2("sample.txt")
			/*
				##..##..##..##..##..##..##..##..##..##..
				###...###...###...###...###...###...###.
				####....####....####....####....####....
				#####.....#####.....#####.....#####.....
				######......######......######......####
				#######.......#######.......#######.....
			*/
		})
	})

	FWhen("given input file for Star2", func() {
		It("it prints the result", func() {
			fmt.Println()
			Star2("input.txt")
			/*
				███░░███░░███░░░██░░███░░░██░░░██░░████░
				█░░█░█░░█░█░░█░█░░█░█░░█░█░░█░█░░█░█░░░░
				█░░█░███░░█░░█░█░░█░█░░█░█░░█░█░░░░███░░
				███░░█░░█░███░░████░███░░████░█░██░█░░░░
				█░█░░█░░█░█░░░░█░░█░█░█░░█░░█░█░░█░█░░░░
				█░░█░███░░█░░░░█░░█░█░░█░█░░█░░███░█░░░░
			*/
		})
	})
})
