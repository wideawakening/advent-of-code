package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {

	/*
			 	n number of red, green, blue cubes in a bag
		 	 	ID(game number): random samples like x red, y green, z blue; random sample x red, y green, z blue; ...
	*/
	Context("Star1", func() {
		/*
			const MinRed = 12
			const MinGreen = 13
			const MinBlue = 14
		*/
		It("identifies each game colors", func() {

			//Expect(GetGameID("Game 1")).To(Equal(1))
			//Expect(GetGameID("Game 10")).To(Equal(10))

			Expect(GetReveal("3 blue, 4 red")).To(Equal(Reveal{4, 0, 3}))
			Expect(GetReveal("1 red")).To(Equal(Reveal{1, 0, 0}))
			Expect(GetReveal("")).To(Equal(Reveal{0, 0, 0}))
			Expect(GetReveal("4 green, 1 red, 2 blue")).To(Equal(Reveal{1, 4, 2}))

			Expect(GetGame("Game 1: 3 blue, 4 red; 1 red")).To(Equal(Game{1, []Reveal{{4, 0, 3}, {1, 0, 0}}}))
		})

		It("resolves if is possible to play", func() {
			Expect(PossibleGame(1, 1, 1, 1)).To(Equal(true))
			Expect(PossibleGame(1, 13, 1, 1)).To(Equal(false))
			Expect(PossibleGame(1, 1, 14, 1)).To(Equal(false))
			Expect(PossibleGame(1, 1, 1, 15)).To(Equal(false))

		})

		When("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star1_sample.txt")).To(Equal(8))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star1("star1_input.txt")).To(Equal(2771))
			})
		})
	})

	Context("Star2", func() {
		It("gets min cubes from reveal", func() {
			Expect(GetMinPossibleFromReveal(Reveal{1, 1, 1}, Reveal{1, 1, 1})).To(Equal(Reveal{red: 1, green: 1, blue: 1}))
			Expect(GetMinPossibleFromReveal(Reveal{2, 1, 1}, Reveal{1, 1, 1})).To(Equal(Reveal{red: 2, green: 1, blue: 1}))
			Expect(GetMinPossibleFromReveal(Reveal{2, 2, 1}, Reveal{1, 1, 1})).To(Equal(Reveal{red: 2, green: 2, blue: 1}))
			Expect(GetMinPossibleFromReveal(Reveal{2, 2, 2}, Reveal{1, 1, 1})).To(Equal(Reveal{red: 2, green: 2, blue: 2}))
		})

		It("gets min cubes from game", func() {
			Expect(GetMinPossibleFromGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")).To(Equal(Reveal{red: 4, green: 2, blue: 6}))
			Expect(GetMinPossibleFromGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")).To(Equal(Reveal{red: 1, green: 3, blue: 4}))
			Expect(GetMinPossibleFromGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")).To(Equal(Reveal{red: 20, green: 13, blue: 6}))
			Expect(GetMinPossibleFromGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")).To(Equal(Reveal{red: 14, green: 3, blue: 15}))
			Expect(GetMinPossibleFromGame("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")).To(Equal(Reveal{red: 6, green: 3, blue: 2}))
		})

		When("given sample", func() {
			It("resolves", func() {
				Expect(Star2("star1_sample.txt")).To(Equal(2286))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star2("star1_input.txt")).To(Equal(70924))
			})
		})
	})
})
