package day5

import (
	"fmt"
	"math"
	"strings"

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


// useMap applies the mapping rules to a seed
func useMap(im islandMap, val int64) int64 {
	for _, m := range im {
		ss := m[1]
		se := ss + m[2]
		if ss <= val && val < se {
			ds := m[0]
			return ds + (val - ss)
		}
	}
	return val
}

func findMinimum(m maps, start, length, step int64, ims []islandMap) int64 {
	if step == int64(len(ims)) { // Terminal condition: reached the 'location' category.
		return start
	}

	im := ims[step]
	minVal := int64(math.MaxInt64)

	for _, mapping := range im {
		sourceStart, sourceEnd := mapping[1], mapping[1]+mapping[2]
		destStart := mapping[0]
		if sourceStart <= start && start < sourceEnd {
			// Current seed falls within the mapping range.
			mappedStart := destStart + (start - sourceStart)
			mappedLength := min(sourceEnd, start+length) - start
			mappedVal := findMinimum(m, mappedStart, mappedLength, step+1, ims)
			if mappedVal < minVal {
				minVal = mappedVal
			}
			if mappedLength < length {
				// Handle the unmapped part of the range.
				unmappedVal := findMinimum(m, start+mappedLength, length-mappedLength, step, ims)
				if unmappedVal < minVal {
					minVal = unmappedVal
				}
			}
			return minVal // Since ranges do not overlap, we can return immediately.
		}
	}

	// If no mapping range includes the current seed, process the next category.
	return findMinimum(m, start, length, step+1, ims)
}

// LowestAnyLocationNumber finds the lowest location number 
// that corresponds to any of the initial seed numbers.
func LowestAnyLocationNumber(m maps) int64 {
	ims := []islandMap{
		m.sts, m.stf, m.ftw, m.wtl, m.ltt, m.tth, m.htl,
	}

	minVal := int64(math.MaxInt64)
	for i := 0; i < len(m.seeds); i += 2 {
		start, length := m.seeds[i], m.seeds[i+1]
		val := findMinimum(m, start, length, 0, ims)
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}


func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}