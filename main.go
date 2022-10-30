package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//type cmd interface {
//	getPrompt() string
//	exec()
//}

type MenuItem struct {
	Title       string
	CMDLong     string
	CMDShort    string
	description string
	minArgLen   int
}

var changePin = MenuItem{
	Title:       "Change Pin",
	CMDShort:    "-c",
	CMDLong:     "--change-pin",
	description: "<-c or --change-pin> <new four-digit pin>",
	minArgLen:   2,
}

var balance = MenuItem{
	Title:       "Check Account Balance",
	CMDShort:    "-b",
	CMDLong:     "--balance",
	description: "<-b or --balance>",
	minArgLen:   1,
}

var withdraw = MenuItem{
	Title:       "Withdraw funds from Account",
	CMDShort:    "-w",
	CMDLong:     "--withdraw",
	description: "<-w or --withdraw> <Amount in digit>",
	minArgLen:   2,
}

var deposit = MenuItem{
	Title:       "Deposit funds from Account",
	CMDShort:    "-d",
	CMDLong:     "--deposit",
	description: "<-d or --deposit> <Amount in digit>",
	minArgLen:   2,
}

var exit = MenuItem{
	Title:       "Exit program",
	CMDShort:    "-e",
	CMDLong:     "--exit",
	description: "<-e or --exit>",
	minArgLen:   1,
}

var menuItems [5]MenuItem = [5]MenuItem{
	changePin, deposit, withdraw, balance, exit,
}

var userPin string = "0000"
var bankBalance int

const TotalIncorrectPinAttemptsAllowed = 3

var pinAttemptCount = 0

func main() {
	welcome()
	runOperation()
}

func newLineUtil(numOfLines int) {
	count := 0
	for count < numOfLines {
		fmt.Println()
		count++
	}
}

func fineOutputUtil(text string) {
	fmt.Println(text)
	newLineUtil(1)
}

func welcome() {
	fineOutputUtil("########## Welcome to ATM CLI APP ##########")
}

func getPin() string {
	fineOutputUtil("Kindly input your 4-digit pin: \n >>>> if you have not changed your pin, use default '0000'")

	args := getPassedArgs(1)
	inputtedPin := args[0]
	if !isPinCorrectlyFormed(inputtedPin) {
		fineOutputUtil("Error: incorrectly formed pin, ensure inputted pin is 4 digits long")
		getPin()
	}

	if inputtedPin != userPin {
		pinAttemptCount += 1
		remainingPinAttempt := TotalIncorrectPinAttemptsAllowed - pinAttemptCount
		if remainingPinAttempt == 0 {
			fineOutputUtil("Too many incorrect pin attempt")
			exitProgram(1)
		}
		fineOutputUtil("Error: incorrect pin try again you have: " + strconv.Itoa(remainingPinAttempt))
		getPin()
	}
	return inputtedPin
}

func getPassedArgs(minArgs int) []string {
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.TrimSuffix(cmdString, "\n")
	args := strings.Fields(cmdString)
	lenOfArgPassed := len(args)
	if lenOfArgPassed < minArgs {
		fineOutputUtil(fmt.Sprintf("Error: At least %v arguments are needed\n >> Got %v Args instead\n", minArgs, lenOfArgPassed))
		exitProgram(1)
	}
	return args
}

func isPinCorrectlyFormed(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 4 {
		return false
	}

	for _, val := range pwR {
		if !unicode.IsNumber(val) {
			return false
		}
	}

	return true
}

func exitProgram(exitCode int) {
	fineOutputUtil("Good Bye 👋")
	os.Exit(exitCode)
}

func displayMenu() {
	getPin()

	fineOutputUtil("What Would you like to do today: ")
	fmt.Println("[shortCommand] [LongCommand] [Title] [usageDescription]")
	for _, menuItem := range menuItems {
		fineOutputUtil(fmt.Sprintf("%v %v '%v' '%v", menuItem.CMDShort, menuItem.CMDLong, menuItem.Title, menuItem.description))
	}
}

func runOperation() {
	displayMenu()
	desiredOp := getPassedArgs(1)

	switch desiredOp[0] {
	case changePin.CMDShort, changePin.CMDLong:
		validateArgSize(desiredOp, changePin)
		pinChange(desiredOp)
	case deposit.CMDLong, deposit.CMDShort:
		validateArgSize(desiredOp, deposit)
		depositFunds(desiredOp)
	case withdraw.CMDShort, withdraw.CMDLong:
		validateArgSize(desiredOp, withdraw)
		withdrawFunds(desiredOp)
	case balance.CMDLong, balance.CMDShort:
		validateArgSize(desiredOp, balance)
		checkBalance()
	case exit.CMDShort, exit.CMDLong:
		validateArgSize(desiredOp, exit)
		exitProgram(0)
	default:
		fineOutputUtil("Error: Invalid Command, check menu options below and try again")
		runOperation()
	}
	fineOutputUtil("Would you like to perform another operation (y for yes, or press any other key to exit)")
	performAgain := getPassedArgs(1)[0]
	if performAgain == "y" || performAgain == "Y" {
		runOperation()
	}
	exitProgram(0)
}

func validateArgSize(args []string, menuItem MenuItem) {
	if len(args) < menuItem.minArgLen {
		fineOutputUtil(fmt.Sprintf("Error: At least %v arguments are needed\n >> Got %v Args instead\n", menuItem.minArgLen, menuItem.minArgLen))
		exitProgram(1)
	}
}

func pinChange(args []string) {
	fineOutputUtil(">>>>>>>>>> PIN CHANGE <<<<<<<<<<")
	if !isPinCorrectlyFormed(args[1]) {
		fineOutputUtil("Error: Invalid pin format, ensure pin is a 4 digit number")
		return
	}
	userPin = args[1]
}

func depositFunds(args []string) {
	fineOutputUtil(">>>>>>>>>> DEPOSIT FUNDS <<<<<<<<<<")
	amount, err := strconv.Atoi(args[1])
	if err != nil || amount < 0 {
		fineOutputUtil("Error: Invalid Amount, Ensure amount is an integer greater than Zero")
		return
	}
	bankBalance = amount
}

func withdrawFunds(args []string) {
	fineOutputUtil(">>>>>>>>>> WITHDRAW FUNDS <<<<<<<<<<")
	amount, err := strconv.Atoi(args[1])
	if err != nil || amount < 0 {
		fineOutputUtil("Error: Invalid Amount, Ensure amount is an integer greater than Zero")
		return
	} else if amount > bankBalance {
		fineOutputUtil("Error: Insufficient Balance for withdrawal operation")
		return
	}
	bankBalance -= amount
}

func checkBalance() {
	fineOutputUtil(">>>>>>>>>> ACCOUNT BALANCE <<<<<<<<<<")
	fineOutputUtil(fmt.Sprintf("Your Account Balance is : ₦%v", bankBalance))
}
