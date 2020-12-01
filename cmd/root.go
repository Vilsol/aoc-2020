package cmd

import (
	"github.com/Vilsol/aoc-2020/registry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:  "aoc2020",
	Args: cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetEnvPrefix("aoc")
		viper.AutomaticEnv()

		_ = viper.ReadInConfig()

		level, err := log.ParseLevel(viper.GetString("log"))

		if err != nil {
			panic(err)
		}

		log.SetFormatter(&log.TextFormatter{
			ForceColors: viper.GetBool("colors"),
		})
		log.SetOutput(os.Stdout)
		log.SetLevel(level)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		day, err := strconv.Atoi(args[0])

		if err != nil {
			return err
		}

		return registry.RunDay(uint8(day))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().String("log", "info", "The log level to output")
	rootCmd.PersistentFlags().Bool("colors", false, "Force output with colors")

	_ = viper.BindPFlag("log", rootCmd.PersistentFlags().Lookup("log"))
	_ = viper.BindPFlag("colors", rootCmd.PersistentFlags().Lookup("colors"))
}
