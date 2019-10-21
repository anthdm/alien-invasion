# Alien invasion
A simulation on how aliens could invade planets and destroy kindoms, written in the Go language.

```
   _____  .__  .__                .___                           .__
  /  _  \ |  | |__| ____   ____   |   | _______  _______    _____|__| ____   ____
 /  /_\  \|  | |  |/ __ \ /    \  |   |/    \  \/ /\__  \  /  ___/  |/  _ \ /    \
/    |    \  |_|  \  ___/|   |  \ |   |   |  \   /  / __ \_\___ \|  (  <_> )   |  \
\____|__  /____/__|\___  >___|  / |___|___|  /\_/  (____  /____  >__|\____/|___|  /
        \/             \/     \/           \/           \/     \/               \/

Starting simulation
> world: data/world.txt
> aliens invading: 3
> cities available: 16
> epoch interval: 3

alien 4039455774    travelled    EAST     thunderbluff            =>    ogrimmar
alien 1879968118    travelled    NORTH    darkshore               =>    winterspring
alien 2949882636    travelled    NORTH    moltencore              =>    thunderbluff
alien 4039455774    travelled    NORTH    ogrimmar                =>    barrens
alien 1879968118    travelled    SOUTH    winterspring            =>    darkshore
alien 2949882636    travelled    WEST     thunderbluff            =>    blackrockmountain
alien 4039455774    travelled    SOUTH    barrens                 =>    ogrimmar
alien 1879968118    travelled    EAST     darkshore               =>    brill
alien 2949882636    travelled    NORTH    blackrockmountain       =>    tarrenmill
alien 4039455774    travelled    SOUTH    ogrimmar                =>    searinggorge
alien 1879968118    travelled    WEST     brill                   =>    darkshore
alien 2949882636    travelled    EAST     tarrenmill              =>    durotar
alien 4039455774    travelled    EAST     searinggorge            =>    winterspring
alien 1879968118    travelled    SOUTH    darkshore               =>    undercity
alien 2949882636    travelled    WEST     durotar                 =>    tarrenmill
alien 4039455774    travelled    SOUTH    winterspring            =>    darkshore
alien 1879968118    travelled    NORTH    undercity               =>    darkshore
alien 2949882636    travelled    SOUTH    tarrenmill              =>    blackrockmountain
alien 4039455774 died gracefully in combat
alien 1879968118 died gracefully in combat
darkshore is destroyed! 1 remaining alien(s)
The simulation is complete, it took 6.003789451s
```

## Installation

```
git clone git@github.com:anthdm/alien-invasion.git
cd alien-invasion
make
```

After running the commands above, the binary program is available as `bin/invasion`

## Tests
It's always a good idea to run tests before running fresh cloned programs.

```
make test
```

## Usage

```
./bin/invasion -aliens 4  -interval 1s -world data/world.txt
```

default values:
- aliens: 3
- interval: 1 second
- world: data/world.txt

## Extra (world generator)
Creating world files by hand is a slow and tidious process. That's why I left a simple and quick script in the `worldgen` folder to generate those world files automatically. 

How to use

```
go run worldgen/main.go > world.txt
```

This will generate 16 cities. You can add more cities by simply adding them to the list. Make sure the `cutoff` parameter is adjusted to fit your needs.
