package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/linkepay/linkepay-sdk-go"
	"github.com/linkepay/linkepay-sdk-go/types"
)

var client *linkepay.Client

func init() {
	client = linkepay.NewClient(&types.Config{
		BaseURL:              "https://api.linkepay.com",
		ProjectID:            "your-project-id",
		PublicKey:            "your-public-key",
		PrivateKey:           "your-private-key",
		PayPlatformPublicKey: "platform-public-key",
	})
}

type CallbackRequest struct {
	VerifyData map[string]interface{} `json:"verify_data"`
	Sig        string                 `json:"sig"`
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var callbackReq CallbackRequest
	if err := json.Unmarshal(body, &callbackReq); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ok, err := client.VerifyPlatformSignature(
		client.GetPlatformPublicKey(),
		callbackReq.VerifyData,
		callbackReq.Sig,
	)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		http.Error(w, "Verification error", http.StatusInternalServerError)
		return
	}

	if !ok {
		log.Printf("Invalid signature")
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	log.Printf("Callback verified successfully: %+v", callbackReq.VerifyData)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/callback", callbackHandler)

	addr := ":8080"
	fmt.Printf("Starting callback server on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
