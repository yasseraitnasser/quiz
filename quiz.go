
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

func	displayScore(index int, reader *csv.Reader, flag int) {
	total	:= index - flag + 1;
	for {
		_, err := reader.Read();
		if err == io.EOF {
			break ;
		} else if err != nil {
			log.Fatal(err)
			return ;
		}
		total++;
	}
	fmt.Println("You scored", index, "out of", total);
	os.Exit(0);
}

func	quizAndDisplayScore(csvFile *os.File) {
	reader	:= csv.NewReader(csvFile);
	index	:= 0;
	for {
		lineTokens, err := reader.Read();
		if err == io.EOF {
			displayScore(index, reader, 1);
		} else if err != nil {
			log.Fatal(err);
		}
		question	:= lineTokens[0];
		answerS		:= lineTokens[1];
		answerI, err	:= strconv.Atoi(answerS);
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Problem #%d:\t%s\t= ", index, question);
		var	userAnswerS	string;
		fmt.Scanln(&userAnswerS);
		userAnswerI, err	:= strconv.Atoi(userAnswerS);
		if err != nil {
			log.Fatal(err);
		}
		if (userAnswerI != answerI) {
			displayScore(index, reader, 0);
		}
		index++;
	}
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
	quizAndDisplayScore(csvFile);
}
