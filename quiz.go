
package main

import	(
	"fmt"
	"os"
	"flag"
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

func	displayScore() {
}

func	main() {
	var	fNameErrorMsg	string	= "a csv file in the format of 'question,answer'";
	var	limitErrorMsg	string	= "the time limit for the quiz in seconds";
	var	fileNamePtr	*string	= flag.String("csv", "problem.csv", fNameErrorMsg);
	var	limitPtr	*int	= flag.Int("limit", 30, limitErrorMsg);

	flag.Parse();
	fmt.Println("file name:", *fileNamePtr);
	fmt.Println("limit:", *limitPtr);
	csvFile, err	:= os.Open(*fileNamePtr);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	defer	csvFile.Close();
	reader	:= csv.NewReader(csvFile);
	index	:= 1;
	for {
		lineTokens, err := reader.Read();
		if err == io.EOF {
			displayScore();
			return ;
		} else if err != nil {
			log.Fatal(err);
			// return ;
		}
		question	:= lineTokens[0];
		answerS		:= lineTokens[1];
		answerI, err	:= strconv.Atoi(answerS);
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Problem #%d:\t%s\t= %d\n", index, question, answerI);
		index++;
	}
}
