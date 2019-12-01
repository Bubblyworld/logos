package main

import (
	"os"

	"github.com/bubblyworld/logos/ops/propositional"
	"github.com/bubblyworld/logos/state"
	"github.com/gookit/color"
)

func main() {
	if err := run(); err != nil {
		color.Error.Printf("error in main loop: %s", err)
		color.Println("")
		os.Exit(1)
	}
}

func run() error {
	_, err := state.New()
	if err != nil {
		return err
	}

	f, err := propositional.Parse("(A →       ¬   (A ∧ B))")
	if err != nil {
		return err
	}

	color.Info.Printf("parser result: %s", f)
	color.Println("")
	return nil
}
