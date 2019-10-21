package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCitiesFromFile(t *testing.T) {
	cities, err := buildCitiesFromFile("../data/world.txt")
	assert.Nil(t, err)
	assert.Equal(t, len(cities), 16)

	for _, c := range cities {
		assert.False(t, c.isDestroyed)
	}

	cities, err = buildCitiesFromFile("data/world.txt")
	assert.NotNil(t, err)
	assert.Equal(t, len(cities), 0)
}

func TestAlienMove(t *testing.T) {
	cities, err := buildCitiesFromFile("../data/world.txt")
	assert.Nil(t, err)

	var (
		state       = newSimulationState(cities)
		currentCity = randomCity(cities)
		a           = newAlien(currentCity)
	)

	cities[currentCity.adjacent[south]].isDestroyed = true
	cities[currentCity.adjacent[north]].isDestroyed = true
	cities[currentCity.adjacent[east]].isDestroyed = true

	// only option for the alien is west
	for i := 0; i < 100; i++ {
		dir, city := a.move(state)
		assert.Equal(t, west, dir)
		assert.NotEqual(t, a.currentCity.name, city.name)
	}
}

func TestCityInvaded(t *testing.T) {
	cities, err := buildCitiesFromFile("../data/world.txt")
	assert.Nil(t, err)

	// pick a random city and let both aliens invade it.
	var (
		city = randomCity(cities)
		a1   = newAlien(city)
		a2   = newAlien(city)
		sim  = &simulator{
			state: newSimulationState(cities, a1, a2),
		}
	)

	city.addAlien(a1)
	city.addAlien(a2)

	assert.False(t, city.isDestroyed)
	assert.Equal(t, 2, sim.remainingAliens())
	assert.True(t, true, city.isUnderAttack())

	// update the simulation, city should be destroyed with 0 aliens alive.
	sim.update()

	assert.True(t, city.isDestroyed)
	assert.Equal(t, 0, sim.remainingAliens())
}

func TestCityDestroy(t *testing.T) {
	cities, err := buildCitiesFromFile("../data/world.txt")
	assert.Nil(t, err)

	city := randomCity(cities)
	city.addAlien(newAlien(city))
	city.addAlien(newAlien(city))
	city.addAlien(newAlien(city))

	assert.False(t, city.isDestroyed)
	assert.Equal(t, 3, len(city.aliens))
	city.destroy()

	assert.True(t, city.isDestroyed)

	for _, a := range city.aliens {
		assert.False(t, a.isAlive)
	}
}

func newSimulationState(cities map[string]*city, aliens ...*alien) *simulationState {
	return &simulationState{
		aliens: aliens,
		cities: cities,
	}
}
