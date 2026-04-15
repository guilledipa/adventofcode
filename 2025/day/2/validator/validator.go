package validator

import (
	"adventofcode/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	segmentPattern = `^(\d+)-(\d+)$`
)

func Validate(file string) (int, int, error) {
	log.Printf("Validating file: %s", file)
	data, err := loadData(file)
	if err != nil {
		return 0, 0, err
	}
	var accP1 int
	var accP2 int
	for _, segment := range data {
		p1, p2 := parsePIDSegment(segment)
		accP1 += p1
		accP2 += p2
	}
	return accP1, accP2, nil
}

func loadData(file string) ([]string, error) {
	s, f, err := utils.CreateScanner(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var data []string
	for s.Scan() {
		line := s.Text()
		data = strings.Split(line, ",")
	}
	return data, nil
}

// RE2 no soporta backreferences, así que no puedo usar el patrón `^(\d+)\1$`
// directamente para detectar PIDs inválidos.
// tengo que parsear el PID mas a manopla.
func isInvalidPID(pid string) bool {
	n := len(pid)
	// Si el largo es impar, no puede ser una repetición exacta de dos partes
	if n == 0 || n%2 != 0 {
		return false
	}
	half := n / 2
	// Comparar la primera mitad con la segunda mitad
	return pid[:half] == pid[half:]
}

func isInvalidPID2(pid string) bool {
	n := len(pid)
	if n == 0 {
		return false
	}
	for size := 1; size <= n/2; size++ {
		if n%size != 0 {
			continue
		}
		if strings.Repeat(pid[:size], n/size) == pid {
			return true
		}
	}
	return false
}

func parsePIDSegment(segment string) (int, int) {
	var accP1 int
	var accP2 int
	re := regexp.MustCompile(segmentPattern)
	strings := re.FindStringSubmatch(segment)[1:]
	start, err := strconv.Atoi(strings[0])
	if err != nil {
		log.Printf("Invalid segment start: %s", strings[0])
		return 0, 0
	}
	end, err := strconv.Atoi(strings[1])
	if err != nil {
		log.Printf("Invalid segment end: %s", strings[1])
		return 0, 0
	}
	for i := start; i <= end; i++ {
		pid := strconv.Itoa(i)
		if isInvalidPID(pid) {
			accP1 += i
		}
		if isInvalidPID2(pid) {
			accP2 += i
		}
	}
	return accP1, accP2
}
