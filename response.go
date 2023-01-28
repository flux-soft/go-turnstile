package turnstile

import "time"

type Response struct {
	Success     bool      `json:"success"`
	ErrorCodes  []string  `json:"error-codes"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
}

func (r *Response) IsSuccess() bool {
	return r.Success
}

func (r *Response) HasErrors() bool {
	return len(r.ErrorCodes) > 0
}
