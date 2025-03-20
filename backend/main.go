package main

import (
	"backend/config"
	"backend/router"
	"fmt"
	"os"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		CustomCallerFormatter: func(f *runtime.Frame) string {
			filename := path.Base(f.File)
			return fmt.Sprintf(" (%s:%d)", filename, f.Line)
		},
		FieldsOrder:     []string{"component", "category"},
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.Info("Logger initialized")
}

func main() {
	var port string
	var conf string

	var rootCmd = &cobra.Command{
		Use:   "CrackSpectra",
		Short: "CrackSpectra is a crack segmentation platform",
	}

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	rootCmd.PersistentFlags().StringVarP(&conf, "conf", "c", "/etc/CrackSpectura/config.prod.toml", "Path to configuration file")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("conf", rootCmd.PersistentFlags().Lookup("conf"))

	config.Init(viper.GetString("conf"))

	var runCommand = &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Println("Running as both producer and consumer")
			engine := router.New()
			port := viper.GetString("port")
			err := engine.Run(":" + port)
			if err != nil {
				panic(err)
			}
		},
	}

	rootCmd.AddCommand(runCommand)
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
