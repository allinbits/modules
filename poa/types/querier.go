package types

const (
	QueryValidatorPOA     = "validator-poa"
	QueryAllValidatorsPOA = "validators"
	QueryVotePOA          = "vote-poa"
	QueryAllVotesPOA      = "votes"
)

// QueryValidatorResolve Queries Result Payload for a resolve query
type QueryValidatorResolve struct {
	Name     string `json:"name"`
	Accepted bool   `json:"accepted"`
}

// QueryValidatorParams
type QueryValidatorParams struct {
	Name string
}

// NewQueryValidatorParams
func NewQueryValidatorParams(name string) QueryValidatorParams {
	return QueryValidatorParams{
		Name: name,
	}
}

// QueryVoteResolve Queries Result Payload for a resolve query
type QueryVoteResolve struct {
	Name    string `json:"name"`
	Voter   string `json:"voter"`
	InFavor bool   `json:"in_favor"`
}

// QueryVoteParams
type QueryVoteParams struct {
	Name  string
	Voter string
}

// NewQueryVoteParams
func NewQueryVoteParams(name string, voter string) QueryVoteParams {
	return QueryVoteParams{
		Name:  name,
		Voter: voter,
	}
}
