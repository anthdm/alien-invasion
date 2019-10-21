package main

import "fmt"

var cityNames = []string{
	"durotar",
	"barrens",
	"undercity",
	"tarrenmill",
	"thunderbluff",
	"ogrimmar",
	"bootybay",
	"blackrockmountain",
	"moltencore",
	"searinggorge",
	"winterspring",
	"arathihighlands",
	"stranglethornvale",
	"silverspineforest",
	"darkshore",
	"brill",
	"redridgemountains",
	"decolace",
	"feralas",
}

const cutoff = 4

func main() {
	cities := make([][]string, len(cityNames)/cutoff)
	for i := 0; i < len(cities); i++ {
		cities[i] = cityNames[i*cutoff : i*cutoff+cutoff]
	}

	var (
		north string
		east  string
		south string
		west  string
	)
	for ri, row := range cities {
		for ci, city := range row {
			if ri == 0 {
				north = cities[len(cities)-1][ci]
			} else {
				north = cities[ri-1][ci]
			}
			if ri == len(cities)-1 {
				south = cities[0][ci]
			} else {
				south = cities[ri+1][ci]
			}
			if ci == len(row)-1 {
				east = cities[ri][0]
			} else {
				east = cities[ri][ci+1]
			}
			if ci == 0 {
				west = row[len(row)-1]
			} else {
				west = row[ci-1]
			}
			fmt.Printf("%s north=%s east=%s south=%s west=%s\n",
				city,
				north,
				east,
				south,
				west)
		}
	}
}
