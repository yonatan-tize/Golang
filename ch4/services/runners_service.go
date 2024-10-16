package services

import (
	"net/http"
	"project/models"
	"project/repositories"
	"strconv"
	"time"
)

type RunnersService struct {
	runnersRepository *repositories.RunnersRepository
	resultsRepository *repositories.ResultsRepository
}

func NewRunnersService(
	runnersRepository *repositories.RunnersRepository,
	resultsRepository *repositories.ResultsRepository) *RunnersService {
	return &RunnersService{
		runnersRepository: runnersRepository,
		resultsRepository: resultsRepository,
	}
}

func (rs RunnersService) CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError) { 
	responseErr := validateRunner(runner)
	if responseErr != nil {
		return nil, responseErr
	}

	return rs.runnersRepository.CreateRunner(runner)	
}

func (rs RunnersService) UpdateRunner(runnerId string, runner *models.Runner) *models.ResponseError {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil{
		return responseErr
	}
	responseErr = validateRunner(runner)
	if responseErr != nil {
		return responseErr
	}
	return rs.runnersRepository.UpdateRunner(runner)
}

func (rs RunnersService) DeleteRunner(runnerId string) *models.ResponseError {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil{
		return responseErr
	}

	return rs.runnersRepository.DeleteRunner(runnerId)
}

func (rs RunnersService) GetRunner(runnerId string) (*models.Runner, *models.ResponseError) {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil{
		return nil, responseErr
	}

	runner, responseErr := rs.runnersRepository.GetRunner(runnerId)
	if responseErr == nil{
		return nil, responseErr
	} 
	results, responseErr := rs.resultsRepository.GetAllRunnerResults(runnerId)
	if responseErr != nil{
		return nil, responseErr
	}

	runner.Results = results

	return runner, nil
}

func (rs RunnersService) GetRunnersBatch(country string, year string) ([]*models.Runner, *models.ResponseError) {
	if country != "" && year != "" {
		return nil, &models.ResponseError{
			Message: "Only one parameter can be passed",
			Status:  http.StatusBadRequest,
		}
	}

	if country != "" {
		return rs.runnersRepository.GetRunnersByCountry(country)
	}
	if year != "" {
		intYear, err := strconv.Atoi(year)
		if err != nil {
			return nil, &models.ResponseError{
				Message: "Invalid year",
				Status:  http.StatusBadRequest,
			}
		}
		currentYear := time.Now().Year()
		if intYear < 0 || intYear > currentYear {
			return nil, &models.ResponseError{
				Message: "Invalid year",
				Status:  http.StatusBadRequest,
			}
		}
		return rs.runnersRepository.GetRunnersByYear(intYear)
	}
	return rs.runnersRepository.GetAllRunner()
}

func validateRunner(runner *models.Runner) *models.ResponseError{
	//check if first name is empty
	if runner.FirstName == ""{
		return &models.ResponseError{
			Message: "Invalid First Name",
			Status: http.StatusBadRequest,
		}
	}

	// check if last name is empty
	if runner.LastName == ""{
		return &models.ResponseError{
			Message: "Invalid Last Name",
			Status: http.StatusBadRequest,
		}
	}

	// check if age is not between 16 and 125
	if runner.Age < 16 || runner.Age > 125{
		return &models.ResponseError{
			Message: "Invalid age",
			Status: http.StatusBadRequest,
		}
	}

	// check if runners country is empty
	if runner.Country == ""{
		return &models.ResponseError{
			Message: "Invalid Country Name",
			Status: http.StatusBadRequest,
		}
	}

	return nil
}

func validateRunnerId(runnerId string) *models.ResponseError{
	if runnerId == "" {
		return &models.ResponseError{
			Message: "Invalid runner ID",
			Status:  http.StatusBadRequest,
		}
	}

	return nil	
}