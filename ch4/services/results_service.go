package services

type ResultsService struct{
	resultsRepository *models.ResultsRepository
}

func NewResultsService(resultsRepository Results){
	return &ResultsService{
		resultsRepository: ,
	}
}