package datetime

import "time"

type Date time.Time

const (
	formatDate = "2006-01-02"
)

func (t Date) String() string {
	return time.Time(t).Format(formatDate)
}

func (t Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(formatDate)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, formatDate)
	b = append(b, '"')
	return b, nil
}

func (t *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatDate+`"`, string(data), time.Local)
	*t = Date(now)
	return
}
