// Inspired by:
// https://www.instagram.com/p/BZkF7Lth8kf/?hl=en&taken-by=mattsurelee

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const INSULTS = `
atomic
steamy
rusty
witless
lumpy
shitty
moist
chunky
lousy
bulbous
trashy
dumbass
nerdy
dotarded
crusty
brainless

knob
bum
turd
prick
bulge
ass
chut
shit
rod
chode
fuck
weiner
jizz
panty
cock
dong

vacuum
general
gremlin
pixie
spasm
fiend
fungus
tunnel
corporal
raider
demon
buccaneer
tyrant
juggler
magician
fiddle
`

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	columns := strings.Split(INSULTS, "\n\n")
	for _, column := range columns {
		words := strings.Split(strings.TrimSpace(column), "\n")
		word := words[rand.Intn(len(words))]
		fmt.Print(word)
		fmt.Print(" ")
	}
	fmt.Println()
}
