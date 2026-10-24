package main

// Approach: convert each "MM-DD" date to its day of the year using a prefix
// sum of month lengths, then compute the length of the intersection of the
// two inclusive intervals: max(0, min(leaves) - max(arrivals) + 1).

var daysInMonth = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func dayOfYear(date string) int {
	month := int(date[0]-'0')*10 + int(date[1]-'0')
	day := int(date[3]-'0')*10 + int(date[4]-'0')
	total := day
	for i := 0; i < month-1; i++ {
		total += daysInMonth[i]
	}
	return total
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	start := max(dayOfYear(arriveAlice), dayOfYear(arriveBob))
	end := min(dayOfYear(leaveAlice), dayOfYear(leaveBob))
	if end < start {
		return 0
	}
	return end - start + 1
}
