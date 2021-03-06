package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

// GenerateTimeInterval finds the time to pair
func GenerateTimeInterval(dayOfWeek string, timeOfDay rune) (string, string) {
	var daysOfWeek = map[string]time.Weekday{
		"Sunday":    time.Sunday,
		"Monday":    time.Monday,
		"Tuesday":   time.Tuesday,
		"Wednesday": time.Wednesday,
		"Thursday":  time.Thursday,
		"Friday":    time.Friday,
		"Saturday":  time.Saturday,
	}

	var start, end time.Time

	for i := 1; i < 6; i++ {
		now := time.Now()
		weekDay := now.AddDate(0, 0, i).Weekday()
		if daysOfWeek[dayOfWeek] == weekDay {
			start = now.AddDate(0, 0, i)
			end = now.AddDate(0, 0, i)
			if timeOfDay == 77 {
				start = start.Add(8 * time.Hour)
				end = end.Add(8 * time.Hour).Add(30 * time.Minute)
			} else if timeOfDay == 65 {
				start = start.Add(12 * time.Hour)
				end = end.Add(12 * time.Hour).Add(30 * time.Minute)
			} else {
				start = start.Add(17 * time.Hour)
				end = end.Add(17 * time.Hour).Add(30 * time.Minute)
			}
		}
	}
	log.Printf("start.Format(time.RFC3339) = %+v\n ", start.Format(time.RFC3339))
	return start.Format(time.RFC3339), end.Format(time.RFC3339)
}

// PrefCheck counts which runes "M", "A", "E" are both in a and b in a map and returns it as a map
// sample input ("ME", "MAE")
func PrefCheck(a, b string) map[rune]bool {
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

func idInSlice(a uint, ids []uint) bool {
	for _, b := range ids {
		if b == a {
			return true
		}
	}
	return false
}

// RemoveIDFromTimeTable removes an Id from the time table
func RemoveIDFromTimeTable(timeTable map[string]map[rune][]uint, idsToRemove []uint) map[string]map[rune][]uint {
	for a, day := range timeTable {
		for time, ids := range day {
			temp := []uint{}
			for _, id := range ids {
				if !idInSlice(id, idsToRemove) {
					temp = append(temp, id)
				}
			}
			timeTable[a][time] = temp
		}
	}

	return timeTable
}

// TableIdsEach checks if the time table is empty
func TableIdsEach(timeTable map[string]map[rune][]uint, test func([]uint) bool) bool {
	b := true
	for _, day := range timeTable {
		for _, ids := range day {
			if test(ids) {
				b = false
			}
		}
	}
	return b
}

// SliceUniqMap returns unique slice of ids
func SliceUniqMap(s []uint) []uint {
	seen := make(map[uint]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

// LeftFromTable returns the ids that are left in the time table
func LeftFromTable(timeTable map[string]map[rune][]uint) (ret []uint) {
	for _, day := range timeTable {
		for _, ids := range day {
			for _, id := range ids {
				ret = append(ret, id)
			}
		}
	}
	return SliceUniqMap(ret)
}

// IsNewerTime checks if the first param time is newer than the second
func IsNewerTime(timeOne string, timeTwo string) bool {
	t1, _ := time.Parse(layoutISO, timeOne)
	t2, _ := time.Parse(layoutISO, timeTwo)
	return t1.After(t2)
}

// DaysL converts day strings to rune arr which is a string (e.g: "Monday, Tuesday" => "MT")
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
