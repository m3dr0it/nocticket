package entity

import (
	"time"
)

var (
	HighSLATime   = 4 * time.Hour
	MediumSLATime = 12 * time.Hour
	LowSLATime    = 24 * time.Hour
)

type Priority string

type PriorityGroup []Priority

const (
	High   Priority = "high"
	Medium Priority = "medium"
	Low    Priority = "low"
)

type SLATime time.Duration

func (p Priority) GetSLATime() time.Duration {
	switch p {
	case High:
		return HighSLATime
	case Medium:
		return MediumSLATime
	default:
		return LowSLATime
	}
}
