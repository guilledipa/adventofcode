package tools

import (
	"adventofcode/2023/day/1/tools"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type Draw struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Id    int
	Draws []Draw
}

type Games struct {
	Games []Game
}

func NewGames(filePath string) (*Games, error) {
	games := new(Games)
	lines, err := tools.ReadInputFile(filePath)
	if err != nil {
		return nil, err
	}
	for _, l := range lines {
		var game Game
		var draws []Draw
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		gameSplit := strings.Split(l, ":") // ["Game 1",  " 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"]
		// Game
		gameID := strings.Split(gameSplit[0], " ")
		id, err := strconv.Atoi(gameID[1])
		if err != nil {
			return nil, err
		}
		game.Id = id
		// Draws
		for _, ds := range strings.Split(gameSplit[1], ";") { // [ " 3 blue, 4 red", " 1 red, 2 green, 6 blue", " 2 green"]
			fmt.Println(ds)
			var draw Draw
			for _, d := range strings.Split(ds, ",") { // [" 3 blue", "4 red"]
				cubesColor := strings.Split(strings.TrimSpace(d), " ") // Careful with spaces.
				cubes, err := strconv.Atoi(cubesColor[0])
				if err != nil {
					return nil, fmt.Errorf("issue converting cubes str to int: %v", err)
				}
				switch cubesColor[1] {
				case "red":
					draw.Red = cubes
				case "green":
					draw.Green = cubes
				case "blue":
					draw.Blue = cubes
				default:
					return nil, fmt.Errorf("unsupported color: %s", cubesColor[1])
				}
			}
			draws = append(draws, draw)
		}
		game.Draws = draws
		games.Games = append(games.Games, game)
	}
	return games, nil
}
