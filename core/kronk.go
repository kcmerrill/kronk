package core

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func say(topic, msg string) {
	Says(topic, msg)
}

// Says is what kronk outputs
func Says(topic, msg string) {
	fmt.Print("[", topic, "] ", msg, "\n")
}

func amConfused(topic, msg string) {
	IsConfused(topic, msg)
}

// IsConfused is what kronk outputs
func IsConfused(topic, msg string) {
	fmt.Print("[", topic, "] ", msg, "\n")
	os.Exit(1)
}

// NewKronk will init a new kronk
func NewKronk(args []string, content []byte) *Kronk {
	// double check the argument length first
	if len(args) == 0 {
		amConfused("error", "Please provide a 'id:expression'")
	}

	// make sure we have something to do
	if len(content) == 0 {
		amConfused("error", "Nothing to kronk")
	}

	k := &Kronk{
		args:    args,
		content: content,
	}

	// initialize/set our defaults
	k.regexes = make(map[string]string)
	k.matches = make(map[string][]string)
	k.cols = make([]string, 0)

	// parse the arguments <id:expression>
	k.parseArgs()

	// get to kronkin!
	k.kronkin()
	return k
}

// Kronk is the obj that stores our regexes and content
type Kronk struct {
	args     []string
	cols     []string
	regexes  map[string]string
	matches  map[string][]string
	content  []byte
	del      string
	passThru bool
}

func (k *Kronk) parseArgs() {
	for _, keyValue := range k.args {
		if col, regex, extractError := k.extractIDRegex(keyValue); extractError == nil {
			// save off the regexes for later
			k.regexes[col] = regex
			// silly maps ... need this to keep the order
			k.cols = append(k.cols, col)
		}
	}
}

func (k *Kronk) extractIDRegex(keyValue string) (string, string, error) {
	if strings.Contains(keyValue, ":") {
		kv := strings.SplitN(keyValue, ":", 2)
		return kv[0], kv[1], nil
	}
	return "", "", fmt.Errorf("Unable to find del in cli arument")
}

func (k *Kronk) kronkin() {
	for id, regex := range k.regexes {
		r, compileError := regexp.Compile(regex)
		if compileError != nil {
			amConfused("error", "Unable to compile '"+regex+"'")
		}
		matches := r.FindAllStringSubmatch(string(k.content), -1)
		for _, matchSet := range matches {
			if len(matchSet) == 2 {
				k.matches[id] = append(k.matches[id], matchSet[1])
			}
		}
	}
}

func (k *Kronk) validate() (bool, error) {
	if len(k.matches) == 0 {
		return false, fmt.Errorf("No matches found")
	}

	if len(k.matches) != len(k.cols) {
		return false, fmt.Errorf("1 or more columns did not yeild results. Please check your expressions and try again")
	}

	// grab the first index's length.
	// all others will be compared to the first
	// godspeed ...
	count := len(k.matches[k.cols[0]])
	for idx := range k.matches {
		fmt.Println("checking", idx)
		if len(k.matches[idx]) > count {
			return false, fmt.Errorf("'" + idx + "' matches greater than '" + k.cols[0] + "' withnumber of matches")
		}

		if len(k.matches[idx]) < count {
			return false, fmt.Errorf("'" + idx + "' matches less than '" + k.cols[0] + "' with number of matches")
		}

		fmt.Println(k.cols, len(k.matches[idx]), count)
	}
	return true, nil
}

// Display determines which output to display
func (k *Kronk) Display(out, del string, passThru bool) {
	if validated, errors := k.validate(); !validated {
		if passThru {
			fmt.Println(string(k.content))
			return
		}
		amConfused("error", errors.Error())
	}

	switch out {
	case "inline":
		k.inline(del)
	case "tsv":
		k.csv("\t")
	case "csv":
		k.csv(del)
	case "simple":
		if len(k.matches) >= 2 {
			k.csv(del)
			return
		}
		fallthrough
	default:
		k.simple()
	}
}

func (k *Kronk) inline(del string) {
	for row := 0; row < len(k.matches[k.cols[0]]); row++ {
		data := make([]string, 0)
		for _, col := range k.cols {
			data = append(data, col+":"+k.matches[col][row])
		}
		fmt.Println(strings.Join(data, del))
	}
}

func (k *Kronk) simple() {
	for row := 0; row < len(k.matches[k.cols[0]]); row++ {
		data := make([]string, 0)
		data = append(data, k.matches[k.cols[0]][row])
		fmt.Println(strings.Join(data, ""))
	}
}

func (k *Kronk) csv(del string) {
	// print out headers
	fmt.Println(strings.Join(k.cols, del))
	for row := 0; row < len(k.matches[k.cols[0]]); row++ {
		data := make([]string, 0)
		for _, col := range k.cols {
			data = append(data, k.matches[col][row])
		}
		fmt.Println(strings.Join(data, del))
	}
}
