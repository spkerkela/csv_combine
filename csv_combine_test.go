package main

func ExampleCombineFiles() {
	fileNames := []string{"tests/001_data.csv", "tests/002_data.csv", "tests/003_data.csv"}
	CombineFiles(fileNames)
	// Output:
	// first,second,third
	// 5,"test 1",21.0
	// 2,"test 2",21.0
	// 3,"test 3",21.0
	// 1,"test 4",21.0
	// 52,"test 5",21.0
	// 21,"test 6",21.0
	// 95,"test 7",21.0
	// 21,"test 9",21.0
	// 12,"test 10",21.0
	// 5,"test 11",21.0
	// 2,"test 12",21.0
	// 3,"test 13",21.0
	// 1,"test 14",21.0
	// 52,"test 15",21.0
	// 21,"test 16",21.0
	// 95,"test 17",21.0
	// 95,"test 18",21.0
	// 21,"test 19",21.0
	// 12,"test 20",21.0
	// 5,"test 21",21.0
	// 2,"test 22",21.0
	// 3,"test 23",21.0
	// 1,"test 24",21.0
	// 52,"test 25",21.0
	// 21,"test 26",21.0
	// 95,"test 27",21.0
	// 95,"test 28",21.0
	// 21,"test 29",21.0
	// 12,"test 30",21.0
}
