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
	dialStartingPosition   = 50
	dialInstructionPattern = `([A-Z])(\d+)`
)

type Dial struct {
	Numbers  int
	Position int
	Zeros    int
}

func UseDial(instructions string, countRevs bool) int {
	s, f, err := utils.CreateScanner(instructions)
	if err != nil {
		log.Fatalf("Unable to read input data: %q", err)
	}
	defer f.Close()
	d := Dial{
		Numbers:  dialNumbers,
		Position: dialStartingPosition,
	}
	for s.Scan() {
		instruction := s.Text()
		if err := d.execInstruction(instruction, countRevs); err != nil {
			log.Fatalf("Error executing instruction %q: %v", instruction, err)
		}
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return d.Zeros
}

func (d *Dial) execInstruction(instruction string, countRevs bool) error {
	var deltaPosition int
	direction, clicks, err := parseInstruction(instruction)
	if err != nil {
		return err
	}
	// Unwrapping the dial into an array it would look like this:
	// dial := [0..99] -> Len(dial) = 100
	// Direction
	//   L -> Cursor moves left in the array -> Negative index
	//   R -> Cursor moves right in the array -> Positive index
	switch direction {
	case "L":
		deltaPosition = d.Position - clicks
		log.Printf("L: dialPosition %d, clicks %d, deltaPosition %d", d.Position, clicks, deltaPosition)
	case "R":
		deltaPosition = d.Position + clicks
		log.Printf("R: dialPosition %d, clicks %d, deltaPosition %d", d.Position, clicks, deltaPosition)
	default:
		return fmt.Errorf("impossible direction: %q", direction)
	}
	// The new position is the deltaPosition modulo the number of positions in the dial.
	// This ensures that the position wraps around correctly when it goes below 0 or above 99.
	newPosition := ((deltaPosition % d.Numbers) + d.Numbers) % d.Numbers
	if countRevs {
		d.Zeros += countZeroCrossings(direction, d.Position, clicks, d.Numbers)
	} else if newPosition == 0 && d.Position != 0 {
		d.Zeros++
	}
	d.Position = newPosition // Store the new position for the next instruction.
	return nil
}

func countZeroCrossings(direction string, position, clicks, dialScale int) int {
	switch direction {
	case "R":
		return (position + clicks) / dialScale
	case "L":
		if clicks < position {
			return 0
		}
		if position == 0 {
			return clicks / dialScale
		}
		return (clicks-position)/dialScale + 1
	default:
		return 0
	}
}

func parseInstruction(instruction string) (string, int, error) {
	re := regexp.MustCompile(`([A-Z])(\d+)`)
	if !re.Match([]byte(instruction)) {
		return "", 0, fmt.Errorf("unexpected instruction: %q", instruction)
	}
	matches := re.FindStringSubmatch(instruction)[1:] // ["R120", "R", "120"] -> ["R", "120"]
	direction := matches[0]
	if direction != "L" && direction != "R" {
		return "", 0, fmt.Errorf("unexpected direction: %q", direction)
	}
	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", 0, err
	}
	return direction, i, nil
}
