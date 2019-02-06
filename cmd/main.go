package main

import (
	"medbrat-server/configs"
	"medbrat-server/network"
	"medbrat-server/usecase"
)

func main() {
	// TODO remove
	blood := usecase.NewBloodGeneral()
	network.NewOCR().RecognizeGeneralBloodTest(&blood, &configs.SampleBase64Image)

	network.Main()
}
