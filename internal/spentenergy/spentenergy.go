// Package spentenergy contains functions for calculating calories burned while walking, running,
// as well as functions for calculating average speed and distance.
// The functions of this package will be used in other packages, so they must be exportable.
package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories calculates the number of calories
// spent while walking and returns the result
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("WalkingSpentCalories: steps <= 0")
	}

	if duration <= 0 {
		return 0, errors.New("WalkingSpentCalories: duration <= 0")
	}

	if weight <= 0 {
		return 0, errors.New("WalkingSpentCalories: weight <= 0")
	}

	if height <= 0 {
		return 0, errors.New("WalkingSpentCalories: height <= 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	return ((weight * meanSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories calculates the number of calories
// spent while running and returns the result
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("RunningSpentCalories: steps <= 0")
	}

	if duration <= 0 {
		return 0, errors.New("RunningSpentCalories: duration <= 0")
	}

	if weight <= 0 {
		return 0, errors.New("RunningSpentCalories: weight <= 0")
	}

	if height <= 0 {
		return 0, errors.New("RunningSpentCalories: height <= 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	return (weight * meanSpeed * duration.Minutes()) / minInH, nil
}

// MeanSpeed takes the number of steps, user height
// and duration of activity and returns the average speed.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / float64(duration.Hours())
}

// Distance takes the number of steps and the height of the user in meters,
// and returns the distance in kilometers
func Distance(steps int, height float64) float64 {
	return float64(steps) * (height * stepLengthCoefficient) / float64(mInKm)
}
