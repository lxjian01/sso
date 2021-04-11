package datetime

import "time"

type Datetime time.Time

const (
	formatDatetime = "2006-01-02 15:04:05"
)

func (t Datetime) String() string {
	return time.Time(t).Format(formatDatetime)
}

func (t Datetime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(formatDatetime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, formatDatetime)
	b = append(b, '"')
	return b, nil
}

func (t *Datetime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatDatetime+`"`, string(data), time.Local)
	*t = Datetime(now)
	return
}



