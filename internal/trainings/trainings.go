// Package trainings contains functionality for parsing a row with training data
// and forming a row with information about them. There are two methods for this:
// the first is for parsing the incoming data string;
// and the second is for generating training information.
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training contains all the necessary training data:
// the number of steps, the type of workout, the duration of the workout,
// as well as data from the personaldata.Personal
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse parses the data string and writes the data
// to the corresponding fields of the Training structure
func (t *Training) Parse(datastring string) (err error) {
	params := strings.Split(datastring, ",")
	if len(params) != 3 {
		return errors.New("parseTraining: invalid data")
	}

	steps, err := strconv.Atoi(params[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("parseTraining: steps <= 0")
	}

	trainingType := params[1]
	if trainingType == "" {
		return errors.New("parseTraining: typeActivity is empty")
	}

	duration, err := time.ParseDuration(params[2])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("parseTraining: duration <= 0")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration

	return nil
}

// ActionInfo collects information about the activity
// and returns it as a formatted string
func (t Training) ActionInfo() (string, error) {
	calories := 0.0
	var err error
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}

	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	info := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	info += fmt.Sprintf("Длительность: %0.2f ч.\n", t.Duration.Hours())
	info += fmt.Sprintf("Дистанция: %0.2f км.\n", distance)
	info += fmt.Sprintf("Скорость: %0.2f км/ч\n", meanSpeed)
	info += fmt.Sprintf("Сожгли калорий: %0.2f\n", calories)

	return info, nil
}
