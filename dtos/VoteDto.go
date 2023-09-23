package dtos

type VoteRequest struct {
	PaslonID  int    `json:"paslon_id" form:"paslon_id" validate:"required"`
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
}

type VoteResponse struct {
	ID        int                  `json:"id"`
	VoterName string               `json:"voter_name"`
	Paslon    PaslonResponseToVote `json:"paslon"`
}
