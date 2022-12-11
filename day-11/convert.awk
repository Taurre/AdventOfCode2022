BEGIN {
	FS = "[ =,:][ =,:]*"
}

/Monkey/ {
	printf("\t\tmonkey{\n")
}

/Starting/ {
	printf("\t\t\titems: []int{")
	
	for (i = 4; i <= NF; i++) {
		if (i < NF) {
			printf("%d, ", $i)
		} else {
			printf("%d", $i)
		}
	}

	printf("},\n")
}

/Operation/ {
	printf("\t\t\toperation: operation{operator: '%c', value: %d},\n", $5, $6)
}

/Test/ {
	printf("\t\t\tdivisible: %d,\n", $NF)
}

/If true/ {
	printf("\t\t\tifTrue: %d,\n", $NF)
}

/If false/ {
	printf("\t\t\tifFalse: %d,\n", $NF)
}

/^$/ {
	printf("\t\t\tinspect: 0,\n")
	printf("\t\t},\n")
}

END {
	printf("\t\t\tinspect: 0,\n")
	printf("\t\t},\n")
}
