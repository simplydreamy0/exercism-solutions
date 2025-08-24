package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0;
	for _, dailyCount := range birdsPerDay {
		total += dailyCount;
	}
	return total;
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekIndex := (week - 1) * 7;
	total := 0;
	for _, dailyCount := range birdsPerDay[weekIndex:weekIndex+7] {
		total += dailyCount;
	}
	return total;
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	fixedSlice := birdsPerDay
	for index, dailyCount := range fixedSlice {
		if index % 2 == 0 {
			fixedSlice[index] = dailyCount + 1;
		}
	}
	return fixedSlice;
}
