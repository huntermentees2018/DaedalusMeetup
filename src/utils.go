package src

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

const (
	layoutISO = "1/02/2006 15:04:05"
)

type configFormat struct {
	postgresURI string
	calendarID  string
}

// GetConfig returns a config string from a config.json file
func GetConfig(key string) string {
	config := make(map[string]string)
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &config)
	return config[key]
}

// PrefCheck counts which runes "M", "A", "E" are both in a and b in a map and returns it as a map
func PrefCheck(a, b []rune) map[rune]bool {
	times := map[rune]uint{77: 0, 65: 0, 69: 0}
	timesC := map[rune]bool{77: false, 65: false, 69: false}
	for _, c := range a {
		times[c]++
	}
	for _, c := range b {
		times[c]++
	}
	for k, v := range times {
		if v == 2 {
			timesC[k] = true
		}
	}
	return timesC
}

// IsNewerTime checks if the first param time is newer than the secondk
func IsNewerTime(timeOne string, timeTwo string) bool {
	t1, _ := time.Parse(layoutISO, timeOne)
	t2, _ := time.Parse(layoutISO, timeTwo)
	return t1.After(t2)
}

// DaysL converts day strings to runes (e.g: "Monday, Tuesday" => ["M", "T"])
func DaysL(days string) string {
	if days == "" {
		return ""
	}
	s := strings.Split(days, ",")
	c := []rune{}
	for _, day := range s {
		day = strings.TrimSpace(day)
		c = append(c, rune(day[0]))
	}
	return string(c)
}

// RegexCommand just checks if the string is a certain string
func RegexCommand(text string) string {
	if strings.Contains(text, "unsubscribe") {
		return "unsubscribe"
	} else if strings.Contains(text, "subscribe") {
		return "subscribe"
	} else if strings.Contains(text, "list") {
		return "list"
	} else if strings.Contains(text, "help") {
		return "help"
	}

	return "idk"
}
