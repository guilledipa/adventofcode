package dial

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	dialNumbers            = 100 // 0..99
	dialInstructionPattern = `([A-Z])(\d+)`
)

func UseDial(rotations string, startingPosition int) int {
	var password int
	s, f, err := utils.CreateScanner(rotations)
	if err != nil {
		log.Fatalf("Unable to read input data: %q", err)
	}
	defer f.Close()
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// Do stuff
	return password
}

func execInstruction(instruction string, dialPosition int) (int, int, error) {
	var deltaPosition int
	direction, clicks, err := parseInstruction(instruction)
	if err != nil {
		return dialPosition, err
	}
	// Unwrapping the dial into an array it would look like this:
	// dial := [0..99] -> Len(dial) = 100
	// Direction
	//   L -> Cursor moves left in the array -> Negative index
	//   R -> Cursor moves right in the array -> Positive index
	if direction == "L" {
		deltaPosition = dialPosition - (-clicks)
	} else if direction == "R" {
		deltaPosition = dialPosition - clicks
	} else {
		return dialPosition, 0, fmt.Errorf("impossible direction: %q", direction)
	}
	zeroCount := deltaPosition / dialNumbers
	newDialPosition := deltaPosition % dialNumbers
	return newDialPosition, zeroCount, nil
}

func parseInstruction(instruction string) (string, int, error) {
	re := regexp.MustCompile(`([A-Z])(\d+)`)
	if !re.Match([]byte(instruction)) {
		return "", 0, fmt.Errorf(fmt.Sprintf("unexpected instruction: %q", instruction))
	}
	matches := re.FindStringSubmatch(instruction)[1:] // ["R120", "R", "120"] -> ["R", "120"]
	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", 0, err
	}
	return matches[0], i, nil
}
