package bkt

import "fmt"

func PrintGlobalHelp() {
	fmt.Println("Usage: bkt [subcommand]")
	fmt.Println("Possible values:")
	fmt.Println(" * load")
	fmt.Println(" * save")
	fmt.Println(" * help")
}

func PrintGetHelp() {
	fmt.Println("Usage: bkt g[et] {label}")
	fmt.Println("Prints saved location of label. Exits with non-zero if label does not exist.")
}

func PrintSaveHelp() {

}

func PrintDeleteHelp() {
	fmt.Println("Usage: bkt d[elete] [options] [label]")
	fmt.Println("Deletes label in current namespace")
}

func PrintListHelp() {
	fmt.Println("Usage: bkt l[ist] [options]")
	fmt.Println("       bkt ls [options]")
	fmt.Println("Lists all labels in current namespace")
}

func PrintHelpHelp() {

}
