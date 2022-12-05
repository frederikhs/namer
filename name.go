package namer

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"strings"
)

var (
	adjectives = []string{"Able", "Angry", "Bad", "Best", "Better", "Big", "Busy", "Clear", "Different", "Early", "Easy", "Excellent", "Free", "Full", "Good", "Great", "Hard", "High", "Huge", "Important", "International", "Late", "Little", "Local", "Long", "Loud", "Low", "Major", "National", "New", "Old", "Only", "Open", "Other", "Personal", "Political", "Possible", "Public", "Real", "Recent", "Right", "Small", "Social", "Special", "Strong", "Sure", "True", "Weird", "Whole", "Young"}
	animals    = []string{"Alpaca", "Anemone", "Ant", "Bee", "Bee", "Beetle", "Butterfly", "Caiman", "Calf", "Camel", "Canary", "Cat", "Caterpillar", "Cattle", "Centipede", "Chick", "Chicken", "Clam", "Cobra", "Cockroach", "Cocoon", "Collared Lizard", "Copperhead", "Coral", "Coral Snake", "Corn Snake", "Cottonmouth", "Cow", "Crab", "Cricket", "Crocodile", "Cuttlefish", "Dog", "Donkey", "Dove", "Dragonfly", "Duck", "Ferret", "Flea", "Fly", "Foal", "GaboonViper", "GarterSnake", "Gavial", "Gecko", "GilaMonster", "GlassLizard", "Goat", "Goldfish", "Goose", "GopherSnake", "Grasshopper", "GreenIguana", "GroundSkink", "Hen", "Horse", "Jellyfish", "Kid", "Ladybug", "Lama", "Lamb", "Lobster", "Louse", "Mamba", "MapTurtle", "Millipede", "Mink", "Monitor Lizard", "Mosquito", "Moth", "Mountain Viper", "Mouse", "Mud Snake", "Mud Turtle", "Mussel", "Octopus", "Ox", "Oyster", "Pig", "Pigeon", "PrayingMantis", "Rabbit", "Rooster", "Scorpion", "SeaUrchin", "Sheep", "Silkmoth", "Slug", "Snail", "Spider", "Squid", "Starfish", "Taipan", "Timber Rattler", "Tortoise", "Turkey", "Turtle", "Wasp", "WaterMoccasin", "Water Turtle", "Whipsnake", "Worm", "Yak", "Zebu"}
	caser      = cases.Title(language.Und)
)

func getRandFromString(s string) *rand.Rand {
	// naive hashing of string to int
	var hash int32
	for _, r := range s {
		hash += r
	}

	return rand.New(rand.NewSource(int64(hash)))
}

func getParts(seed string) (string, string) {
	r := getRandFromString(seed)

	adjective := adjectives[r.Intn(len(adjectives))]
	animal := animals[r.Intn(len(animals))]

	return adjective, animal
}

// GeneratePascalName generates a new reproducible random name consisting of an adjective and animal name in pascal case
func GeneratePascalName(seed string) string {
	adjective, animal := getParts(seed)

	return caser.String(adjective) + caser.String(animal)
}

// GenerateCamelCaseName generates a new reproducible random name consisting of an adjective and animal name in camel case
func GenerateCamelCaseName(seed string) string {
	adjective, animal := getParts(seed)

	return strings.ToLower(adjective) + caser.String(animal)
}
