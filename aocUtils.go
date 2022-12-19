package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/toast.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var config *Config
var template string
var startTime = time.Date(2022, 12, 1, 5, 0, 0, 0, time.UTC)

type Config struct {
	SessionId string `json:"sessionId"`
}

func LoadConfig() error {
	config = &Config{}
	configFile, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		return err
	}
	return nil
}

func getInput(year int, day int) (input string, err error) {
	path := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	header := http.Header{}
	header.Add("Cookie", fmt.Sprintf("session=%s", config.SessionId))
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return "", err
	}
	request.Header = header
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(string(body), "Please don't repeatedly request this endpoint before it unlocks!") {
		time.Sleep(1 * time.Second)
		return getInput(year, day)
	}
	return string(body), nil
}

func saveInput(year int, day int, input string) (err error) {
	dir := fmt.Sprintf("%d", year)
	err = os.Mkdir(dir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	path := fmt.Sprintf("%s/%d", dir, day)
	err = os.Mkdir(path, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	path = fmt.Sprintf("./%d/%d/input.txt", year, day)
	err = ioutil.WriteFile(path, []byte(input), 0644)
	if err != nil {
		return err
	}
	return nil
}

func saveTemplate(year int, day int) (err error) {
	if _, err := os.Stat(fmt.Sprintf("./%d/%d/main.go", year, day)); !os.IsNotExist(err) {
		return nil
	}
	dir := fmt.Sprintf("%d", year)
	err = os.Mkdir(dir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	path := fmt.Sprintf("%s/%d", dir, day)
	err = os.Mkdir(path, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	path = fmt.Sprintf("./%d/%d/main.go", year, day)
	err = ioutil.WriteFile(path, []byte(template), 0644)
	return err
}

func allInputFromYear(year int, maxDay int) {
	wg := sync.WaitGroup{}
	for i := 0; i < maxDay; i++ {
		wg.Add(1)
		i := i
		go func() {
			input, err := getInput(year, i+1)
			if err != nil {
				panic(err)
			}
			log.Println("Saving input for year", year, "day", i+1)
			err = saveInput(year, i+1, input)
			if err != nil {
				panic(err)
			}
			err = saveTemplate(year, i+1)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func timeTillNextDay() time.Duration {
	now := time.Now().UTC()
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 5, 0, 0, 0, time.UTC)
	diff := next.Sub(now)
	return diff
}

func nextDayCount() int {
	now := time.Now().UTC()
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 5, 0, 0, 0, time.UTC)
	diff := next.Sub(startTime)
	day := diff.Hours() / 24
	return roundUp(day) + 1
}

func roundUp(day float64) int {
	return int(day + 0.5)
}

func alert(year int, day int) {
	notification := toast.Notification{
		AppID:   "Advent of Code Utils",
		Title:   "Time to code!",
		Message: "The next day of Advent of Code has started!",
		Icon:    "D:\\GolandProjects\\adventofcode\\img.png",
		Actions: []toast.Action{
			{"protocol", "Open Url", fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func alertLoop(year int) {
	for {
		day := nextDayCount()
		duration := timeTillNextDay()
		log.Println("Sleeping for", duration, "until day", day)
		time.Sleep(duration + time.Second)
		alert(year, day)
		input, err := getInput(year, day)
		if err != nil {
			panic(err)
		}
		err = saveInput(year, day, input)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	YEAR := 2022
	DAY := 19

	err := LoadConfig()
	if err != nil {
		panic(err)
	}
	tmp, err := os.ReadFile("template.txt")
	if err != nil {
		panic(err)
	}
	template = string(tmp)
	template = fmt.Sprintf(template, YEAR, DAY)
	input, err := getInput(YEAR, DAY)
	if err != nil {
		panic(err)
	}
	err = saveInput(YEAR, DAY, input)
	if err != nil {
		panic(err)
	}
	err = saveTemplate(YEAR, DAY)
}
