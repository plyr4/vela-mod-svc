package main

import (
	"encoding/json"
	"io"
	"net/http"

	compiler "github.com/go-vela/server/compiler/native"
	"github.com/sirupsen/logrus"
)

// server is a simple http router for running the application as a service
func server(cfg *Config) error {
	// http health endpoint to run this application as an http service
	http.HandleFunc("/health", health)

	// modification endpoint for analyzing/modifying pipelines
	http.HandleFunc("/modify", modify(cfg))

	// serve http
	err := http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

// health responds with 200 OK for indicating application health
func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

func modify(cfg *Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// read the incoming pipeline modification request
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Errorf("unable to read incoming modification request: %v", err)
			return
		}

		modReq := &compiler.ModifyRequest{}
		err = json.Unmarshal(b, modReq)
		if err != nil {
			logrus.Errorf("unable to unmarshal incoming modification request: %v", err)
			return
		}

		// prepare a modification response
		//  store the original pipeline so in scenarios where modification returns
		//  an error, the service will respond with the original pipeline
		modResp := &compiler.ModifyResponse{
			Pipeline: modReq.Pipeline,
		}

		defer func() {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(modResp)
		}()

		modifiedBuild, err := modifyPipeline(cfg, modReq)
		if err != nil {
			logrus.Errorf("unable to modify pipeline: %v, returning original pipeline", err)

			return
		}

		// marshal the modified pipeline
		modifiedBuild_, err := json.Marshal(modifiedBuild)
		if err != nil {
			logrus.Errorf("unable to marshal modified pipeline: %v, returning original pipeline", err)

			return
		}

		// return the modified pipeline as a modification response
		modResp.Pipeline = string(modifiedBuild_)
	}
}
