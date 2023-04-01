package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
)

///					 ///
// OpenRedirectFinder //
///					 ///

///
// Misc. Functions
///

// function to print help
func PrintHelp() {
	fmt.Println(color.Red + "orf - Open Redirect Finder" + color.Reset + "\n\n" + color.InBold("Usage:") + "\n$ ./orf <endpoints file> <runtime mode> <payload file>\n\n" + color.InBold("Example 1 - (Only tests with single payload):") + " \n\n\t$ ./orf endpoints.txt 0\n\n" + color.InBold("Example 2 - (Tests with every payload defined in payloads.txt - warning - lots more traffic):") + "\n\n\t$ ./orf endpoints.txt 1 payloads.txt\n\n")
	os.Exit(1)
}

/// function to check whether user input contains a value.
func CheckUserInput() {
	if len(os.Args) == 1 {
		PrintHelp()
	}
}

/* Opens a wordlist file and places each line into a string array. */
func ReadFileToStringArray(wordlist_file_path string) []string {
	//open wordlist
	wordlist, _ := os.Open(wordlist_file_path)
	defer wordlist.Close()
	// read lines from wordlist
	scanner := bufio.NewScanner(wordlist)
	scanner.Split(bufio.ScanLines)
	var wordlist_lines []string
	for scanner.Scan() {
		// put lines into string array
		wordlist_lines = append(wordlist_lines, scanner.Text())
	}
	return wordlist_lines
}

func regexCheck(result string) bool {
	r, _ := regexp.MatchString("=http", result)
	if r {
		return true
	} else {
		return false
	}
}

func ExtractEndpointsContainingURLs(endpoints_file []string) []string {
	var endpointsWithUrls []string
	for _, endpoint := range endpoints_file {
		// if endpoint contains =http add to new string array
		//string := "http://foo.com/?foo=http://"
		if regexCheck(endpoint) {
			endpointsWithUrls = append(endpointsWithUrls, endpoint)
		}
	}
	return endpointsWithUrls
}

// returns a new string array of endpoints where every instance of =http is replaced with the basic payload "=https://www.google.com/"
func InsertBasicPayload(endpoints []string) []string {
	var ModifiedWithBasicPayloadsEndpointArray []string
	for _, endpoint := range endpoints {
		modifiedEndpoint := strings.Replace(endpoint, "=http", "=https://www.google.com/", -1)
		ModifiedWithBasicPayloadsEndpointArray = append(ModifiedWithBasicPayloadsEndpointArray, modifiedEndpoint)
	}
	return ModifiedWithBasicPayloadsEndpointArray
}

// returns a new string array (ModifiedWithPayloadsListArray) of endpoints. for each string passed, but each instance of =http with each line in the payloads file.
func InsertPayloadList(endpoints []string, payloads []string) []string {
	var ModifiedWithPayloadsListArray []string

	return ModifiedWithPayloadsListArray
}

func handleQueryResponse() {

}

func makeQuery(modifiedEndpoints []string) {

}

func main() {
	// declare variables
	//var wg sync.WaitGroup // for running cmd commands simultaneously
	var arg1 string // to store the name of the payloads file
	var mode int    // to store the name of the endpoints file
	var arg3 string // for payloads file
	var modifiedEndpoints []string

	// get date
	//date := time.Now().Format("01-02-2006")

	// get full tool run time
	//full_runtime := time.Now()

	// check user inputted an the needed arguements
	CheckUserInput()

	// handle arguements.

	//handle hosts file
	arg1 = os.Args[1]
	//handle run time mode -- can be 1 or 0. 0 = run with one payload, 1 = run with bypasses from payloads file,
	mode, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("\n ERROR - error processing run time mode. Did you input \"1\" or \"0\"?")
	}

	fmt.Println(os.Args)

	if len(os.Args) == 4 {
		arg3 = os.Args[3]
	}

	// open user input files, confirming they exist, if not exit with error.

	// print title
	fmt.Println(color.Ize(color.Red, "orf- Open Redirect Finder"))

	// check user inputted an the needed arguements
	CheckUserInput()

	//read endpoints file & payloads files to string arrays

	payloadsFile := ReadFileToStringArray(arg3)
	endpointsFile := ReadFileToStringArray(arg1)

	endpointsContainingUrls := ExtractEndpointsContainingURLs(endpointsFile)

	if mode == 1 {
		fmt.Println("foo")
		modifiedEndpoints = InsertPayloadList(endpointsContainingUrls, payloadsFile)
	} else {
		modifiedEndpoints = InsertBasicPayload(endpointsContainingUrls)
	}
	//fmt.Println(payloadsFile)
	//fmt.Println(endpointsContainingUrls)
	fmt.Print(modifiedEndpoints)
}
