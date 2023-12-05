package day5

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/burtenshaw/advent/src/utils"
)

type maps struct {
	seeds []int64
	sts   islandMap
	stf   islandMap
	ftw   islandMap
	wtl   islandMap
	ltt   islandMap
	tth   islandMap
	htl   islandMap
}

type islandMap [][3]int64

func Run(inputPath string) {
	lines := utils.ReaderSplit(inputPath)
	m := constructMaps(lines)
	fmt.Printf("Part One: %d\n", LowestLocationNumber(m))
	fmt.Printf("Part Two: %d\n", LowestAnyLocationNumber(m))
}

func constructMaps(lines []string) maps {
	return maps{
		seeds: utils.MustStringToInt64Slice(strings.Split(lines[0], ":")[1]),
		sts:   constructMap(lines, "seed-to-soil"),
		stf:   constructMap(lines, "soil-to-fertilizer"),
		ftw:   constructMap(lines, "fertilizer-to-water"),
		wtl:   constructMap(lines, "water-to-light"),
		ltt:   constructMap(lines, "light-to-temperature"),
		tth:   constructMap(lines, "temperature-to-humidity"),
		htl:   constructMap(lines, "humidity-to-location"),
	}
}

func constructMap(lines []string, mapName string) islandMap {
	im := islandMap{}
	idx := findIndex(lines, mapName) + 1
	line := lines[idx]
	for line != "" {
		sn := strings.Fields(line)
		n := [3]int64{}
		for i := 0; i < 3; i++ {
			n[i] = utils.MustParseInt64(sn[i])
		}
		im = append(im, n)

		idx++
		if idx < len(lines) {
			line = lines[idx]
		} else {
			break
		}
	}
	return im
}

func findIndex(lines []string, s string) int {
	for i, l := range lines {
		if strings.HasPrefix(l, s) {
			return i
		}
	}
	return -1
}

func useMap(im islandMap, val int64) int64 {
	for _, m := range im {
		ss := m[1]
		se := ss + m[2]

		if ss <= val && val < se {
			ds := m[0]
			dist := val - ss
			return ds + dist
		}
	}
	return val
}

func LowestLocationNumber(m maps) int64 {
	ims := []islandMap{
		m.sts,
		m.stf,
		m.ftw,
		m.wtl,
		m.ltt,
		m.tth,
		m.htl,
	}
	locs := []int64{}
	for _, val := range m.seeds {
		for _, im := range ims {
			val = useMap(im, val)
		}
		locs = append(locs, val)
	}

	minVal := locs[0]
	for _, l := range locs {
		if l < minVal {
			minVal = l
		}
	}
	return minVal
}

func LowestAnyLocationNumber(m maps) int64 {
	ims := []islandMap{
		m.sts,
		m.stf,
		m.ftw,
		m.wtl,
		m.ltt,
		m.tth,
		m.htl,
	}

	// Determine an appropriate number of workers
	workerCount := runtime.NumCPU()
	workChan := make(chan int, workerCount)
	results := make(chan int64)
	var wg sync.WaitGroup

	// Create workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range workChan {
				startTime := time.Now()
				ml := int64(math.MaxInt64)
				for seed := m.seeds[i]; seed < m.seeds[i]+m.seeds[i+1]; seed++ {
					val := seed
					for _, im := range ims {
						val = useMap(im, val)
					}
					if val < ml {
						ml = val
					}
				}
				endTime := time.Now()
				fmt.Printf("Finished goroutine for range starting at %d, %f seconds\n", i, endTime.Sub(startTime).Seconds())
				results <- ml
			}
		}()
	}

	// spread it out
	go func() {
		for i := 0; i < len(m.seeds); i += 2 {
			workChan <- i
		}
		close(workChan)
	}()

	// pick up results
	go func() {
		wg.Wait()
		close(results)
	}()

	minLoc := int64(math.MaxInt64)
	for ans := range results {
		if ans < minLoc {
			minLoc = ans
		}
	}

	return minLoc
}
