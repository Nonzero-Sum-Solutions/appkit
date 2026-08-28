// TODO Move this to its own project.
package appkit

import (
	"fmt"
	"os"

	"github.com/Nonzero-Sum-Solutions/appkit/cmd"
	"github.com/Nonzero-Sum-Solutions/appkit/misc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	klog "github.com/go-kit/kit/log"

	gmodel "github.com/Nonzero-Sum-Solutions/appkit/model/gen/appkit"
)

var (
	version = "1.0.0"
	cfgFile string
)

func DoInit(cli1 gmodel.CLI) {
	logger := klog.NewLogfmtLogger(klog.NewSyncWriter(os.Stdout))

	misc.LogMessage(logger, "Logger initialized.")

	cli1.SetLogger(logger)

	misc.LogMessage(cli1.Logger(), fmt.Sprintf("CLI initialized. Appkit version %s", version))

	initConfig2 := func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		}

		viper.SetConfigName(cli1.App().ConfigName())
		viper.AddConfigPath("$HOME")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			misc.LogMessage(cli1.Logger(), fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
		}
	}

	cobra.OnInitialize(initConfig2)
}

type cliImpl struct {
	app              gmodel.App
	name             string
	shortDescription string
	logger           klog.Logger
	commandFactory   gmodel.CommandFactory
}

func (c *cliImpl) App() gmodel.App {
	return c.app
}

func (c *cliImpl) Name() string {
	return c.name
}

func (c *cliImpl) ShortDescription() string {
	return c.shortDescription
}

func (c *cliImpl) Logger() klog.Logger {
	return c.logger
}

func (c *cliImpl) SetLogger(logger klog.Logger) {
	c.logger = logger
}

func (c *cliImpl) NewRootCommand() (*cobra.Command, error) {
	return c.commandFactory.New()
}

func (c *cliImpl) Execute() {
	rootCmd, err := c.NewRootCommand()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func NewCLI(app gmodel.App, name, shortDescription string, commandFactory gmodel.CommandFactory) gmodel.CLI {
	return &cliImpl{
		app:              app,
		name:             name,
		shortDescription: shortDescription,
		commandFactory:   commandFactory,
	}
}
func NewCommandBase(cli gmodel.CLI) *cobra.Command {
	newCmd := &cobra.Command{
		Use:   cli.Name(),
		Short: cli.ShortDescription(),
	}
	newCmd.AddCommand(newVersionCommand(cli))
	newCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		fmt.Sprintf("config file (default is $HOME/%s.yaml)", cli.App().ConfigName()))

	return newCmd
}

// TODO Maybe make method on CLI.
func Execute(cli gmodel.CLI) {
	rootCmd, err := cli.NewRootCommand()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func newVersionCommand(cli gmodel.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version number of this command.",
		Run:   cmd.WrapRunner(doVersion, cli),
	}
}

func doVersion(cli gmodel.CLI, _ *cobra.Command, _ []string) error {
	fmt.Println(cli.App().Version())
	return nil
}
