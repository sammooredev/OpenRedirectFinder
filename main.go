package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	fmt.Println(color.Red + "OpenRedirectFinder" + color.Reset + "\n\n" + color.InBold("Usage:") + "\n$ ./OpenRedirectFinder <hosts file> <payload file> <verbose>\n\n" + color.InBold("Example 1 - (Only prints successful redirects):") + " \n\n\t$ ./OpenRedirectFinder alive_hosts.txt payloads.txt\n\n\\" + color.InBold("Example 2 - (Prints every request + status code):") + "\n\n\t$ ./OpenRedirectFinder alive_hosts.txt payloads.txt verbose\n\n")
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

func InsertPayloads(endpoints []string) []string {
	var ModifiedEndpointArray []string
	for _, endpoint := range endpoints {
		modifiedEndpoint := strings.Replace(endpoint, "=http", "=https://www.google.com/", -1)
		ModifiedEndpointArray = append(ModifiedEndpointArray, modifiedEndpoint)
	}
	return ModifiedEndpointArray
}

func handleQueryResponse()

func makeQuery(modifiedEndpoints []string) {

}

func main() {
	// declare variables
	//var wg sync.WaitGroup // for running cmd commands simultaneouslya
	var arg1 string // to store the name of the payloads file
	var arg2 string // to store the name of the endpoints file

	// get date
	//date := time.Now().Format("01-02-2006")

	// get full tool run time
	//full_runtime := time.Now()

	// check user inputted an the needed arguements
	CheckUserInput()

	// handle arguements.
	// 1. payloads file 2. hosts file
	arg1 = os.Args[1]
	arg2 = os.Args[2]

	// open user input files, confirming they exist, if not exit with error.

	// print title
	fmt.Println(color.Ize(color.Red, "OpenRedirectFinder"))

	// check user inputted an the needed arguements
	CheckUserInput()

	//payloadsFile := ReadFileToStringArray(arg1)
	ReadFileToStringArray(arg2)
	endpointsFile := ReadFileToStringArray(arg1)

	endpointsContaingUrls := ExtractEndpointsContainingURLs(endpointsFile)

	modifiedEndpoints := InsertPayloads(endpointsContaingUrls)

	//fmt.Println(payloadsFile)
	fmt.Println(endpointsContaingUrls)
	fmt.Print(modifiedEndpoints)
}
