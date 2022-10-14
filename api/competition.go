package api

import (
	"errors"
)

type CompetitionStatus string

const (
	Upcoming   CompetitionStatus = "Upcoming"
	Inprogress CompetitionStatus = "Inprogress"
	Completed  CompetitionStatus = "Completed"
)

func (st CompetitionStatus) IsValid() error {
	switch st {
	case Upcoming, Inprogress, Completed:
		return nil
	}
	return errors.New("invalid competition status")
}
