package main

import (
	"fmt"
	"mexok/bkt/package/bkt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	var force bool
	saveFlagSet := pflag.NewFlagSet("save", pflag.ContinueOnError)
	saveFlagSet.BoolVarP(&force, "force", "f", false, "force saving, overwriting previous labels")

	var namespaces bool
	listFlagSet := pflag.NewFlagSet("list", pflag.ContinueOnError)
	listFlagSet.BoolVarP(&namespaces, "namespaces", "n", false, "List namespaces instead")

	var all bool
	deleteFlagSet := pflag.NewFlagSet("delete", pflag.ContinueOnError)
	deleteFlagSet.BoolVar(&all, "all", false, "")

	if len(os.Args) < 2 {
		bkt.PrintGlobalHelp()
		os.Exit(1)
	}
	var err error
	switch os.Args[1] {
	case "g", "ge", "get":
		if len(os.Args) != 3 {
			bkt.PrintGetHelp()
			os.Exit(1)
		}
		err = bkt.GetCmd(os.Args[2])
	case "s", "sa", "sav", "save":
		if len(os.Args) < 3 {
			bkt.PrintSaveHelp()
			os.Exit(1)
		}
		err = saveFlagSet.Parse(os.Args[2 : len(os.Args)-1])
		if err == nil {
			err = bkt.SaveCmd(os.Args[len(os.Args)-1], force)
		}
	case "d", "de", "del", "dele", "delet", "delete":
		err = deleteFlagSet.Parse(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		unparsed := deleteFlagSet.Args()
		label := ""
		if len(unparsed) > 1 {
			bkt.PrintDeleteHelp()
			os.Exit(1)
		} else if len(unparsed) == 1 {

		}
		err = bkt.DeleteCmd(label, all)
	case "l", "li", "lis", "list", "ls":
		err = listFlagSet.Parse(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			bkt.PrintListHelp()
			fmt.Printf("\nOptions:\n")
			listFlagSet.PrintDefaults()
			os.Exit(1)
		}
		err = bkt.ListCmd(namespaces)
	case "h", "he", "hel", "help":
		bkt.PrintHelpHelp()
		os.Exit(1)
	default:
		bkt.PrintGlobalHelp()
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
