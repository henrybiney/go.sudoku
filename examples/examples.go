package examples

var (
	//this requires speculation
	EX1 = []int{
		0, 3, 0, 0, 0, 0, 0, 4, 0,
		0, 1, 0, 0, 9, 7, 0, 5, 0,
		0, 0, 2, 5, 0, 8, 6, 0, 0,
		0, 0, 3, 0, 0, 0, 8, 0, 0,
		9, 0, 0, 0, 0, 4, 3, 0, 0,
		0, 0, 7, 6, 0, 0, 0, 0, 4,
		0, 0, 9, 8, 0, 5, 4, 0, 0,
		0, 7, 0, 0, 0, 0, 0, 2, 0,
		0, 5, 0, 0, 7, 1, 0, 8, 0}

	//this requires no speculation
	//when finding the solution
	EX2 = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 7,
		7, 0, 4, 0, 0, 0, 8, 9, 3,
		0, 0, 6, 8, 0, 2, 0, 0, 0,
		0, 0, 7, 5, 2, 8, 6, 0, 0,
		0, 8, 0, 0, 0, 6, 7, 0, 1,
		9, 0, 3, 4, 0, 0, 0, 8, 0,
		0, 0, 0, 7, 0, 4, 9, 0, 0,
		6, 0, 0, 0, 9, 0, 0, 0, 0,
		4, 5, 9, 0, 0, 0, 1, 0, 8}

	EX3 = []int{
		0, 0, 0, 0, 0, 7, 0, 0, 0,
		0, 9, 0, 1, 0, 0, 6, 2, 0,
		3, 1, 5, 0, 2, 0, 0, 4, 0,
		0, 0, 0, 3, 0, 0, 0, 0, 7,
		9, 0, 0, 7, 0, 0, 3, 0, 0,
		0, 3, 2, 0, 9, 8, 0, 6, 0,
		5, 0, 0, 2, 0, 0, 8, 0, 0,
		0, 0, 0, 9, 0, 0, 1, 7, 0,
		6, 0, 9, 0, 3, 0, 4, 0, 0}

	//SPECULATION REQUIRED
	EXT4 = []int{
		0, 0, 0, 2, 0, 0, 9, 0, 0,
		0, 0, 0, 0, 5, 0, 0, 8, 0,
		0, 7, 4, 0, 0, 0, 0, 0, 3,
		5, 0, 8, 0, 6, 1, 0, 0, 0,
		0, 4, 0, 8, 0, 9, 0, 2, 0,
		0, 1, 0, 0, 4, 0, 0, 0, 6,
		0, 0, 5, 0, 0, 3, 7, 0, 0,
		3, 0, 0, 6, 9, 0, 1, 0, 0,
		0, 9, 0, 0, 0, 4, 0, 0, 0}

	//NO SPECULATION REQUIRED
	METRO_21_03_18_EASY = []int{
		4, 0, 6, 3, 0, 0, 5, 0, 0,
		0, 0, 0, 4, 5, 0, 0, 6, 0,
		5, 0, 2, 0, 1, 0, 4, 0, 8,
		1, 3, 0, 5, 0, 8, 0, 0, 0,
		0, 6, 4, 0, 0, 0, 1, 5, 0,
		0, 0, 0, 1, 0, 3, 0, 8, 4,
		6, 0, 9, 0, 3, 0, 8, 0, 2,
		0, 1, 0, 0, 2, 4, 0, 0, 0,
		0, 0, 8, 0, 0, 5, 3, 0, 7}

	//REQUIRES SPECULATION
	METRO_21_03_18_MODERATE = []int{
		0, 5, 0, 3, 9, 1, 0, 0, 0,
		3, 7, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 0, 8, 0, 5, 4, 0,
		2, 0, 0, 1, 0, 9, 0, 0, 8,
		5, 0, 1, 0, 0, 0, 4, 0, 7,
		6, 0, 0, 8, 0, 4, 0, 0, 1,
		0, 2, 9, 0, 1, 0, 0, 0, 0,
		0, 0, 3, 0, 0, 0, 0, 1, 5,
		0, 0, 0, 6, 3, 8, 0, 7, 0}

	BAD1 = []int{0, 0, 0, 2, 1, 2, 2, 0, 88}
)