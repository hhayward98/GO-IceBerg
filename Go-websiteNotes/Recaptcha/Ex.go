package main

import (

	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

)


const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`

}

type SiteVerifyResponse struct {
	Success bool `json:"success"`
	Score float64 `json:"score"`
	Action string `json:"action"`
	ChallengeTS time.Time `json:"challange_ts"`
	Hostname string `json:"hostname"`
	ErrorCodes []string `json:"error-codes"`

}

type SiteVerifyRequest struct {
	RecaptchaResponse string `json:"g-recaptch-response"`
}


func Login(secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		//check credentials 

		// if authorized

		w.WriteHeader(http.StatusOk)
	}
}


func CheckRecaptcha(secret, reponse string) error {
	req, err : = http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	Qu := req.URL.Query()
	Qu.Add("secret", secret)
	Qu.Add("response", reponse)
	req.URL.RawQuery = Qu.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err !+ nil {
		return err
	}

	defer resp.Body.Close()

	var body SiteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil
	}

	if !body.Success {
		return errors.New("unsuccessful recaptcha verify request")
	}

	if body.Score < 0.5 {
		return errors.New("score was lower than expected")
	}

	if body.Action != "login" {
		return errors.New("Error, smismatched recaptcha action")
	}

	return nil

}

func RecaptchaMiddleware(secret string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			var body SiteVerifyRequest
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				http.Error(w,"Unauthorized", http.StatusUnauthorized)
				return
			}


			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			
			if err := CheckRecaptcha(secret, body.RecaptchaResponse); err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}


func main() {
	secret := os.Getenv("RECAPTCHA_SECRET")

	http.Handle("/login", RecaptchaMiddleware(secret)(Login()))

	log.Println("Listening.....")

	log.Fatal(http.ListenAndServe(":8080", nil))
}