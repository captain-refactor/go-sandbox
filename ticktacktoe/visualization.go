package ticktacktoe

import (
	"fmt"
	"strings"
)

func visualizeCell(value Side, x, y int) string {
	switch value {
	case Empty:
		return fmt.Sprintf("[%d%d]", x, y)
	case Circle:
		return " O  "
	case Cross:
		return " X  "
	}
	return value.String()
}

func (s State) String() string {
	var result []string
	result = append(result, "\n")
	for x := 0; x < width; x++ {
		result = append(result, "  ")
		for y := 0; y < height; y++ {
			result = append(result, visualizeCell(s.board[x][y], x, y), " ")
		}
		result = append(result, "\n")
	}
	return strings.Join(result, "")
}
