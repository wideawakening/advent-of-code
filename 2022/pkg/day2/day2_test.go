package day2

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day2 - Shape1", func() {
	/*
								shape score
		  	A/X - rock			- 1 pts
			B/Y - paper			- 2 pts
			C/Z - scissors		- 3 pts
	*/
	When("Evaluating Shape points", func() {
		It("returns expected points", func() {
			Expect(GetShapePoints("A")).To(Equal(1))
			Expect(GetShapePoints("X")).To(Equal(1))
			Expect(GetShapePoints("B")).To(Equal(2))
			Expect(GetShapePoints("Y")).To(Equal(2))
			Expect(GetShapePoints("C")).To(Equal(3))
			Expect(GetShapePoints("Z")).To(Equal(3))
		})
	})

	/*
		   round outcome score
		   0 - lost
		   3 - draw
		   6  - win

			win/loose
			C - X		6 pts - piedra - tijera
			A - Z		0 pts - loose

			A - Y		6 pts - papel	- piedra
			B - X		0 pts - papel - piedra

			B - Z		6 pts - tijera - papel

			draws	-	3 pts
			A - X
			B - Y
			C - Z
	*/
	When("Evaluating Round points", func() {
		It("draws return 3 points", func() {
			Expect(EvaluateRound("A", "X")).To(Equal(3))
			Expect(EvaluateRound("B", "Y")).To(Equal(3))
			Expect(EvaluateRound("C", "Z")).To(Equal(3))
		})

		It("looses return 0 points", func() {
			Expect(EvaluateRound("A", "Z")).To(Equal(0))
			Expect(EvaluateRound("B", "X")).To(Equal(0))
			Expect(EvaluateRound("C", "Y")).To(Equal(0))
		})

		It("wins return 6 points", func() {
			Expect(EvaluateRound("C", "X")).To(Equal(6))
			Expect(EvaluateRound("A", "Y")).To(Equal(6))
			Expect(EvaluateRound("B", "Z")).To(Equal(6))
		})

		When("Evaluating Round Shape points", func() {
			It("returns expected points", func() {
				// A = piedra; Y = papel
				Expect(EvaluatePointsByShape("A", "Y")).To(Equal(8))
				// B = papel; X = piedra
				Expect(EvaluatePointsByShape("B", "X")).To(Equal(1))
				Expect(EvaluatePointsByShape("C", "Z")).To(Equal(6))
			})
		})

		When("Star1 input", func() {
			It("return total score", func() {
				Expect(Star1("./input.txt", true)).To(Equal(14827))
			})
		})
	})
})

var _ = Describe("Day2 - Shape1", func() {

	When("given action point", func() {
		It("returns the points of each action", func() {
			Expect(GetActionPoints("X")).To(Equal(0))
			Expect(GetActionPoints("Y")).To(Equal(3))
			Expect(GetActionPoints("Z")).To(Equal(6))
		})
	})
	When("Evaluating Round Shape points", func() {
		It("returns expected points", func() {
			// A = piedra; Y = draw; resultado: piedra = 1 punto
			// 1 punto piedra + 3 puntos draw
			Expect(EvaluatePointsByAction("A", "Y")).To(Equal(4))
			Expect(EvaluatePointsByAction("B", "X")).To(Equal(1))
			Expect(EvaluatePointsByAction("C", "Z")).To(Equal(7))
		})
	})

	When("Star2 input", func() {
		It("return total score", func() {
			Expect(Star1("./input.txt", false)).To(Equal(13889))
		})
	})
})
