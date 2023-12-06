package challenge

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const seedToSoil = "seed-to-soil map:"
const soilToFertilizer = "soil-to-fertilizer map:"
const fertilizerToWater = "fertilizer-to-water map:"
const waterToLight = "water-to-light map:"
const lightToTemperature = "light-to-temperature map:"
const temperatureToHumidity = "temperature-to-humidity map:"
const humidityToLocation = "humidity-to-location map:"

type SeedAlmanac struct {
	seed     int
	stateMap map[string]int
}

type almanacPerState struct {
	source      int
	destination int
	drange      int
}

func Star2(inputFileName string, seeds []int) int {

	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(runtime.NumCPU())

	_, err := ReadAlmanac(inputFileName)
	if err != nil {
		return 0
	}

	almanac, err := ReadAlmanac(inputFileName)
	if err != nil {
		return 0
	}

	var m sync.Mutex
	var lowestLocation = 999999999999999999

	//var processedSeeds []int
	var wgPerSeedBlock sync.WaitGroup
	//var wgPerSeed sync.WaitGroup

	var seedNumber = 0
	i := 0
	for i < len(seeds) {

		wgPerSeedBlock.Add(1)

		go func() {
			defer wgPerSeedBlock.Done()
			sourceSeed := seeds[i]
			drange := seeds[i+1]
			i = i + 2
			// fmt.Printf("processing seed %d with drange %d\n", sourceSeed, drange)

			for seed := sourceSeed; seed < sourceSeed+drange; seed++ {
				//wgPerSeed.Add(1)
				//if slices.Contains(processedSeeds, seed) {
				//	continue
				//}
				//processedSeeds = append(processedSeeds, seed)

				//go func() {
				//	defer wgPerSeed.Done()
				seedAlmanac := GetSeedAlmanac(almanac, seed)
				// fmt.Printf("%v\n", seedAlmanac)

				if seedAlmanac.stateMap[humidityToLocation] < lowestLocation {
					fmt.Printf("new lowest location %d for seed %d\n", seedAlmanac.stateMap[humidityToLocation], seed)
					m.Lock()
					lowestLocation = seedAlmanac.stateMap[humidityToLocation]
					m.Unlock()
				}
				m.Lock()
				seedNumber++
				m.Unlock()
				//}()
			}
			//time.Sleep(10 * time.Millisecond)
			//wgPerSeed.Wait()
		}()
		//time.Sleep(10 * time.Millisecond)
		wgPerSeedBlock.Wait()
	}

	// fmt.Printf("processed %d seeds\n", seedNumber)
	return lowestLocation
}

func Star1(inputFileName string, seeds []int) int {

	_, err := ReadAlmanac(inputFileName)
	if err != nil {
		return 0
	}

	almanac, err := ReadAlmanac(inputFileName)
	if err != nil {
		return 0
	}

	var lowestLocation = 999999999999999999
	for _, seed := range seeds {
		seedAlmanac := GetSeedAlmanac(almanac, seed)
		if seedAlmanac.stateMap[humidityToLocation] < lowestLocation {
			lowestLocation = seedAlmanac.stateMap[humidityToLocation]
		}
	}

	return lowestLocation
}

func ReadAlmanac(inputFileName string) (map[string][]almanacPerState, error) {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	almanacStates := make(map[string][]almanacPerState)

	var currentMap string
	for _, dataLine := range strings.Split(string(data), "\n") {
		// fmt.Printf("%s\n", dataLine)
		if strings.HasPrefix(dataLine, "seeds: ") {
			continue
		}
		if strings.TrimSpace(dataLine) == "" {
			continue
		}
		switch dataLine {
		case seedToSoil:
			currentMap = seedToSoil
			break
		case soilToFertilizer:
			currentMap = soilToFertilizer
			break
		case fertilizerToWater:
			currentMap = fertilizerToWater
			break
		case waterToLight:
			currentMap = waterToLight
			break
		case lightToTemperature:
			currentMap = lightToTemperature
			break
		case temperatureToHumidity:
			currentMap = temperatureToHumidity
			break
		case humidityToLocation:
			currentMap = humidityToLocation
			break
		default:
			SetMapForAlmanac(almanacStates, currentMap, dataLine)
		}
	}
	return almanacStates, nil
}

func GetSeedAlmanac(almanac map[string][]almanacPerState, seed int) SeedAlmanac {

	seedAlmanac := SeedAlmanac{seed: seed, stateMap: make(map[string]int)}

	destination := seed
	for _, state := range []string{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation} {
		destination = GetSeedAlamanacPerState(almanac[state], destination)
		seedAlmanac.stateMap[state] = destination
	}
	return seedAlmanac
}

func GetSeedAlamanacPerState(stateAlmanac []almanacPerState, seed int) int {
	destination := seed
	for _, state := range stateAlmanac {
		if seed >= state.source && seed < state.source+state.drange {
			destination = seed - state.source + state.destination
		}
	}
	return destination
}

func SetMapForAlmanac(almanac map[string][]almanacPerState, currentMap string, dataLine string) {

	mapConversion := strings.Split(dataLine, " ")
	destination, err := strconv.Atoi(mapConversion[0])
	if err != nil {
		return
	}
	source, err := strconv.Atoi(mapConversion[1])
	if err != nil {
		return
	}
	drange, err := strconv.Atoi(mapConversion[2])
	if err != nil {
		return
	}

	almanac[currentMap] = append(almanac[currentMap], almanacPerState{
		destination: destination,
		source:      source,
		drange:      drange,
	})
}
