
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

func	displayScore(index int, reader *csv.Reader) {
	total	:= index;
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
	fmt.Println("You scored", index - 1, "out of", total);
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
			displayScore(index, reader);
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
		fmt.Printf("Problem #%d:\t%s\t= ", index, question);
		var	userAnswerS	string;
		fmt.Scanln(&userAnswerS);
		userAnswerI, err	:= strconv.Atoi(userAnswerS);
		if err != nil {
			log.Fatal(err);
		}
		if (userAnswerI != answerI) {
			displayScore(index, reader);
			return ;
		}
		index++;
	}
}
