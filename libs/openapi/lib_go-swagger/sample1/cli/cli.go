// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/syunkitada/programming_go/libs/openapi/lib_go-swagger/sample1/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.TodoList, error) {
	hostname := viper.GetString("hostname")
	viper.SetDefault("base_path", client.DefaultBasePath)
	basePath := viper.GetString("base_path")
	scheme := viper.GetString("scheme")

	r := httptransport.New(hostname, basePath, []string{scheme})
	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/io.goswagger.examples.todo-list.v1+json"] = runtime.JSONConsumer()

	r.Producers["application/io.goswagger.examples.todo-list.v1+json"] = runtime.JSONProducer()

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))
	rootCmd.PersistentFlags().String("base-path", client.DefaultBasePath, fmt.Sprintf("For example: %v", client.DefaultBasePath))
	viper.BindPFlag("base_path", rootCmd.PersistentFlags().Lookup("base-path"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	// register security flags
	// add all operation groups
	operationGroupTodosCmd, err := makeOperationGroupTodosCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupTodosCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

func makeOperationGroupTodosCmd() (*cobra.Command, error) {
	operationGroupTodosCmd := &cobra.Command{
		Use:  "todos",
		Long: ``,
	}

	operationAddOneCmd, err := makeOperationTodosAddOneCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationAddOneCmd)

	operationDestroyOneCmd, err := makeOperationTodosDestroyOneCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationDestroyOneCmd)

	operationFindTodosCmd, err := makeOperationTodosFindTodosCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationFindTodosCmd)

	operationUpdateOneCmd, err := makeOperationTodosUpdateOneCmd()
	if err != nil {
		return nil, err
	}
	operationGroupTodosCmd.AddCommand(operationUpdateOneCmd)

	return operationGroupTodosCmd, nil
}
