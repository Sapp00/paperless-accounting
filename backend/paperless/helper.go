package paperless

import "time"

func (t *PaperlessTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (t *PaperlessTime) MarshalJSON() (b []byte, err error) {
	date := t.Time.Format(`"2006-01-02"`)

	return []byte(date), nil
}
