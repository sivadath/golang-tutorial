package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"dath"
)

var scoreBoard ScoreCard


func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,`<a href="http://localhost:8071/live-score/">click here for latest scores</a>`)
}

func liveScore (w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, `<h1><meta http-equiv="refresh" content="1" />Current Score:%d</h1><h1>Overs count:%f</h1>`, scoreBoard.Score, scoreBoard.Overs)
}

func EventListener() {
	reader := bufio.NewReader(os.Stdin)
	for {
		event, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		scoreBoard.Notify(event[0:len(event)-1])
	}
}

func main() {
	fmt.Println(dath.D)
	http.HandleFunc("/", home)
	http.HandleFunc("/live-score/", liveScore)
	scoreBoard.AddListener(Balls{})
	scoreBoard.AddListener(Score{})
	go EventListener()
	log.Fatal(http.ListenAndServe(":8071",nil))
}



















type ScoreCard struct {
	Balls int
	Overs float64
	Score int
	listeners map[scoreUpdater]struct{}
}

func (sc *ScoreCard)AddListener(updater scoreUpdater) {
	if sc.listeners == nil {
		sc.listeners = map[scoreUpdater]struct{}{}
	}
	sc.listeners[updater] = struct{}{}
}

func (sc *ScoreCard)RemoveListener(updater scoreUpdater) {
	delete(sc.listeners,updater)
}

func (sc *ScoreCard)Notify(event string) {
	for eventHandler, _ := range sc.listeners {
		eventHandler.Update(event)
	}
}

type scoreUpdater interface {
	Update(event string)
}

type Balls struct {}

func (ball Balls)ParseEvent(event string) bool{
	if strings.HasPrefix(strings.ToUpper(event),"B") {
		event = event[1:]
		if event[len(event)-1] == '\r' {
			event = event[0:len(event)-1]
		}
		if ball, err := strconv.Atoi(event);err == nil {
			scoreBoard.Balls = ball
			return true
		}
	}
	return false
}

func (ball Balls)Update(event string){
	if ball.ParseEvent(event) {
		scoreBoard.Overs = float64(scoreBoard.Balls)/6
	}
}

type Score struct {}

func (score Score)Update(event string) {
	if strings.HasPrefix(strings.ToUpper(event),"S"){
		event = event[1:]
		if event[len(event)-1] == '\r' {
			event = event[0:len(event)-1]
		}
		if score, err := strconv.Atoi(event);err == nil {
			scoreBoard.Score = score
		}
	}
}