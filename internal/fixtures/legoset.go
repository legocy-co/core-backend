package fixtures

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	models "legocy-go/internal/domain/lego/models"
	repository2 "legocy-go/internal/domain/lego/repository"
	"os"
	"strconv"
)

const (
	legoSetsJSONFilepath = "/static/json/lego_sets.json"
)

type legoSets []legoSet

type legoSet struct {
	SetNumber string `json:"set_number"`
	Title     string `json:"title"`
	Series    string `json:"series"`
	NDetails  int    `json:"n_details"`
}

func (l legoSet) toLegoSetValueObject(
	r repository2.LegoSeriesRepository) *models.LegoSetValueObject {

	series, err := r.GetLegoSeriesByName(context.Background(), l.Series)
	if err != nil {
		return nil
	}

	setNumber, _err := strconv.Atoi(l.SetNumber)
	if _err != nil {
		return nil
	}

	return &models.LegoSetValueObject{
		Number:   setNumber,
		Name:     l.Title,
		NPieces:  l.NDetails,
		SeriesID: series.ID,
	}
}

func LoadLegoSets(
	setsRepo repository2.LegoSetRepository,
	seriesRepo repository2.LegoSeriesRepository) {

	cwd, _ := os.Getwd()
	jsonFile, err := os.Open(cwd + legoSetsJSONFilepath)
	if err != nil {
		logrus.Errorf("Error opening file %v", err)
		return
	}

	defer jsonFile.Close()

	var setsList legoSets
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err = json.Unmarshal(byteValue, &setsList); err != nil {
		logrus.Errorf("Error unmarshalling %v", err)
		return
	}

	for _, set := range setsList {
		setCreate := set.toLegoSetValueObject(seriesRepo)
		if setCreate == nil {
			continue
		}

		if err = setsRepo.CreateLegoSet(context.Background(), setCreate); err != nil {
			logrus.Errorf("Error creating LegoSet %v", err)
			continue
		}
		logrus.Infof("Successfully created LegoSet %v", setCreate.Number)
	}
}
