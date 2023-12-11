package challenge

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	It("FindStart", func() {
		mapPoint := FindStart("star_sample.txt")
		Expect(mapPoint.x).To(Equal(1))
		Expect(mapPoint.y).To(Equal(1))

		// TODO more tests if required
	})

	// TODO test IsValidAdjacent id needed

	It("GetAdjacentValidMoves", func() {
		validMoves := GetAdjacentValidMoves("star_sample.txt", 1, 1)
		Expect(validMoves).To(Equal([]MapPoint{MapPoint{2, 1, "|"}, MapPoint{1, 2, "-"}}))

		// TODO more tests if required
	})

	It("GetsLoop", func() {
		loop := GetLoop("star_sample.txt")
		Expect(loop).To(Equal([]MapPoint{
			{1, 1, "S"},
			{2, 1, "|"},
			{3, 1, "L"},
			{3, 2, "-"},
			{3, 3, "J"},
			{2, 3, "|"},
			{1, 3, "7"},
			{1, 2, "-"},
			{1, 1, "S"}}))

		Expect(float64((len(loop)) / 2)).To(Equal(float64(4)))
	})

	Context("Star1", func() {
		When("given sample", func() {
			It("resolves", func() {
				loop := GetLoop("star_sample2.txt")
				Expect(math.Ceil(float64(len(loop) / 2))).To(Equal(float64(8)))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				loop := GetLoop("star_input.txt")
				farestPoint := math.Ceil(float64(len(loop) / 2))
				Expect(farestPoint).NotTo(Equal(float64(6883)))
				Expect(farestPoint).NotTo(Equal(float64(6884)))
				Expect(farestPoint).NotTo(Equal(float64(5119)))
				Expect(farestPoint).NotTo(Equal(float64(6191)))
				Expect(farestPoint).NotTo(Equal(float64(6192)))
				Expect(farestPoint).NotTo(Equal(float64(6193)))
				Expect(farestPoint).To(Equal(float64(6882)))
			})
		})
	})
	FContext("Star2", func() {
		When("given sample", func() {
			It("resolves", func() {
				//Expect(FindInsideObjectsWithinLoop("star2_sample1.txt")).To(Equal(4))
				Expect(FindInsideObjectsWithinLoop("star2_sample2.txt")).To(Equal(8))

			})
		})
		//XWhen("given input", func() {
		//	It("resolves", func() {
		//		Expect(Star1("star_input.txt")).To(Equal(0))
		//	})
		//})
	})
})
