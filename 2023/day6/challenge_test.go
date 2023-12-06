package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {
		It("returns options per lastingTime", func() {
			Expect(GetDistanceOptions(7)).To(Equal(map[int]int{
				0: 0,
				1: 6,
				2: 10,
				3: 12,
				4: 12,
				5: 10,
				6: 6,
				7: 0,
			}))
		})

		It("returns options that beat maxDistance", func() {
			Expect(GetBeatingOptions(map[int]int{
				0: 0,
				1: 7,
				2: 10,
				3: 12,
				4: 12,
				5: 10,
				6: 6,
				7: 0,
			}, 9)).To(Equal(4))
		})

		It("returns options that beat maxDistance", func() {
			Expect(GetBeatingOptions(GetDistanceOptions(15), 40)).To(Equal(8))
			Expect(GetBeatingOptions(GetDistanceOptions(30), 200)).To(Equal(9))

		})

		When("given sample", func() {
			It("resolves", func() {
				race1 := GetBeatingOptions(GetDistanceOptions(7), 9)
				race2 := GetBeatingOptions(GetDistanceOptions(15), 40)
				race3 := GetBeatingOptions(GetDistanceOptions(30), 200)
				Expect(race1 * race2 * race3).To(Equal(288))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				race1 := GetBeatingOptions(GetDistanceOptions(58), 478)
				race2 := GetBeatingOptions(GetDistanceOptions(99), 2232)
				race3 := GetBeatingOptions(GetDistanceOptions(64), 1019)
				race4 := GetBeatingOptions(GetDistanceOptions(69), 1071)
				Expect(race1 * race2 * race3 * race4).To(Equal(128700))
			})
		})
	})
	Context("Star2", func() {
		When("given sample", func() {
			It("resolves", func() {
				race1 := GetBeatingOptions(GetDistanceOptions(71530), 940200)
				Expect(race1).To(Equal(71503))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				race1 := GetBeatingOptions(GetDistanceOptions(58996469), 478223210191071)
				Expect(race1).To(Equal(39594072))
			})
		})
	})
})
