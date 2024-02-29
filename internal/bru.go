package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type BruFile struct {
	Summary Summary `json:"summary"`
	Results []Result `json:"result"`
}

type Summary struct {
	TotalRequests int `json:"totalRequests"`
	PassedRequests int `json:"passedRequests"`
	FailedRequests int`json:"failedRequests"`
	TotalAssertions int`json:"totalAssertions"`
	PassedAssertions int`json:"passedAssertions"`
	FailedAssertions int`json:"failedAssertions"`
	TotalTests int`json:"totalTests"`
	PassedTests int`json:"passedTests"`
	FailedTests int`json:"failedTests"`
}

type Result struct {
	Test TestSummary`json:"test"`
	Request  BruRequest`json:"request"`
	Response BruResponse`json:"response"`
	Error string`json:"error"`
	AssertionResults []AssertionResults`json:"assertionResults"`
	TestResults []TestResults`json:"testResults"`
	Runtime float64`json:"runtime"`
	Suitename string`json:"suitename"`
}

type TestSummary struct {
	Filename string `json:"filename"`
}

type BruRequest struct {
	Method string`json:"Method"`
	Url string`json:"url"`
	Data string`json:"data"`
}

type BruResponse struct {
	Status int`json:"status"`
	StatusText string`json:"statusText"`
	Data string`json:"data"`
}
type AssertionResults struct {
	Uid string`json:"uid"`
	LhsExp string`json:"lhsExp"`
	RhsExp string`json:"rhsExp"`
	RhsOperand string`json:"rhsOperand"`
	Operator string`json:"operator"`
	Status string`json:"Status"`
}

type TestResults struct {
	Description string`json:"description"`
	Status string`json:"status"`
	Uid string`json:"uid"`
} 

func NewBruFile(filename string) BruFile{

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to read file " + filename, err)
	}
	var burFile BruFile
	if err := json.Unmarshal(fileBytes, &burFile); err != nil {
		log.Fatal(err)
	}

	fmt.Println(burFile.Summary.TotalTests)

	return burFile
	
}
