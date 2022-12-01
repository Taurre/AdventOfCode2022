#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct elve {
	size_t calories;
	struct elve *next;
};

static void *
memdup(void *data, size_t size) {
	assert(data != NULL);
	assert(size > 0);
	void *copy = malloc(size);

	if (copy == NULL) {
		perror("malloc");
		exit(EXIT_FAILURE);
	}

	memcpy(copy, data, size);
	return copy;
}

static void
elve_append_sorted(struct elve *elves, struct elve *elve) {
	assert(elves != NULL);
	assert(elve != NULL);
	struct elve *e = elves;

	for (; e->next != NULL; e = e->next) {
		if (elve->calories > e->next->calories) {
			elve->next = e->next;
			e->next = elve;
			return;
		}
	}

	e->next = elve;
	return;
}

static size_t
elve_calories_sum(struct elve *elves, size_t count) {
	assert(elves != NULL);
	assert(count > 0);
	size_t total = 0;
	struct elve *elve = elves->next;

	for (size_t i = 0; i < count; i++) {
		assert(elve != NULL);
		total += elve->calories;
		elve = elve->next;
	}

	return total;
}

static size_t
read_calories(FILE *fp) {
	char line[255];
	size_t total = 0;

	assert(fp != NULL);

	while (fgets(line, sizeof line, fp) != NULL) {
		size_t calories = 0;

		if (line[0] == '\n')
			return total;
		if (sscanf(line, "%zu", &calories) != 1) {
			perror("fscanf");
			exit(EXIT_FAILURE);
		}

		total += calories;
	}

	if (ferror(fp)) {
		perror("fgets");
		exit(EXIT_FAILURE);
	}

	return total ? total : (size_t)-1;
}

int
main(int argc, char *argv[]) {
	if (argc == 1) {
		fprintf(stderr, "Usage: %s file\n", argv[0]);
		exit(EXIT_FAILURE);
	}

	FILE *fp = fopen(argv[1], "r");

	if (fp == NULL) {
		perror("fopen");
		exit(EXIT_FAILURE);
	}

	struct elve *elves = &(struct elve){ .next = NULL };
	size_t calories = 0;

	while ((calories = read_calories(fp)) != (size_t)-1) {
		struct elve *elve = memdup(&(struct elve) { .calories = calories }, sizeof *elve);
		elve_append_sorted(elves, elve);
	}

	printf("Max of calories: %zu\n", elve_calories_sum(elves, 1));
	printf("Sum of top three calories: %zu\n", elve_calories_sum(elves, 3));
	return 0;
}
