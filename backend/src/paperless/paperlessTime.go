package paperless

import (
	"errors"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type PaperlessTime struct {
	time.Time
}

func NewPaperlessTime(t time.Time) *PaperlessTime {
	p := &PaperlessTime{}
	p.Time = t

	return p
}

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

func MarshalPaperlessTime(t PaperlessTime) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, t.Time.Format(`"2006-01-02"`))
	})
}

func UnmarshalPaperlessTime(v interface{}) (PaperlessTime, error) {
	var err error
	var date time.Time
	if tmpStr, ok := v.(string); ok {
		date, err = time.Parse(`"2006-01-02"`, tmpStr)
		if err != nil {
			return PaperlessTime{}, err
		}
	} else if tmpStr, ok := v.([]byte); ok {
		date, err = time.Parse(`"2006-01-02"`, string(tmpStr))
		if err != nil {
			return PaperlessTime{}, err
		}
	} else {
		return PaperlessTime{}, errors.New("type is invalid")
	}

	return *NewPaperlessTime(date), nil
}
