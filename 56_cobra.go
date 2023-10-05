package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "myapp"}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a hello message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets a user by name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("Hello, %s!\n", args[0])
		} else {
			fmt.Println("Hello!")
		}
	},
}

func main() {
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
