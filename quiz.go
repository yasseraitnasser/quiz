
package main

import	(
	"fmt"
	"os"
	"flag"
	"encoding/csv"
	"strings"
)

type	problem	struct {
	q	string
	a	string
}

func	displayScore(index int, total int) {
	fmt.Println("You scored", index, "out of", total);
	os.Exit(0);
}

func	parseLines(lines [][]string) []problem {
	ret	:= make([]problem, len(lines));
	for i, line := range lines {
		ret[i] = problem {
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return	ret;
}

func	main() {
	var	fNameErrorMsg	string	= "a csv file in the format of 'question,answer'";
	var	limitErrorMsg	string	= "the time limit for the quiz in seconds";
	var	fileNamePtr	*string	= flag.String("csv", "problem.csv", fNameErrorMsg);
	var	limitPtr	*int	= flag.Int("limit", 30, limitErrorMsg);

	flag.Parse();
	// debug
	fmt.Println("file name:", *fileNamePtr);
	fmt.Println("limit:", *limitPtr);
	// debug
	csvFile, err	:= os.Open(*fileNamePtr);
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *fileNamePtr);
		os.Exit(1);
	}
	defer	csvFile.Close();
	reader	:= csv.NewReader(csvFile);
	lines, err := reader.ReadAll();
	if err != nil {
		fmt.Printf("Failed to read the CSV file: %s\n", *fileNamePtr);
		os.Exit(1);
	}
	problems	:= parseLines(lines);
	length		:= len(problems);
	for i, p := range problems {
		fmt.Printf("Problem #%d:\t%s = ", i+1, p.q);
		var	answer	string;
		fmt.Scanf("%s\n", &answer);
		if (answer != p.a) {
			displayScore(i, length);
		}
	}
	displayScore(length, length);
}
