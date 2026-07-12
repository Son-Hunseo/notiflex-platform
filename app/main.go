package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

// version은 배포 버전을 나타낸다. 빌드 시 이미지 태그와 함께 갱신한다.
const version = "v0.1.1"

// idCounter는 /id 요청마다 순차 증가하는 인메모리 카운터다.
var idCounter atomic.Uint64

// podName은 요청을 처리한 Pod의 이름이다. K8s가 HOSTNAME 환경 변수로 주입한다.
var podName = func() string {
	if h := os.Getenv("HOSTNAME"); h != "" {
		return h
	}
	return "unknown"
}()

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{
		"version":      version,
		"generated_by": podName,
	})
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	id := idCounter.Add(1)
	writeJSON(w, map[string]string{
		"id":           strconv.FormatUint(id, 10),
		"generated_by": podName,
	})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/id", idHandler)
	mux.HandleFunc("/version", versionHandler)

	addr := ":8080"
	log.Printf("notiflex-api listening on %s (pod=%s)", addr, podName)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
