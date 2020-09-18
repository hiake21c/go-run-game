package main

import (
	"go-run-game/global"
	"go-run-game/scenemanager"
	"go-run-game/scenes"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {

	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 2.0, "Jongmin Kim's GO run")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
