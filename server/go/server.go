package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/protobuf/types/known/structpb"
)

type algoliaAnswersReq struct {
	Query                   string   `json:"query"`
	QueryLanguages          []string `json:"queryLanguages"`
	AttributesForPrediction []string `json:"attributesForPrediction"`
	NbHits                  int      `json:"nbHits"`
}

type algoliaAnswersHit struct {
	Question string `json:"q"`
	Answer   string `json:"a"`
}

type algoliaAnswersRes struct {
	Hits []algoliaAnswersHit `json:"hits"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(os.Getenv("STATIC_DIR"))))
	http.HandleFunc("/webhook", handleWebhook)

	addr := "0.0.0.0:4242"
	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var webhook_req dialogflow.WebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&webhook_req); err != nil {
		writeJSON(w, nil, err)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	// Query the Algolia Answers API.
	// https://www.algolia.com/doc/guides/algolia-ai/answers/#finding-answers-in-your-index
	url := fmt.Sprintf(
		"https://%s-dsn.algolia.net/1/answers/%s/prediction",
		os.Getenv("ALGOLIA_APP_ID"),
		os.Getenv("ALGOLIA_INDEX_NAME"),
	)
	data := &algoliaAnswersReq{
		Query:                   webhook_req.QueryResult.QueryText,
		QueryLanguages:          []string{"en"},
		AttributesForPrediction: []string{"q", "a"},
		NbHits:                  1,
	}
	body, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("ContentType", "application/json")
	req.Header.Set("X-Algolia-Api-Key", os.Getenv("ALGOLIA_API_KEY"))
	req.Header.Set("X-Algolia-Application-ID", os.Getenv("ALGOLIA_APP_ID"))
	if err != nil {
		writeJSON(w, nil, err)
		log.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		writeJSON(w, nil, err)
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var algoliaRes algoliaAnswersRes
	err = json.NewDecoder(r.Body).Decode(&algoliaRes)
	if err != nil {
		writeJSON(w, nil, err)
		log.Println(err)
		return
	}

	payload := dialogflow.Intent_Message_Payload{
		Payload: &structpb.Struct{
			richContent: "test",
		},
	}

	if len(algoliaRes.Hits) == 0 {
		webhookRes := dialogflow.WebhookResponse{
			FulfillmentMessages: []*dialogflow.Intent_Message{
				Message: &payload,
			},
		}
	} else {

	}
}

type errResp struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, v interface{}, err error) {
	var respVal interface{}
	if err != nil {
		msg := err.Error()
		w.WriteHeader(http.StatusBadRequest)
		var e errResp
		e.Error = msg
		respVal = e
	} else {
		respVal = v
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(respVal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
