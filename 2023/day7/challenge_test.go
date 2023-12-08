package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", Ordered, func() {
	Context("Star1", func() {

		It("validates hand", func() {
			Expect(IsValidHand(Hand("AAAAA"))).To(Equal(true))
			Expect(IsValidHand(Hand("AAAAAA"))).To(Equal(false))
			Expect(IsValidHand(Hand("AAAAB"))).To(Equal(false))
			Expect(IsValidHand(Hand("22221"))).To(Equal(false))
			Expect(IsValidHand(Hand("A"))).To(Equal(false))

		})
		It("computes hand type", func() {
			Expect(Borrowed_GetHandType(Hand("AAAAA"))).To(Equal(FiveOfAKind))
			Expect(Borrowed_GetHandType(Hand("AA8AA"))).To(Equal(FourOfAKind))
			Expect(Borrowed_GetHandType(Hand("23332"))).To(Equal(FullHouse))
			Expect(Borrowed_GetHandType(Hand("TTT98"))).To(Equal(ThreeOfAKind))
			Expect(Borrowed_GetHandType(Hand("23432"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("12233"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("22344"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("22334"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("A23A4"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("23456"))).To(Equal(HighCard))

			Expect(Borrowed_GetHandType(Hand("11223"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("12233"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("11223"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("11233"))).To(Equal(TwoPair))

			Expect(Borrowed_GetHandType(Hand("112J3"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("121J3"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("121J3"))).To(Equal(OnePair))

			Expect(Borrowed_GetHandType(Hand("122JJ"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("122J1"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("122J1"))).To(Equal(TwoPair))

			Expect(Borrowed_GetHandType(Hand("J1234"))).To(Equal(HighCard))
			Expect(Borrowed_GetHandType(Hand("32T3K"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("KK677"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("T55J5"))).To(Equal(ThreeOfAKind))
			Expect(Borrowed_GetHandType(Hand("KTJJT"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("QQQJA"))).To(Equal(ThreeOfAKind))
			Expect(Borrowed_GetHandType(Hand("JJJJJ"))).To(Equal(FiveOfAKind))
			Expect(Borrowed_GetHandType(Hand("96J66"))).To(Equal(ThreeOfAKind))
			Expect(Borrowed_GetHandType(Hand("J6569"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("JKK92"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("234JJ"))).To(Equal(OnePair))

			Expect(Borrowed_GetHandType(Hand("23456"))).To(Equal(HighCard))
			Expect(Borrowed_GetHandType(Hand("32T3K"))).To(Equal(OnePair))
			Expect(Borrowed_GetHandType(Hand("KK677"))).To(Equal(TwoPair))
			Expect(Borrowed_GetHandType(Hand("JJJ23"))).To(Equal(ThreeOfAKind))
			Expect(Borrowed_GetHandType(Hand("KKKQQ"))).To(Equal(FullHouse))
			Expect(Borrowed_GetHandType(Hand("KKKKQ"))).To(Equal(FourOfAKind))
			Expect(Borrowed_GetHandType(Hand("KKKKK"))).To(Equal(FiveOfAKind))
			Expect(Borrowed_GetHandType(Hand("22222"))).To(Equal(FiveOfAKind))

		})

		It("computes highest card", func() {
			Expect(GetHighestHand(Hand("33332"), Hand("2AAAA"))).To(Equal(Hand("33332")))
			Expect(GetHighestHand(Hand("77888"), Hand("77788"))).To(Equal(Hand("77888")))
			Expect(GetHighestHand(Hand("KTJJT"), Hand("KK677"))).To(Equal(Hand("KK677")))
			Expect(GetHighestHand(Hand("KK677"), Hand("KTJJT"))).To(Equal(Hand("KK677")))
			Expect(GetHighestHand(Hand("T55J5"), Hand("QQQJA"))).To(Equal(Hand("QQQJA")))

			Expect(GetHighestHand(Hand("11115"), Hand("11116"))).To(Equal(Hand("11116")))
			Expect(GetHighestHand(Hand("33332"), Hand("2AAAA"))).To(Equal(Hand("33332")))
			Expect(GetHighestHand(Hand("2KQ83"), Hand("2K59J"))).To(Equal(Hand("2KQ83")))
			Expect(GetHighestHand(Hand("T55J5"), Hand("KK677"))).To(Equal(Hand("T55J5")))

			Expect(GetHighestHand(Hand("Q78QQ"), Hand("Q7K77"))).To(Equal(Hand("Q7K77")))
			Expect(GetHighestHand(Hand("J345A"), Hand("32T3K"))).To(Equal(Hand("32T3K")))
			Expect(GetHighestHand(Hand("AAKQJ"), Hand("22234"))).To(Equal(Hand("22234")))

			Expect(GetHighestHand(Hand("32T3K"), Hand("T55J5"))).To(Equal(Hand("T55J5")))
			Expect(GetHighestHand(Hand("32T3K"), Hand("KK677"))).To(Equal(Hand("KK677")))
			Expect(GetHighestHand(Hand("32T3K"), Hand("KTJJT"))).To(Equal(Hand("KTJJT")))
			Expect(GetHighestHand(Hand("32T3K"), Hand("QQQJA"))).To(Equal(Hand("QQQJA")))

			Expect(GetHighestHand(Hand("KK677"), Hand("KTJJT"))).To(Equal(Hand("KK677")))

		})

		When("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt")).To(Equal(6440))
			})
		})
		When("given curated input", func() {
			It("resolves", func() {
				Expect(Star1("curated_input.txt")).To(Equal(6592))
				Expect(Star1("curated_input2.txt")).To(Equal(1343))
				Expect(Star1("curated_input3.txt")).To(Equal(2237))

			})
		})
		When("given input", func() {
			It("resolves", func() {
				result := Star1("star_input.txt")

				Expect(result).NotTo(Equal(249040044))
				Expect(result).NotTo(Equal(249040722))
				Expect(result).NotTo(Equal(249403760))
				Expect(result).NotTo(Equal(249644722))
				Expect(result).NotTo(Equal(250024151))
				Expect(result).To(Equal(249483956))

			})
		})
	})
	Context("Star2", func() {
		When("given sample", func() {
			It("resolves", func() {
				JokerOn = true
				Expect(Star1("star_sample.txt")).To(Equal(5905))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				JokerOn = true
				Expect(Star1("star_input.txt")).To(Equal(252137472))
			})
		})
	})
})
