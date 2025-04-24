// Package daysteps contains functionality for parsing a string with
// data about walks and forming a string with information about them.
// The DaySteps structure and two exported methods for it are described.
package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps It contains all the necessary data about daytime walks:
// the number of steps, the duration, as well as data from the personaldata.Personal,
// that is, the user's name, weight, and height
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse parses a string with data in the format "678,0h50m"
// and writes the data to the corresponding fields of the DaySteps structure.
func (ds *DaySteps) Parse(datastring string) (err error) {
	params := strings.Split(datastring, ",")
	if len(params) != 2 {
		return errors.New("parsePackage: invalid data")
	}

	steps, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("parsePackage: steps <= 0")
	}

	duration, err := time.ParseDuration(params[1])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("parsePackage: duration <= 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo generates and returns a string with data about the walk
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	info += fmt.Sprintf("Дистанция составила %0.2f км.\n", distance)
	info += fmt.Sprintf("Вы сожгли %0.2f ккал.\n", calories)

	return info, nil
}
