
package main

import	(
	"fmt"
	"os"
)

func	print_and_exit(msg string, exit_code int) {
	fmt.Println(msg);
	os.Exit(exit_code);
}

func	not_enough_args() {
	print_and_exit(`Error:	the program takes at least one paramter
	type './quiz -h' for help`, 1);
}

func	usage() {
	print_and_exit(`Usage of ./quiz:
	-csv string
		a csv file in the format of 'question,answer' (default "problems.csv")
	-limit int
		the time limit for the quiz in seconds (default 10)`, 0);
}

func	main() {
	args	:= os.Args[1:];

	if (len(args) < 1) {
		not_enough_args();}
	if (args[0] == "-h") {
		usage();}
}
