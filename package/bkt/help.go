package bkt

import (
	"fmt"

	"github.com/spf13/pflag"
)

func PrintGlobalHelp() {
	fmt.Println("Usage: bkt <subcommand> [...]")
	fmt.Println("       bkt h[elp] <subcommand>")
	fmt.Println("Possible values for subcommand:")
	fmt.Println(" * s[ave]      ... saves directories using labels")
	fmt.Println(" * g[et]       ... returns a saved directory")
	fmt.Println(" * u[se]       ... switch between namespaces")
	fmt.Println(" * l[ist], ls  ... list labels and namespaces")
	fmt.Println(" * o[verview]  ... shortcut for 'bkt ls -ln'")
	fmt.Println(" * d[elete]    ... deletes labels and namespaces")
	fmt.Println("Use bkt help <subcommand> to show help for a specific subcommand")
}

func PrintSaveHelp(flagSet *pflag.FlagSet) {
	fmt.Println("Usage: bkt s[ave] [options] <label>")
	fmt.Println("Saves current directory in current namespace using label.")
	fmt.Println("")
	fmt.Println("Options:")
	flagSet.PrintDefaults()
}

func PrintGetHelp(flagSet *pflag.FlagSet) {
	fmt.Println("Usage: bkt g[et] <label>")
	fmt.Println("Prints saved location of label. Exits with non-zero if label does not exist.")
	fmt.Println("")
	fmt.Println("Options:")
	flagSet.PrintDefaults()
}

func PrintUseHelp(flagSet *pflag.FlagSet) {
	fmt.Println("Usage: bkt u[se] [options] <namespace>")
	fmt.Println("Switch to namespace")
	fmt.Println("")
	fmt.Println("Options:")
	flagSet.PrintDefaults()
}

func PrintListHelp(flagSet *pflag.FlagSet) {
	fmt.Println("Usage: bkt l[ist] [options]")
	fmt.Println("       bkt ls [options]")
	fmt.Println("Lists all labels in current namespace")
	fmt.Println("")
	fmt.Println("Options:")
	flagSet.PrintDefaults()
}

func PrintOverviewHelp() {
	fmt.Println("Usage: bkt o[verview]")
	fmt.Println("Shortcut for 'bkt ls -ln'")
}

func PrintDeleteHelp(flagSet *pflag.FlagSet) {
	fmt.Println("Usage: bkt d[elete] [options] [label]")
	fmt.Println("Deletes label in current namespace")
	fmt.Println("")
	fmt.Println("Options:")
	flagSet.PrintDefaults()
}
