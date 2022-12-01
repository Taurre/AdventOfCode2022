function top_three(calories) {
	if (calories > first) {
		third = second
		second = first
		first = calories
	} else if (calories > second) {
		third = second
		second = calories
	} else if (calories > third) {
		third = calories
	}
}

/[0-9]/ {
	calories += $1	
}

/^$/ {
	top_three(calories)
	calories = 0
}

END {
	top_three(calories)
	print first
	print first+second+third
}
