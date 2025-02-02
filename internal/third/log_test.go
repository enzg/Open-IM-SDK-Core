package third

import (
	"testing"
)

func TestLogMatch(t *testing.T) {

	filenames := []string{
		"log1.txt",
		"log2.log",
		"log3.log.txt",
		"log4.log.2022-01-01",
		"log5.log.2022-01-01.txt",
		"log20230918.log",
		"OpenIM.CronTask.log.all.2023-09-18", "OpenIM.log.all.2023-09-18",
	}

	expected := []string{
		"OpenIM.CronTask.log.all.2023-09-18", "OpenIM.log.all.2023-09-18",
	}

	var actual []string
	for _, filename := range filenames {
		if checkLogPath(filename) {
			actual = append(actual, filename)
		}
	}

	if len(actual) != len(expected) {
		t.Errorf("Expected %d matches, but got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Expected match %d to be %q, but got %q", i, expected[i], actual[i])
		}
	}
}
