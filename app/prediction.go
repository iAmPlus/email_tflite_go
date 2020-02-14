package app
package nlp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"bufio"
	"dialogmgr/core"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	architecturePath     = "/model"
	intentsBasePath      = architecturePath + "/classifier"
	word2Idx             = architecturePath + "/word2Idx.npy"
	mainIntentModelPath  = intentsBasePath + "/email_classifier/model.tflite"
	mainIntentLabelsPath = intentsBasePath + "/email_classifier/idx2Label.npy"
	minProbability       = 0.4
	maxSentenceLength    = int(100)
)

var case2Idx = map[string]int{
	"numeric":        0,
	"allLower":       1,
	"allUpper":       2,
	"initialUpper":   3,
	"other":          4,
	"mainly_numeric": 5,
	"contains_digit": 6,
	"PADDING_TOKEN":  7,
}

type localNlp struct {
	interpreter *tflite.Interpreter
	idx2Label   map[int]string
	word2Idx    map[string]int
	assetsPath  string
}
// SkillRequest represents a valid request for the skill
type RequestBody struct {
	Query   string `json:"query"`
	Content string `json:"content"`
}

type ResponseBody struct {
	Matches []string `json:"matches"`
}

func (a *Application) addPredictionRoute() {
	a.router.HandleFunc("/predict", a.Predict)
	a.router.HandleFunc("/classify", a.Classify)
}
'''
// Predict the answer from a given content
func (a *Application) Predict(w http.ResponseWriter, r *http.Request) {
	input := RequestBody{}
	w.Header().Add("Content-Type", "application/json")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return
	}

	err = json.Unmarshal(data, &input)
	result, err := a.predictor.Predict(input.Query, input.Content)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(result) <= 0 {
		http.Error(w, "no matches found", http.StatusNotFound)
		return
	}
	response := ResponseBody{Matches: result}
	body, _ := json.Marshal(response)

	_, err = w.Write(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
'''

func (a *Application) Classify(r ClassifierRequest) ([]ClassifierResult, error) {

	intentModel := tflite.NewModelFromFile(assetsPath + mainIntentModelPath)
	if intentModel == nil {
		fmt.Println("Exception : Cannot load model")
	}
	word2Idx, _ := loadWord2Idx(assetsPath + word2Idx)
	idx2Label, _ := loadIdx2Label(assetsPath + mainIntentLabelsPath)

	interpreter := tflite.NewInterpreter(intentModel, nil)


	tokens := r.BasicNlpRequest.Tokens

	wordIndices := createTensorForClassification(tokens, word2Idx)

	tokensData := padSequences1DFloat32(wordIndices, maxSentenceLength)

	interpreter.AllocateTensors()

	input := interpreter.GetInputTensor(0)
	inputType := input.Type()

	if inputType == tflite.Float32 {
		copy(input.Float32s(), tokensData)
	}

	status := interpreter.Invoke()
	if status != tflite.OK {
		fmt.Println("Exception : Classification failed")
	}

	output := interpreter.GetOutputTensor(0)
	intent, score := getIntent(output.Float32s(), idx2Label)

	interpreter.ResetVariableTensors()

	var result []ClassifierResult
	result = append(result, ClassifierResult{
		Score:     float64(score),
		Algorithm: "lstm",
		Intent:    intent,
		Language:  "en",
	})
	return result, nil
}


func createTensorForClassification(sentence []string, word2Idx map[string]int) []int {
	/*
	* input : sentence = ["play", "hello", etc...]
	* output : wordIndices []int
	 */
	unknownIdx := word2Idx["UNKNOWN_TOKEN"]

	wordIndices := make([]int, 0)
	wordIdx := 0
	for _, word := range sentence {
		if val, ok := word2Idx[word]; ok {
			wordIdx = val
		} else if val, ok := word2Idx[strings.ToLower(word)]; ok {
			wordIdx = val
		} else {
			wordIdx = unknownIdx
		}

		wordIndices = append(wordIndices, wordIdx)
	}
	return wordIndices
}

func padSequences1DFloat32(array []int, length int) []float32 {
	returnArray := []float32{}
	for _, value := range array {
		returnArray = append(returnArray, float32(value))
	}

	for index := 0; index < length-len(array); index++ {
		returnArray = append(returnArray, float32(0))
	}
	return returnArray
}

func (n *localNlp) getIntent(probabs []float32, labels map[int]string) (string, float32) {
	var indexFirst int
	first := probabs[0]
	for i, prob := range probabs {
		if prob > minProbability {
			if prob >= first {
				first = prob
				indexFirst = i
			}
		}
	}
	return labels[indexFirst], first
}