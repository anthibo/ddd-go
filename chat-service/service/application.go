package service

import (
	"chat-service/app"
	"context"

	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
// TODO: inject repositories here and domain factories here

	// factoryConfig := hour.FactoryConfig{
	// 	MaxWeeksInTheFutureToSet: 6,
	// 	MinUtcHour:               12,
	// 	MaxUtcHour:               20,
	// }

	// datesRepository := adapters.NewDatesFirestoreRepository(firestoreClient, factoryConfig)

	// hourFactory, err := hour.NewFactory(factoryConfig)
	// if err != nil {
	// 	panic(err)
	// }

	// hourRepository := adapters.NewFirestoreHourRepository(firestoreClient, hourFactory)

	logger := logrus.NewEntry(logrus.StandardLogger())

	// inject repos and loggers in commands and queries
	return app.Application{
		Commands: app.Commands{
			// CancelTraining:       command.NewCancelTrainingHandler(hourRepository, logger, metricsClient),
			// ScheduleTraining:     command.NewScheduleTrainingHandler(hourRepository, logger, metricsClient),
			// MakeHoursAvailable:   command.NewMakeHoursAvailableHandler(hourRepository, logger, metricsClient),
			// MakeHoursUnavailabl`	e: command.NewMakeHoursUnavailableHandler(hourRepository, logger, metricsClient),
		},
		Queries: app.Queries{
			// HourAvailability:      query.NewHourAvailabilityHandler(hourRepository, logger, metricsClient),
			// TrainerAvailableHours: query.NewAvailableHoursHandler(datesRepository, logger, metricsClient),
		},
	}
}
