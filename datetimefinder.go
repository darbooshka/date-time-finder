package datematchgo

import (
	"time"
	"regexp"
	"strings"
)


type DateTimeFinderComponent interface {
	FindDateTime(text string) []time.Time
}


type DateTimeFinder struct {
	components []DateTimeFinderComponent
}
func (d *DateTimeFinder) FindDateTime(text string) []time.Time {
	textn := replaceUkrainianMonths(
		replaceRussianMonths(text))

	uniqueTimes := make(map[time.Time]bool)
	var allTimes []time.Time

	// Iterate over each component and collect the time arrays
	for _, component := range d.components {
		times := component.FindDateTime(textn)

		// allTimes = append(allTimes, times...)
		// Add unique times to the map
		for _, t := range times {
			if !uniqueTimes[t] {
				uniqueTimes[t] = true
				allTimes = append(allTimes, t)
			}
		}
	}

	return allTimes
}
func NewDateTimeFinder() DateTimeFinderComponent {
	a := &patternDateTimeFinder {
		regexPattern: `\d{1,2}/\d{1,2}/\d{4} \d{1,2}:\d{2} [AP]M`,
		parsePattern: "01/02/2006 3:04 PM",
	}
	b := &patternDateTimeFinder {
		regexPattern: `\d{4}/\d{2}/\d{2} \d{1,2}:\d{2} (?:AM|PM)`,
		parsePattern: "2006/01/02 3:04 PM",
	}
	c := &patternDateTimeFinder {
		regexPattern: `\d{2} \p{L}+ \d{4} \d{1,2}:\d{2} (?:AM|PM)`,
		parsePattern: "02 Jan 2006 03:04 PM",
	}
	d := &patternDateTimeFinder {
		regexPattern: `\d{2} \p{L}+ \d{4} \d{1,2}:\d{2} (?:AM|PM)`,
		parsePattern: "02 January 2006 03:04 PM",
	}
	
	return &DateTimeFinder{
		components: []DateTimeFinderComponent{ a, b, c, d, },
	}
}

type patternDateTimeFinder struct {
	regexPattern string
	parsePattern string
}
func (p *patternDateTimeFinder) FindDateTime(text string) []time.Time {
	// Regular expression pattern A to find date and time in the given text
	pattern := p.regexPattern
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)

	// Parse matched strings into time.Time objects
	var times []time.Time
	for _, match := range matches {
		dt, err := time.Parse(p.parsePattern, match)
		if err == nil {
			times = append(times, dt)
		}
	}

	return times
}

func replaceUkrainianMonths(dateString string) string {
	months := map[string]string{
		"січня":   "January",
		"лютого":  "February",
		"березня": "March",
		"квітня":  "April",
		"травня":  "May",
		"червня":  "June",
		"липня":   "July",
		"серпня":  "August",
		"вересня": "September",
		"жовтня":  "October",
		"листопада": "November",
		"грудня":  "December",
	}

	for ukrMonth, engMonth := range months {
		dateString = strings.ReplaceAll(dateString, ukrMonth, engMonth)
	}

	return dateString
}

func replaceRussianMonths(dateString string) string {
	months := map[string]string{
		"января":   "January",
		"февраля":  "February",
		"марта": "March",
		"апреля":  "April",
		"мая":  "May",
		"июня":  "June",
		"июля":   "July",
		"августа":  "August",
		"сентября": "September",
		"октября":  "October",
		"ноября": "November",
		"декабря":  "December",
	}

	for ukrMonth, engMonth := range months {
		dateString = strings.ReplaceAll(dateString, ukrMonth, engMonth)
	}

	return dateString
}
