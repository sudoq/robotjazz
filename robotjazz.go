package robotjazz

import (
	"errors"
	"github.com/SudoQ/robotjazz/chord"
	"github.com/SudoQ/robotjazz/model"
)

/*
func main() {
	fmt.Println("Robot Jazz v0.1")
	m := model.New()
	m.Load("resources/chords-v1.csv")
	dataItem, _ := m.Classify([]float64{1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0})

	// Prints top ten most relevant chords
	for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
		fmt.Println(dataItem.ClosestCentroids[i].Tag)
	}
}
*/
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var mainModel *model.Model

func init() {
	mainModel = model.New()
	// TODO read env for root dir of robot jazz project
	mainModel.Load("resources/chords-v1.csv")
}

func GetMatchingChords(notes []float64) ([]*chord.Chord, error) {
	if len(notes) != 12 {
		return nil, errors.New("First input argument must be a float64 slice of length 12")
	}
	dataItem, err := mainModel.Classify(notes)
	if err != nil {
		return nil, err
	}
	topTenChords := make([]*chord.Chord, 0)
	for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
		centroid := dataItem.ClosestCentroids[i]

		name := centroid.Tag
		noteWeights := centroid.Attributes
		chrd := chord.New(name, noteWeights)
		topTenChords = append(topTenChords, chrd)
	}
	return topTenChords, nil
}
