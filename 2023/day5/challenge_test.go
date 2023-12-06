package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {

		It("gets almanac info", func() {

			almanac, err := ReadAlmanac("star_sample.txt")
			Expect(err).NotTo(HaveOccurred())
			Expect(len(almanac)).To(Equal(7))

			Expect(GetSeedAlamanacPerState([]almanacPerState{{
				source:      50,
				destination: 52,
				drange:      48,
			}}, 79)).To(Equal(81))

			Expect(GetSeedAlmanac(almanac, 79)).To(Equal(SeedAlmanac{
				seed: 79,
				stateMap: map[string]int{
					seedToSoil:            81,
					soilToFertilizer:      81,
					fertilizerToWater:     81,
					waterToLight:          74,
					lightToTemperature:    78,
					temperatureToHumidity: 78,
					humidityToLocation:    82,
				},
			}))

			Expect(GetSeedAlmanac(almanac, 82)).To(Equal(SeedAlmanac{
				seed: 82,
				stateMap: map[string]int{
					seedToSoil:            84,
					soilToFertilizer:      84,
					fertilizerToWater:     84,
					waterToLight:          77,
					lightToTemperature:    45,
					temperatureToHumidity: 46,
					humidityToLocation:    46,
				},
			}))
		})
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt", []int{79, 14, 55, 13})).To(Equal(35))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				Expect(Star1("star_input.txt", []int{
					1187290020, 247767461, 40283135, 64738286, 2044483296, 66221787, 1777809491, 103070898, 108732160, 261552692, 3810626561, 257826205, 3045614911, 65672948, 744199732, 300163578, 3438684365, 82800966, 2808575117, 229295075,
				})).To(Equal(282277027))
			})
		})
	})

	Context("Star2", func() {
		When("given sample", func() {
			It("resolves", func() {
				Expect(Star2("star_sample.txt", []int{79, 14, 55, 13})).To(Equal(46))
			})
		})
		When("given input", func() {
			It("resolves", func() {
				result := Star2("star_input.txt", []int{
					1187290020, 247767461, 40283135, 64738286, 2044483296, 66221787, 1777809491, 103070898, 108732160, 261552692, 3810626561, 257826205, 3045614911, 65672948, 744199732, 300163578, 3438684365, 82800966, 2808575117, 229295075,
				})
				Expect(result).NotTo(Equal(11554134))
				Expect(result).To(Equal(11554135))
			})
		})
	})
})
