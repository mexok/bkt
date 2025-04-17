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

	var create bool
	namespaceFlagSet := pflag.NewFlagSet("namespace", pflag.ContinueOnError)
	namespaceFlagSet.BoolVarP(&create, "create", "c", false, "Create namespace instead")

	var namespaces bool
	listFlagSet := pflag.NewFlagSet("list", pflag.ContinueOnError)
	listFlagSet.BoolVarP(&namespaces, "namespaces", "n", false, "List namespaces instead")

	var namespace bool
	var yes bool
	deleteFlagSet := pflag.NewFlagSet("delete", pflag.ContinueOnError)
	deleteFlagSet.BoolVarP(&namespace, "namespace", "n", false, "Deletes current namespace and switches to default namespaces")
	deleteFlagSet.BoolVarP(&yes, "yes", "y", false, "Confirm to delete. Only required for deletion of namespace")

	if len(os.Args) < 2 {
		bkt.PrintGlobalHelp()
		os.Exit(1)
	}
	var err error
	switch os.Args[1] {
	case "s", "sa", "sav", "save":
		err = saveFlagSet.Parse(os.Args[2:])
		if err == nil {
			if len(saveFlagSet.Args()) != 1 {
				bkt.PrintSaveHelp(saveFlagSet)
				os.Exit(1)
			}
			err = bkt.SaveCmd(saveFlagSet.Arg(0), force)
		}
	case "g", "ge", "get":
		if len(os.Args) != 3 {
			bkt.PrintGetHelp()
			os.Exit(1)
		}
		err = bkt.GetCmd(os.Args[2])
	case "n", "na", "nam", "name", "names", "namesp", "namespa", "namespac", "namespace":
		err = namespaceFlagSet.Parse(os.Args[2:])
		if err == nil {
			if len(namespaceFlagSet.Args()) != 1 {
				bkt.PrintNamespaceHelp(namespaceFlagSet)
				os.Exit(1)
			}
			err = bkt.NamespaceCmd(namespaceFlagSet.Arg(0), create)
		}
	case "l", "li", "lis", "list", "ls":
		err = listFlagSet.Parse(os.Args[2:])
		if err == nil {
			if len(listFlagSet.Args()) != 0 {
				bkt.PrintListHelp(listFlagSet)
				os.Exit(1)
			}
			err = bkt.ListCmd(namespaces)
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
			bkt.PrintDeleteHelp(deleteFlagSet)
			os.Exit(1)
		} else if len(unparsed) == 1 {
			label = unparsed[0]
		}
		err = bkt.DeleteCmd(label, namespace, yes)
	case "h", "he", "hel", "help":
		if len(os.Args) != 3 {
			bkt.PrintGlobalHelp()
			os.Exit(1)
		}
		switch os.Args[2] {
		case "s", "sa", "sav", "save":
			bkt.PrintSaveHelp(saveFlagSet)
		case "g", "ge", "get":
			bkt.PrintGetHelp()
		case "n", "na", "nam", "name", "names", "namesp", "namespa", "namespac", "namespace":
			bkt.PrintNamespaceHelp(namespaceFlagSet)
		case "l", "li", "lis", "list", "ls":
			bkt.PrintListHelp(listFlagSet)
		case "d", "de", "del", "dele", "delet", "delete":
			bkt.PrintDeleteHelp(deleteFlagSet)
		default:
			bkt.PrintGlobalHelp()
		}
	default:
		bkt.PrintGlobalHelp()
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
