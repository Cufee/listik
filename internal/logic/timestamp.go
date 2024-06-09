package logic

import (
	"fmt"
	"time"
)

func FormatTimestamp(timestamp time.Time) string {
	if timestamp.IsZero() {
		return ""
	}

	diff := time.Since(timestamp)
	if diff < time.Minute {
		return "just now"
	} else if diff < time.Hour {
		minutes := int(diff.Minutes())
		return fmt.Sprintf("%d minute%s ago", minutes, StringIfElse(minutes != 1, "s", ""))
	} else if diff < time.Hour*24 {
		hours := int(diff.Hours())
		return fmt.Sprintf("%d hour%s ago", hours, StringIfElse(hours != 1, "s", ""))
	} else if diff < time.Hour*24*7 {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%d day%s ago", days, StringIfElse(days != 1, "s", ""))
	} else if diff < time.Hour*24*30 {
		weeks := int(diff.Hours() / 24 / 7)
		return fmt.Sprintf("%d week%s ago", weeks, StringIfElse(weeks != 1, "s", ""))
	} else if diff < time.Hour*24*365 {
		months := int(diff.Hours() / 24 / 30)
		return fmt.Sprintf("%d month%s ago", months, StringIfElse(months != 1, "s", ""))
	} else {
		years := int(diff.Hours() / 24 / 365)
		return fmt.Sprintf("%d year%s ago", years, StringIfElse(years != 1, "s", ""))
	}
}
