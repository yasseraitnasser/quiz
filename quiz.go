
package main

import	(
	"fmt"
	"os"
	"flag"
)

func	main() {
	var	fNameErrorMsg	string	= "a csv file in the format of 'question,answer'";
	var	limitErrorMsg	string	= "the time limit for the quiz in seconds";
	var	fileNamePtr	*string	= flag.String("csv", "problem.csv", fNameErrorMsg);
	var	limitPtr	*int	= flag.Int("limit", 30, limitErrorMsg);

	flag.Parse();
	fmt.Println("file name:", *fileNamePtr);
	fmt.Println("limit:", *limitPtr);
	file, err	:= os.Open(*fileNamePtr);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	defer	file.Close();
}
