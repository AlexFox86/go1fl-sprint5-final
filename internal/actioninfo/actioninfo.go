// Package actioninfo implements the output of general information about all workouts and walks.
// The package has created an interface that will contain two methods: Parse() and Action Info(),
// and a function to which slices with data strings and
// instances of Training and DaySteps structures will be passed.
package actioninfo

import "log"

// DataParser contains two methods: Parse() and ActionInfo()
// These methods are implemented in the training and
// daysteps packages for their Training and DaySteps types, respectively.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info accepts a slice of rows with training or walking data and
// an instance of one of the Training or DaySteps structures.
// Generates and outputs a string with information about the activity
func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		if err := dp.Parse(value); err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}

		log.Println(info)
	}
}
