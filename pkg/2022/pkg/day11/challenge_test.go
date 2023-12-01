package day11

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("challenge", func() {
	When("given sample", func() {
		It("converts data to monkeys", func() {
			monkeys, err := InputFileToMonkeys("sample.txt")
			Expect(err).NotTo(HaveOccurred())
			Expect(len(monkeys)).To(Equal(4))

			Expect(monkeys[0].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(int64(79)), *big.NewInt(int64(98))}))
			Expect(monkeys[0].DivisibleValue).To(Equal(23))
			Expect(monkeys[0].MonkeyTrue).To(Equal(2))
			Expect(monkeys[0].MonkeyFalse).To(Equal(3))

			Expect(monkeys[1].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(int64(54)), *big.NewInt(int64(65)), *big.NewInt(int64(75)), *big.NewInt(int64(74))}))
			Expect(monkeys[1].DivisibleValue).To(Equal(19))
			Expect(monkeys[1].MonkeyTrue).To(Equal(2))
			Expect(monkeys[1].MonkeyFalse).To(Equal(0))

			Expect(monkeys[2].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(int64(79)), *big.NewInt(int64(60)), *big.NewInt(int64(97))}))
			Expect(monkeys[2].DivisibleValue).To(Equal(13))
			Expect(monkeys[2].MonkeyTrue).To(Equal(1))
			Expect(monkeys[2].MonkeyFalse).To(Equal(3))

			Expect(monkeys[3].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(int64(74))}))
			Expect(monkeys[3].DivisibleValue).To(Equal(17))
			Expect(monkeys[3].MonkeyTrue).To(Equal(0))
			Expect(monkeys[3].MonkeyFalse).To(Equal(1))

			Expect(monkeys[0].Operation(big.NewInt(1))).To(Equal(big.NewInt(int64(19))))
			Expect(monkeys[1].Operation(big.NewInt(1))).To(Equal(big.NewInt(int64(7))))
			Expect(monkeys[2].Operation(big.NewInt(1))).To(Equal(big.NewInt(int64(1))))
			Expect(monkeys[3].Operation(big.NewInt(1))).To(Equal(big.NewInt(int64(4))))

			Expect(monkeys[0].Operation(big.NewInt(2))).To(Equal(big.NewInt(int64(38))))
			Expect(monkeys[1].Operation(big.NewInt(2))).To(Equal(big.NewInt(int64(8))))
			Expect(monkeys[2].Operation(big.NewInt(2))).To(Equal(big.NewInt(int64(4))))
			Expect(monkeys[3].Operation(big.NewInt(2))).To(Equal(big.NewInt(int64(5))))
		})

		It("return expected worry levels per monkey after ROUND 1", func() {
			monkeys, err := InputFileToMonkeys("sample.txt")
			err = PlayRound(monkeys, big.NewInt(3))
			Expect(err).NotTo(HaveOccurred())

			Expect(monkeys[0].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(20), *big.NewInt(23), *big.NewInt(27), *big.NewInt(26)}))
			Expect(monkeys[1].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(2080), *big.NewInt(25), *big.NewInt(167), *big.NewInt(207), *big.NewInt(401), *big.NewInt(1046)}))
			Expect(monkeys[2].StartingItemsWorryLevel).To(Equal([]big.Int{}))
			Expect(monkeys[3].StartingItemsWorryLevel).To(Equal([]big.Int{}))
		})

		It("return expected worry levels per monkey after ROUND 20", func() {
			monkeys, err := InputFileToMonkeys("sample.txt")
			for i := 0; i < 20; i++ {
				err = PlayRound(monkeys, big.NewInt(3))
				Expect(err).NotTo(HaveOccurred())
			}

			Expect(monkeys[0].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(10), *big.NewInt(12), *big.NewInt(14), *big.NewInt(26), *big.NewInt(34)}))
			Expect(monkeys[1].StartingItemsWorryLevel).To(Equal([]big.Int{*big.NewInt(245), *big.NewInt(93), *big.NewInt(53), *big.NewInt(199), *big.NewInt(115)}))
			Expect(monkeys[2].StartingItemsWorryLevel).To(Equal([]big.Int{}))
			Expect(monkeys[3].StartingItemsWorryLevel).To(Equal([]big.Int{}))

			Expect(monkeys[0].ItemsInspected).To(Equal(101))
			Expect(monkeys[1].ItemsInspected).To(Equal(95))
			Expect(monkeys[2].ItemsInspected).To(Equal(7))
			Expect(monkeys[3].ItemsInspected).To(Equal(105))
		})

		DescribeTable("return inspected items after ROUND X with NO WORRY LEVEL DIVISOR", func(numRounds int, expectedTotalItems []int) {
			monkeys, _ := InputFileToMonkeys("sample.txt")
			divisor := 96577
			//for divisorAttempt := 5; true; ;{
			for i := 0; i < numRounds; i++ {
				//fmt.Printf("divisorAttempt: %d | round:%d \n", divisor, i)
				_ = PlayRound(monkeys, big.NewInt(int64(divisor)))
				//Expect(err).NotTo(HaveOccurred())
			}

			//if monkeys[0].ItemsInspected == expectedTotalItems[0] &&
			//	monkeys[1].ItemsInspected == expectedTotalItems[1] &&
			//	monkeys[2].ItemsInspected == expectedTotalItems[2] &&
			//	monkeys[3].ItemsInspected == expectedTotalItems[3] {
			//	//divisor = divisor
			//	fmt.Printf("divisor is: %d", divisor)
			//	//break
			//}

			Expect(monkeys[0].ItemsInspected).To(Equal(expectedTotalItems[0]))
			Expect(monkeys[1].ItemsInspected).To(Equal(expectedTotalItems[1]))
			Expect(monkeys[2].ItemsInspected).To(Equal(expectedTotalItems[2]))
			Expect(monkeys[3].ItemsInspected).To(Equal(expectedTotalItems[3]))
		},
			Entry("_", 1, []int{2, 4, 3, 6}),
			Entry("_", 20, []int{99, 97, 8, 103}),
			Entry("_", 1000, []int{5204, 4792, 199, 5192}),
			Entry("_", 2000, []int{10419, 9577, 392, 10391}),
			Entry("_", 10000, []int{52166, 47830, 1938, 52013}),
		)
	})

	When("given sample file", func() {
		It("resolve the multiplication of the TWO monkeys with MOST ItemsInspected", func() {
			monkeys, err := InputFileToMonkeys("sample.txt")
			for i := 0; i < 20; i++ {
				err = PlayRound(monkeys, big.NewInt(3))
				Expect(err).NotTo(HaveOccurred())
			}

			Expect(monkeys[0].StartingItemsWorryLevel).To(Equal([]int{10, 12, 14, 26, 34}))
			Expect(monkeys[1].StartingItemsWorryLevel).To(Equal([]int{245, 93, 53, 199, 115}))
			Expect(monkeys[2].StartingItemsWorryLevel).To(Equal([]int{}))
			Expect(monkeys[3].StartingItemsWorryLevel).To(Equal([]int{}))

			Expect(monkeys[0].ItemsInspected).To(Equal(101))
			Expect(monkeys[1].ItemsInspected).To(Equal(95))
			Expect(monkeys[2].ItemsInspected).To(Equal(7))
			Expect(monkeys[3].ItemsInspected).To(Equal(105))

			Expect(101 * 105).To(Equal(10605))
		})
	})

	When("star 1 - given input file", func() {
		It("resolve the multiplication of the TWO monkeys with MOST ItemsInspected", func() {
			monkeys, err := InputFileToMonkeys("input.txt")
			for i := 0; i < 20; i++ {
				err = PlayRound(monkeys, big.NewInt(3))
				Expect(err).NotTo(HaveOccurred())
			}

			Expect(monkeys[0].StartingItemsWorryLevel).To(Equal([]int{5, 5, 5, 3, 3, 3, 3, 3, 3, 19}))
			Expect(monkeys[1].StartingItemsWorryLevel).To(Equal([]int{64, 7, 7, 7, 7, 7, 7, 9, 5, 270023884279, 3}))
			Expect(monkeys[2].StartingItemsWorryLevel).To(Equal([]int{}))
			Expect(monkeys[3].StartingItemsWorryLevel).To(Equal([]int{8, 8, 8, 8, 8, 535519, 5, 5, 3, 3, 3, 3, 3, 3, 3}))

			Expect(monkeys[0].ItemsInspected).To(Equal(163))
			Expect(monkeys[1].ItemsInspected).To(Equal(258))
			Expect(monkeys[2].ItemsInspected).To(Equal(51))
			Expect(monkeys[3].ItemsInspected).To(Equal(247))
			Expect(monkeys[4].ItemsInspected).To(Equal(271))
			Expect(monkeys[5].ItemsInspected).To(Equal(22))
			Expect(monkeys[6].ItemsInspected).To(Equal(105))
			Expect(monkeys[7].ItemsInspected).To(Equal(229))

			result := 271 * 258
			Expect(result).ToNot(Equal(63726))
			Expect(result).To(Equal(69918))
		})
	})

	FIt("return inspected items after ROUND X with NO WORRY LEVEL DIVISOR", func() {
		monkeys, _ := InputFileToMonkeys("input.txt")
		//for divisorAttempt := 5; true; ;{
		for i := 0; i < 10000; i++ {
			//fmt.Printf("divisorAttempt: %d | round:%d \n", divisor, i)
			err := PlayRound(monkeys, big.NewInt(9699690))
			Expect(err).NotTo(HaveOccurred())
		}

		//if monkeys[0].ItemsInspected == expectedTotalItems[0] &&
		//	monkeys[1].ItemsInspected == expectedTotalItems[1] &&
		//	monkeys[2].ItemsInspected == expectedTotalItems[2] &&
		//	monkeys[3].ItemsInspected == expectedTotalItems[3] {
		//	//divisor = divisor
		//	fmt.Printf("divisor is: %d", divisor)
		//	//break
		//}

		for _, monkey := range monkeys {
			fmt.Println(monkey.ItemsInspected)
		}

		Expect(monkeys[0].ItemsInspected).To(Equal(80160))
		Expect(monkeys[1].ItemsInspected).To(Equal(139887))
		Expect(monkeys[2].ItemsInspected).To(Equal(99585))
		Expect(monkeys[3].ItemsInspected).To(Equal(139923))
		Expect(monkeys[4].ItemsInspected).To(Equal(119975))
		Expect(monkeys[5].ItemsInspected).To(Equal(19972))
		Expect(monkeys[6].ItemsInspected).To(Equal(99605))
		Expect(monkeys[7].ItemsInspected).To(Equal(40344))

		monkeyChaos := 139923 * 139887
		Expect(monkeyChaos).NotTo(Equal(2567194800))
		Expect(monkeyChaos).NotTo(Equal(2637102168))
		Expect(monkeyChaos).NotTo(Equal(16787261925))
		Expect(monkeyChaos).To(Equal(19573408701))

	})
})
