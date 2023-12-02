package userdm

type Career struct {
	id     CareerID
	detail string
}

func newCareer(id CareerID, detail string) (*Career, error) {
	// validationチェックが入る
	return &Career{
		id:     id,
		detail: detail,
	}, nil
}
