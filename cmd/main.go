package main

import (
	"medbrat-server/configs"
	"medbrat-server/network"
)

func main() {
	// TODO remove
	network.NewOCR().RecognizeGeneralBloodTest(&configs.SampleBase64Image)

	network.Main()
}
