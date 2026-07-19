package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/valkey-io/valkey-go"
)

// version은 배포 버전을 나타낸다. 빌드 시 이미지 태그와 함께 갱신한다.
const version = "v0.3.0"

// valkeyClient는 Pod 간 공유되는 ID 카운터를 보관하는 Valkey 클라이언트다.
var valkeyClient valkey.Client

// podName은 요청을 처리한 Pod의 이름이다. K8s가 HOSTNAME 환경 변수로 주입한다.
var podName = func() string {
	if h := os.Getenv("HOSTNAME"); h != "" {
		return h
	}
	return "unknown"
}()

func connectValkey() valkey.Client {
	addr := os.Getenv("VALKEY_ADDR")
	password := os.Getenv("VALKEY_PASSWORD")
	if pwFile := os.Getenv("VALKEY_PASSWORD_FILE"); pwFile != "" {
		data, err := os.ReadFile(pwFile)
		if err != nil {
			log.Fatalf("Valkey 비밀번호 파일 읽기 실패: %v", err)
		}
		password = string(data)
	}

	var client valkey.Client
	var err error
	for i := 0; i < 10; i++ {
		client, err = valkey.NewClient(valkey.ClientOption{
			InitAddress: []string{addr},
			Password:    password,
		})
		if err == nil {
			return client
		}
		log.Printf("Valkey 연결 재시도 %d/10: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}
	log.Fatalf("Valkey 연결 실패: %v", err)
	return nil
}

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
	resp := valkeyClient.Do(r.Context(), valkeyClient.B().Incr().Key("notiflex:id").Build())
	id, err := resp.ToInt64()
	if err != nil {
		http.Error(w, "id 생성 실패", http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{
		"id":           strconv.FormatInt(id, 10),
		"generated_by": podName,
	})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	valkeyClient = connectValkey()
	defer valkeyClient.Close()

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
