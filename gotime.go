package gotime

import (
	"regexp"
	"time"
)

var fmtmap = make(map[*regexp.Regexp]string)

func setValues() {
	if len(fmtmap) == 0 {
		// Full month name
		fmtmap[regexp.MustCompile("%B")] = "January"
		// Short month name
		fmtmap[regexp.MustCompile("%b")] = "Jan"
		// Month as zero-padded decimal
		fmtmap[regexp.MustCompile("%m")] = "01"
		// Full day name
		fmtmap[regexp.MustCompile("%A")] = "Monday"
		// Short day name
		fmtmap[regexp.MustCompile("%a")] = "Mon"
		// Day of month as zero-padded decimal
		fmtmap[regexp.MustCompile("%d")] = "02"
		// Hour (12-hour clock) as zero-padded decimal
		fmtmap[regexp.MustCompile("%I")] = "03"
		// Hour (24-hour clock) as zero-padded decimal
		fmtmap[regexp.MustCompile("%H")] = "15"
		// Minute as zero-padded decimal
		fmtmap[regexp.MustCompile("%M")] = "04"
		// Second as zero-padded decimal
		fmtmap[regexp.MustCompile("%S")] = "05"
		// Year with century as decimal
		fmtmap[regexp.MustCompile("%Y")] = "2006"
		// Year without century as zero-padded decimal
		fmtmap[regexp.MustCompile("%y")] = "06"
		// AM or PM, or locale equivilent
		fmtmap[regexp.MustCompile("%p")] = "pm"
		// Timezone
		fmtmap[regexp.MustCompile("%Z")] = "MST"
	}
}

// FormatDateTime uses standard verbs.
// Verbs available are:
//   %B : Full month name
//   %b : Short month name
//   %m : Month as zero-padded decimal
//   %A : Full day name
//   %a : Short day name
//   %d : Day of month as zero-padded decimal
//   %I : Hour (12-hour clock) as zero-padded decimal
//   %H : Hour (24-hour clock) as zero-padded decimal
//   %M : Minute as zero-padded decimal
//   %S : Second as zero-padded decimal
//   %Y : Year with century as decimal
//   %y : Year without century as zero-padded decimal
//   %p : AM or PM, or locale equivilent
//   %Z : Timezone
// Verbs were taken from the Python3 documentation: https://docs.python.org/3/library/datetime.html#strftime-and-strptime-behavior
func FormatDateTime(DateTime time.Time, fmt string) string {
	setValues()
	for re, val := range fmtmap {
		fmt = re.ReplaceAllString(fmt, val)
	}
	return DateTime.Format(fmt)
}
