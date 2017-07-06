package Lib1

import (
	"Testprojekt5/Testprojekt5_func_Lib2"
	"fmt"
	"log"
	_ "os/exec"
	"strconv"

	_ "github.com/howeyc/gopass"
	"github.com/skratchdot/open-golang/open"
)

func AskForConfirmation() string {

	var response string
	fmt.Println("Enter your Username:")
	_, err := fmt.Scanln(&response)
	if err != nil {
		return AskForConfirmation()
	}
	okayResponses := []string{"FynnOke"}
	nokayResponses := []string{"Lasse"}
	Simon := []string{"Simon"}
	terminated := []string{"terminate", "term"}
	if containsString(okayResponses, response) {

		return LogInOke()
	} else if containsString(nokayResponses, response) {

		return LogInLasse()
	} else if containsString(Simon, response) {

		return LogInSimon()
	} else if containsString(terminated, response) {

		return terminate()
	} else {
		fmt.Println("Please type yes or no and then press enter, or terminate the program with: terminate")
		return AskForConfirmationreturn()
	}

}

func AskForConfirmationreturn() string {
	return AskForConfirmation()
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func LogInOke() string {
	fmt.Println("Enter your Password, Fynn Oke")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return AskForConfirmation()
	}
	pass := []string{"qwer1234"}
	if containsString(pass, response) {
		return WDYWTD()
	} else {
		return AskForConfirmation()
	}
}

func LogInLasse() string {
	pass := []string{"abc123"}
	fmt.Println("Enter your Password, Lasse")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return AskForConfirmation()
	}
	if containsString(pass, response) {
		return WDYWTD()
	} else {
		return AskForConfirmation()
	}
}

func LogInSimon() string {
	pass := []string{"asdf6789"}
	fmt.Println("Enter your Password, Simon")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return AskForConfirmation()
	}
	if containsString(pass, response) {
		return WDYWTD()
	} else {
		return AskForConfirmation()
	}
}

func WDYWTD() string {
	fmt.Println("Welcome. What do you want to do?")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return WDYWTDreturn()
	}
	response1 := []string{"terminate", "term"}
	response2 := []string{"calculate", "calc"}
	response3 := []string{"website", "web"}
	response4 := []string{"help"}
	response5 := []string{"locateIP", "locIP", "IP", "ip"}
	if containsString(response1, response) {
		return terminate()
	} else if containsString(response2, response) {
		return calc()
	} else if containsString(response3, response) {
		return openWebsite()
	} else if containsString(response4, response) {
		fmt.Println("Possible commands are:")
		fmt.Println(" ")
		fmt.Println("terminate, term:	Terminates the program.")
		fmt.Println(" ")
		fmt.Println("calculate, calc:	Is a calculator.")
		fmt.Println(" ")
		fmt.Println("website, web:		Opens a website.")
		fmt.Println(" ")
		return WDYWTDreturn()
	} else if containsString(response5, response) {
		return IPv4()
	} else {
		fmt.Println("Command not recognized. Commands are: terminate, calculate, website, help")
		return WDYWTDreturn()
	}
}

func WDYWTDreturn() string {
	return WDYWTD()
}

func terminate() string {
	var response string
	fmt.Println("Program was terminated. Do you want to close the command prompt?")
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return "Bye!"
	} else if containsString(nokayResponses, response) {
		return AskForConfirmation()
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return terminatereturn()
	}
}

func terminatereturn() string {
	return terminate()
}

func terminatenotImplemented() string {
	var response string
	fmt.Println("Function not implemented right now, program terminated. Do you want to close the command prompt?")
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return "Bye!"
	} else if containsString(nokayResponses, response) {
		return AskForConfirmation()
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return terminatenotImplementedreturn()
	}
}

func terminatenotImplementedreturn() string {
	return terminatenotImplementedreturn()
}

func calc() string {
	var X float64
	var Y float64
	var response string
	fmt.Println("Enter your first Number:")
	_, err := fmt.Scanln(&X)
	if err != nil {
		return calcreturn()
	}
	fmt.Println("Enter your second Number:")
	_, err = fmt.Scanln(&Y)
	if err != nil {
		return calcreturn()
	}
	fmt.Println("What do you want to do with these? Commands are: add, sub, mul, div, sqrt, pow")
	_, err = fmt.Scanln(&response)
	if err != nil {
		return calcreturn()
	}
	addi := []string{"add"}
	subt := []string{"sub"}
	mult := []string{"mul"}
	divi := []string{"div"}
	sqrt := []string{"sqrt"}
	pow := []string{"pow"}
	cancel := []string{"cancel"}
	if containsString(addi, response) {
		fmt.Println(strconv.FormatFloat(Lib2.Add(X, Y), 'f', 6, 64))
		return calcreturn()
	} else if containsString(subt, response) {
		fmt.Println(strconv.FormatFloat(Lib2.Sub(X, Y), 'f', 6, 64))
		return calcreturn()
	} else if containsString(mult, response) {
		fmt.Println(strconv.FormatFloat(Lib2.Mul(X, Y), 'f', 6, 64))
		return calcreturn()
	} else if containsString(divi, response) {
		fmt.Println(strconv.FormatFloat(Lib2.Div(X, Y), 'f', 6, 64))
		return calcreturn()
	} else if containsString(cancel, response) {
		fmt.Println("Canceling...")
		return WDYWTD()
	} else if containsString(sqrt, response) {
		fmt.Println("Squareroot of your first Number: ", strconv.FormatFloat(Lib2.Sqrt(X), 'f', 6, 64))
		fmt.Println("Squareroot of your second Number: ", strconv.FormatFloat(Lib2.Sqrt(Y), 'f', 6, 64))
		return calcreturn()
	} else if containsString(pow, response) {
		fmt.Println(strconv.FormatFloat(Lib2.Pow(X, Y), 'f', 6, 64))
		return calcreturn()
	} else {
		return calcreturn()
	}
}

func calcreturn() string {
	return calc()
}

func openWebsite() string {
	fmt.Println("Which website do you want to open? Possible websites: Google, YouTube, Wikipedia, QSC, https://your.website")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		return openWebsitereturn()
	}
	googleDOTcom := []string{"google", "Google", "google.com"}
	youtubeDOTcom := []string{"youtube", "Youtube", "YouTube", "youtube.com"}
	wikipediaDOTorg := []string{"wikipedia", "WikipediA", "Wikipedia", "wikipedia.org", "wiki"}
	qscDOTde := []string{"qsc", "QSC", "qsc ag", "qscag", "QSCAG", "qsc ag", "qsc.de"}
	cancel := []string{"cancel"}
	if containsString(googleDOTcom, response) {
		fmt.Println("Website is getting opened...")
		open.Run("https://google.com/")
		return openWebsitereturn()
	} else if containsString(youtubeDOTcom, response) {
		fmt.Println("Website is getting opened...")
		open.Run("https://youtube.com/")
		return openWebsitereturn()
	} else if containsString(wikipediaDOTorg, response) {
		fmt.Println("Website is getting opened...")
		open.Run("https://wikipedia.org/")
		return openWebsitereturn()
	} else if containsString(qscDOTde, response) {
		fmt.Println("Website is getting opened...")
		open.Run("https://qsc.de/")
		return openWebsitereturn()
	} else if containsString(cancel, response) {
		fmt.Println("Canceling...")
		return WDYWTD()
	} else {
		fmt.Println("Website is getting opened... If not, check your input.")
		open.Run(response)
		return openWebsitereturn()
	}
}

func openWebsitereturn() string {
	return openWebsite()
}

func IPv4() string {
	fmt.Println("")
	fmt.Println(Lib2.LocalIP())
	fmt.Println("")
	return WDYWTD()
}
