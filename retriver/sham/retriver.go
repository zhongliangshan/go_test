package sham

type Retriver struct {
	Contents string
}

func (r Retriver) Get(url string) string {
	return r.Contents
}

