package fixtures

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	models "legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/lego/repository"
	"os"
)

const (
	legoSeriesJSONFilepath = "/static/json/lego_themes_final.json"
)

type legoSeriesList []legoSeries

type legoSeries struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (l legoSeries) toLegoSeriesValueObject() *models.LegoSeriesValueObject {
	return &models.LegoSeriesValueObject{
		Name: l.Name,
	}
}

func LoadLegoSeries(r repository.LegoSeriesRepository) {
	cwd, _ := os.Getwd()
	jsonFile, err := os.Open(cwd + legoSeriesJSONFilepath)
	if err != nil {
		logrus.Errorf("Error opening file %v", err)
		return
	}

	defer jsonFile.Close()

	var seriesList legoSeriesList
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err = json.Unmarshal(byteValue, &seriesList); err != nil {
		logrus.Errorf("Error unmarshalling data %v", err)
		return
	}

	for _, series := range seriesList {
		err = r.CreateLegoSeries(context.Background(), series.toLegoSeriesValueObject())
		if err != nil {
			logrus.Errorf("Error creating LegoSeries %v", err)
			continue
		}
		logrus.Infof("Created LegoSeries object %v", series.Name)
	}
}
