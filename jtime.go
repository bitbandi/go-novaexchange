package novaexchange

import (
	"encoding/json"
	"time"
)

const (
	TIME_FORMAT_SORT = "2006-01-02 15:04"
	TIME_FORMAT_LONG = "2006-01-02 15:04:05"
)

type jTime struct {
	time.Time
}

func (jt *jTime) UnmarshalJSON(data []byte) error {
	var s string
	var format string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if len(s) == len(TIME_FORMAT_SORT) {
		format = TIME_FORMAT_SORT
	} else {
		format = TIME_FORMAT_LONG
	}
	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	jt.Time = t
	return nil
}

func (jt jTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&jt.Time).Format(TIME_FORMAT_LONG))
}
