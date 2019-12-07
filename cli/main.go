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

	f, err := propositional.Parse("(((A ^ B) | (C ^ D)) | E)")
	if err != nil {
		return err
	}

	color.Info.Printf("parser result: %s\n", f)
	color.Info.Printf("identifiers: %v\n", f.ListIdentifiers())
	color.Info.Printf("truth table: %v\n", propositional.GetTruthTable(
		f, f.ListIdentifiers()))

	nf := propositional.GetConjunctiveNormalForm(f)
	color.Info.Printf("\nnormal form: %s\n", nf)
	color.Info.Printf("truth table: %v\n", propositional.GetTruthTable(
		nf, nf.ListIdentifiers()))

	color.Println("")
	return nil
}
