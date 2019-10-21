package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var (
		nAliens       int
		epochInterval time.Duration
		worldFile     string
	)

	flag.IntVar(&nAliens, "aliens", 3, "")
	flag.DurationVar(&epochInterval, "interval", 1000*time.Millisecond, "")
	flag.StringVar(&worldFile, "world", "data/world.txt", "")
	flag.Parse()

	if nAliens < 2 {
		log.Fatalf("The simulation needs at least 2 invading aliens, %d given", nAliens)
	}

	s, err := newSimulator(worldFile, epochInterval, nAliens)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(asciBanner)

	s.start()
}

var asciBanner = `
   _____  .__  .__                .___                           .__               
  /  _  \ |  | |__| ____   ____   |   | _______  _______    _____|__| ____   ____  
 /  /_\  \|  | |  |/ __ \ /    \  |   |/    \  \/ /\__  \  /  ___/  |/  _ \ /    \ 
/    |    \  |_|  \  ___/|   |  \ |   |   |  \   /  / __ \_\___ \|  (  <_> )   |  \
\____|__  /____/__|\___  >___|  / |___|___|  /\_/  (____  /____  >__|\____/|___|  /
		\/             \/     \/           \/           \/     \/               \/ 
`
