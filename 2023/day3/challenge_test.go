package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {

		It("is able to determine if a rune/string is a number", func() {
			Expect(IsRuneNumber('.')).To(BeFalse())
			Expect(IsRuneNumber('0')).To(BeTrue())
			Expect(IsRuneNumber('1')).To(BeTrue())
			Expect(IsRuneNumber('a')).To(BeFalse())

			Expect(GetNumber(".")).To(Equal(-1))
			Expect(GetNumber("1")).To(Equal(1))
			Expect(GetNumber("350")).To(Equal(350))
			Expect(GetNumber("a")).To(Equal(-1))
			Expect(GetNumber("1111")).To(Equal(1111))

		})

		It("is able to determine if a char is a symbol", func() {
			Expect(IsSymbol('.')).To(BeFalse())
			Expect(IsSymbol('1')).To(BeFalse())
			Expect(IsSymbol('a')).To(BeTrue())
			Expect(IsSymbol('#')).To(BeTrue())
			Expect(IsSymbol('+')).To(BeTrue())
			Expect(IsSymbol('*')).To(BeTrue())
			Expect(IsSymbol('$')).To(BeTrue())
		})

		When("given sample", func() {

			It("is able to determine combined numbers and its positions", func() {
				Expect(GetNumbersAndPositions("star1_sample.txt")).To(Equal([]NumberAndPositions{
					{
						number: 467,
						positions: []Position{
							{row: 0, column: 0},
							{row: 0, column: 1},
							{row: 0, column: 2},
						}}, {

						number: 114,
						positions: []Position{
							{row: 0, column: 5},
							{row: 0, column: 6},
							{row: 0, column: 7},
						}}, {

						number: 35,
						positions: []Position{
							{row: 2, column: 2},
							{row: 2, column: 3},
						}}, {

						number: 633,
						positions: []Position{
							{row: 2, column: 6},
							{row: 2, column: 7},
							{row: 2, column: 8},
						}}, {

						number: 617,
						positions: []Position{
							{row: 4, column: 0},
							{row: 4, column: 1},
							{row: 4, column: 2},
						}}, {

						number: 58,
						positions: []Position{
							{row: 5, column: 7},
							{row: 5, column: 8},
						}}, {

						number: 592,
						positions: []Position{
							{row: 6, column: 2},
							{row: 6, column: 3},
							{row: 6, column: 4},
						}}, {

						number: 755,
						positions: []Position{
							{row: 7, column: 6},
							{row: 7, column: 7},
							{row: 7, column: 8},
						}}, {

						number: 664,
						positions: []Position{
							{row: 9, column: 1},
							{row: 9, column: 2},
							{row: 9, column: 3},
						}}, {

						number: 598,
						positions: []Position{
							{row: 9, column: 5},
							{row: 9, column: 6},
							{row: 9, column: 7},
						}},
				}))
			})

			It("is able to determine symbol position", func() {
				Expect(GetSymbolsPosition("star1_sample.txt")).To(Equal([]Position{
					{
						row:    1,
						column: 3,
					}, {
						row:    3,
						column: 6,
					}, {
						row:    4,
						column: 3,
					}, {
						row:    5,
						column: 5,
					}, {
						row:    8,
						column: 3,
					}, {
						row:    8,
						column: 5,
					},
				}))
			})

			It("is able to determine if a number is adjacent to a symbol", func() {
				symbolPositions := GetSymbolsPosition("star1_sample.txt")

				// 467
				Expect(AreAdjacent(symbolPositions, []Position{
					{row: 0, column: 0},
					{row: 0, column: 1},
					{row: 0, column: 2},
				})).To(BeTrue())

				// 114
				Expect(AreAdjacent(symbolPositions, []Position{
					{row: 0, column: 5},
					{row: 0, column: 6},
					{row: 0, column: 7},
				})).To(BeFalse())

				// 35
				Expect(AreAdjacent(symbolPositions, []Position{
					{row: 2, column: 2},
					{row: 2, column: 3},
				})).To(BeTrue())

				// 633
				Expect(AreAdjacent(symbolPositions, []Position{
					{row: 2, column: 6},
					{row: 2, column: 7},
					{row: 2, column: 8},
				})).To(BeTrue())

				// 617
				Expect(AreAdjacent(symbolPositions, []Position{
					{row: 4, column: 0},
					{row: 4, column: 1},
					{row: 4, column: 2},
				})).To(BeTrue())

				// 58
				Expect(AreAdjacent(symbolPositions, []Position{
					{row: 5, column: 7},
					{row: 5, column: 8},
				})).To(BeFalse())

				//number: 58,
				//	positions: []Position{
				//		{row: 5, column: 7},
				//		{row: 5, column: 8},
				//	}}, {
				//
				//number: 592,
				//	positions: []Position{
				//		{row: 6, column: 2},
				//		{row: 6, column: 3},
				//		{row: 6, column: 4},
				//	}}, {
				//
				//number: 755,
				//	positions: []Position{
				//		{row: 7, column: 6},
				//		{row: 7, column: 7},
				//		{row: 7, column: 8},
				//	}}, {
				//
				//number: 664,
				//	positions: []Position{
				//		{row: 9, column: 1},
				//		{row: 9, column: 2},
				//		{row: 9, column: 3},
				//	}}, {
				//
				//number: 598,
				//	positions: []Position{
				//		{row: 9, column: 5},
				//		{row: 9, column: 6},
				//		{row: 9, column: 7},
				//	}},

			})

			It("resolves", func() {
				Expect(Star1("star1_sample.txt")).To(Equal(4361))
			})
		})

		When("given input", func() {
			It("resolves", func() {
				Expect(Star1("star1_input.txt")).To(Equal(0))
			})
		})
	})
	FContext("Star2", func() {
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star2("star1_sample.txt")).To(Equal(467835))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Star2("err_sample.txt")

				result := Star2("star1_input.txt")

				Expect(result).To(Not(Equal(56442513)))
				Expect(result).To(Not(Equal(55589908)))
				Expect(result).To(Not(Equal(76129174)))
				Expect(result).To(Not(Equal(82113134)))
				Expect(result).To(Not(Equal(82824352)))
				Expect(result).To(Not(Equal(82965739)))
				Expect(result).To(Not(Equal(83236519)))
				Expect(result).To(Not(Equal(96085965)))
				Expect(result).To(Not(Equal(295327471)))

				Expect(result).To(Equal(82824352))
			})
		})
	})
})
