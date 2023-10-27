package lotrsdk

type Movie struct {
	ID                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           int     `json:"runtimeInMinutes"`
	BudgetInMillions           float64 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float64 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float64 `json:"rottenTomatoesScore"`
}

type MoviesResponse struct {
	Docs []Movie `json:"docs"`
}

type Quote struct {
	ID        string `json:"_id"`
	Dialog    string `json:"dialog"`
	MovieID   string `json:"movie"`
	Character string `json:"character"`
	QuoteID   string `json:"id"`
}

type QuotesResponse struct {
	Docs []Quote `json:"docs"`
}

type Documenter interface {
	GetDocs() []interface{}
}

func (m MoviesResponse) GetDocs() []interface{} {
	docs := make([]interface{}, len(m.Docs))
	for i, v := range m.Docs {
		docs[i] = v
	}
	return docs
}

func (q QuotesResponse) GetDocs() []interface{} {
	docs := make([]interface{}, len(q.Docs))
	for i, v := range q.Docs {
		docs[i] = v
	}
	return docs
}
