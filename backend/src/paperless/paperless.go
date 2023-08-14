package paperless

import (
	"errors"
	"sapp/paperless-accounting/config"
)

type Paperless struct {
	conf           *config.Config
	correspondents map[int]*PaperlessCorrespondent
}

func Init(conf *config.Config) (*Paperless, error) {
	p := Paperless{
		conf: conf,
	}

	return &p, nil
}

func (p *Paperless) GetCorrespondent(id int) (*PaperlessCorrespondent, error) {
	if val, ok := p.correspondents[id]; ok {
		return val, nil
	}
	return nil, errors.New("cannot find the correspondent")
}
