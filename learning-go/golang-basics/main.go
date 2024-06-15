package main

const fileName = "sampleFiles.txt"

func main() {
	cards := deckFromFile(fileName)
	cards.saveToFile(fileName)
}
