package chess

import (
	"errors"
	"math"
	"regexp"
)

func CanKnightAttack(white, black string) (bool, error) {
	if !validate(white) || !validate(black) || white == black {
		return false, errors.New("not valid")
	}

	if (math.Abs(float64(white[0])-float64(black[0])) == 1 && math.Abs(float64(white[1])-float64(black[1])) == 2) ||
		(math.Abs(float64(white[0])-float64(black[0])) == 2 && math.Abs(float64(white[1])-float64(black[1])) == 1) {
		return true, nil
	}
	return false, nil
}

func validate(pos string) bool {
	match, _ := regexp.MatchString("[a-h][1-8]", pos)
	return match
}
