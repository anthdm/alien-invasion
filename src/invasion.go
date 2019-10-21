package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var errInvalidFileFormat = errors.New("invalid file format")

type direction int

const (
	north direction = iota
	east
	south
	west
)

// String implements the Stringer interface.
func (d direction) String() string {
	switch d {
	case north:
		return "NORTH"
	case east:
		return "EAST"
	case south:
		return "SOUTH"
	case west:
		return "WEST"
	default:
		return "UNKNOWN"
	}
}

func newDirection(s string) direction {
	switch s {
	case "north":
		return north
	case "east":
		return east
	case "south":
		return south
	case "west":
		return west
	default:
		panic("invalid string direction given")
	}
}

type alien struct {
	// Random identifier
	id uint32

	// Whether this alien is dead or alive, dead aliens are not in the simulation
	// update loop.
	isAlive bool

	// The current city this alien is located/invading.
	currentCity *city
}

func newAlien(c *city) *alien {
	return &alien{
		id:          rand.Uint32(),
		currentCity: c,
		isAlive:     true,
	}
}

func (a *alien) die() {
	a.isAlive = false
	fmt.Printf("alien %d died gracefully in combat\n", a.id)
}

func (a *alien) move(s *simulationState) (direction, *city) {
	dir := randomDirection()
	cityName := a.currentCity.adjacent[dir]
	city := s.cities[cityName]

	if city.isDestroyed {
		fmt.Printf("alien %d need to rethink his direction, %s is destroyed\n", a.id, city.name)
		// IMPROVEMENT: We can make this faster by excluding the current
		// direction when picking a random one.
		return a.move(s)
	}

	return dir, city
}

func (a *alien) update(s *simulationState) {
	var (
		currentCity = a.currentCity
		dir, city   = a.move(s)
	)

	a.currentCity = city
	city.addAlien(a)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "alien %d\ttravelled\t%s\t%s\t=>\t%s\t\n",
		a.id,
		fillSpace(dir.String(), 5),
		fillSpace(currentCity.name, 20),
		city.name)
	w.Flush()
}

// quick hack to make the internal tabwriter write equal width columns.
func fillSpace(s string, n int) string {
	if len(s) < n {
		return fmt.Sprintf("%s%s", s, strings.Repeat(" ", n-len(s)))
	}
	return s
}

type city struct {
	name        string
	isDestroyed bool

	// aliens that are currently invading this city
	aliens []*alien

	// map of adjacent cities
	adjacent map[direction]string
}

func (c *city) addAlien(a *alien) {
	c.aliens = append(c.aliens, a)
}

func (c *city) isUnderAttack() bool {
	return len(c.aliens) >= 2
}

// When a city is destroyed, all current invading aliens in that city die during
// that invasion.
func (c *city) destroy() {
	for _, a := range c.aliens {
		a.die()
	}
	c.isDestroyed = true
}

type simulationState struct {
	aliens []*alien
	cities map[string]*city
}

type simulator struct {
	epochInterval time.Duration
	worldFile     string
	state         *simulationState
	quitch        chan struct{}
}

func newSimulator(worldFile string, interval time.Duration, nAliens int) (*simulator, error) {
	cities, err := buildCitiesFromFile(worldFile)
	if err != nil {
		return nil, err
	}

	s := &simulator{
		worldFile:     worldFile,
		epochInterval: interval,
		state: &simulationState{
			cities: cities,
			aliens: make([]*alien, nAliens),
		},
		quitch: make(chan struct{}),
	}

	// Create n aliens and place them in a random city.
	for i := 0; i < nAliens; i++ {
		s.state.aliens[i] = newAlien(randomCity(cities))
	}

	return s, nil
}

func (s *simulator) start() {
	var (
		start    = time.Now()
		interval = time.NewTicker(s.epochInterval)
	)

	fmt.Println("Starting simulation")
	fmt.Printf("> world: %s\n", s.worldFile)
	fmt.Printf("> aliens invading: %d\n", len(s.state.aliens))
	fmt.Printf("> cities available: %d\n", len(s.state.cities))
	fmt.Printf("> epoch interval: %d\n", len(s.state.aliens))
	fmt.Println()

loop:
	for {
		select {
		case <-s.quitch:
			break loop
		case <-interval.C:
			s.update()
		}
	}

	fmt.Printf("The simulation is complete, it took %s\n", time.Since(start))
}

func (s *simulator) remainingAliens() int {
	remAliens := 0
	for _, a := range s.state.aliens {
		if a.isAlive {
			remAliens++
		}
	}
	return remAliens
}

func (s *simulator) update() {
	// update all aliens in the simulation
	for _, a := range s.state.aliens {
		if a.isAlive {
			a.update(s.state)
		}
	}

	// check if cities are under attack, destroy city and its invading aliens if so.
	for _, c := range s.state.cities {
		if c.isUnderAttack() {
			c.destroy()
			fmt.Printf("%s is destroyed! %d remaining alien(s)\n", c.name, s.remainingAliens())
		}
	}

	if s.isTerminated() {
		go func() { s.quitch <- struct{}{} }()
	}

	// reset all aliens inside cities before the next simulation.
	for _, c := range s.state.cities {
		c.aliens = []*alien{}
	}
}

func (s *simulator) isTerminated() bool {
	aliensAlive := 0
	for _, a := range s.state.aliens {
		if a.isAlive {
			aliensAlive++
		}
	}
	return aliensAlive < 2
}

func randomDirection() direction {
	return direction(rand.Intn(4))
}

func randomCity(cities map[string]*city) *city {
	citySlice := make([]*city, len(cities))
	i := 0
	for _, v := range cities {
		citySlice[i] = v
		i++
	}
	return citySlice[rand.Intn(len(citySlice))]
}

func buildCitiesFromFile(src string) (map[string]*city, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cities := make(map[string]*city)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			l     = scanner.Text()
			parts = strings.Split(l, " ")
		)
		if len(parts) != 5 {
			return nil, errInvalidFileFormat
		}

		c := &city{
			name:     parts[0],
			adjacent: make(map[direction]string),
			aliens:   []*alien{},
		}

		// parse directions
		for _, v := range parts[1:] {
			parts := strings.Split(v, "=")
			if len(parts) != 2 {
				return nil, errInvalidFileFormat
			}
			direction := newDirection(parts[0])
			cityname := parts[1]
			c.adjacent[direction] = cityname
		}

		cities[c.name] = c
	}

	return cities, nil
}
