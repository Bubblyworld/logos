package cmd

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "logos",
	Short: "Logos is a CLI multitool for manipulating formal logic sentences",
	Long: `Logos is a CLI multitool for manipulating sentences in various formal logics,
such as propositional, modal and description logics. Logos can currently only 
perform normal form conversions, with future plans to do satisfiability
checking and automatic proof discovery.`,
}

func Execute() {
	defer func() {
		color.Print("") // reset ANSI codes on exit
	}()

	if err := rootCmd.Execute(); err != nil {
		errorln(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config", "", "config file (default is $HOME/.cli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			errorln(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		infoln("Using config file: " + viper.ConfigFileUsed())
	}
}

func errorln(err error) {
	color.Error.Println(err.Error())
}

func printf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func println(msg string) {
	fmt.Println(msg)
}

func infof(msg string, args ...interface{}) {
	color.Info.Printf(msg, args...)
}

func infoln(msg string) {
	color.Info.Println(msg)
}

func lightf(msg string, args ...interface{}) {
	color.Light.Printf(msg, args...)
}

func lightln(msg string) {
	color.Light.Println(msg)
}

func primaryf(msg string, args ...interface{}) {
	color.Primary.Printf(msg, args...)
}

func primaryln(msg string) {
	color.Primary.Println(msg)
}
