# Notiflex 여정 기록

이 파일은 독자가 실제로 진행한 내용을 기록한다. AI가 각 챕터 완료 시 자동으로 업데이트한다.

## 진행 현황

| 챕터 | 서브챕터 | 상태 | 완료일 | 비고 |
|------|---------|------|--------|------|
| ch2 | 2.2 설치 확인 | ✅ | 2026-07-05 | |
| ch2 | 2.3 gcloud 설정 | ✅ | 2026-07-05 | 프로젝트 my-gitaiops, 리전 asia-northeast3 |
| ch2 | 2.4 GitHub 저장소 | ✅ | 2026-07-05 | Son-Hunseo/notiflex-platform |
| ch2 | 2.5 GKE 클러스터 | ✅ | 2026-07-05 | notiflex-cluster (Zonal, Spot) |
| ch2 | 2.6 빌드/배포 | ✅ | 2026-07-05 | api:v0.1.0, notiflex 네임스페이스 |
| ch2 | 2.7 첫 커밋 | ✅ | 2026-07-05 | |
| ch3 | 3.2 GitOps 도구 | ✅ | 2026-07-12 | ArgoCD v3.4.5, notiflex-smb Application Synced/Healthy |
| ch3 | 3.3 기능 추가 | ✅ | 2026-07-12 | `/version` 엔드포인트 추가, api:v0.1.1 롤링 업데이트 배포 |
| ch3 | 3.4 CI | ✅ | 2026-07-12 | `.github/workflows/ci.yaml` (방식 A: docker build+push), GCP_PROJECT_ID/GCP_SA_KEY Secret 등록, github-ci SA에 artifactregistry.writer 권한 확인 |
| ch3 | 3.5 CI-CD 연결 | ✅ | 2026-07-12 | CI가 매니페스트 이미지 태그 자동 업데이트 후 git push, ArgoCD 자동 Sync. api:v0.1.2(sha-aed9c36)로 e2e 검증 완료 |
| ch4 | 4.2 메트릭 모니터링 | ✅ | 2026-07-12 | kube-prometheus-stack(Prometheus+Grafana+Alertmanager) 설치, `helm-values/kube-prometheus.yaml`로 리소스 축소. port-forward로 Grafana 접속(admin/시크릿 PW) 확인, Prometheus 타겟 16/18 up 확인(coredns 2개 down은 GKE 관리형 환경의 정상 현상) |
| ch4 | 4.3 로그 수집 | ✅ | 2026-07-12 | Loki(SingleBinary) + Fluent Bit 설치, `helm-values/loki.yaml`·`helm-values/fluent-bit.yaml` 작성. `k8s/monitoring/loki-datasource.yaml`로 Grafana Loki 데이터소스 추가 배포. Grafana Explore에서 `{namespace="monitoring"}` 등으로 실제 로그 조회 확인 |
| ch4 | 4.4 알림 | ✅ | 2026-07-12 | `k8s/monitoring/pod-restart-alert.yaml`(PrometheusRule) 배포, Alertmanager는 kube-prometheus-stack 기본값 사용. Prometheus가 규칙 로드함(`notiflex-alerts`/`PodRestartTooMany`) 확인 |
| ch5 | 5.2 트래픽 관리 | ✅ | 2026-07-19 | Gateway API(`gke-l7-regional-external-managed`), Gateway+HTTPRoute+HealthCheckPolicy 배포. proxy-only 서브넷 없어 생성 필요했음. 외부 IP로 /health,/id,/version 확인 |
| ch5 | 5.3 무중단 배포 | ✅ | 2026-07-19 | Argo Rollouts v1.9.1 설치, Deployment→Rollout(Blue/Green) 전환, CI sed 대상을 rollout.yaml로 갱신, api:v0.2.0 배포로 auto-promote(30s) 동작 확인 |
| ch6 | 6.1 캐시 | ✅ | 2026-07-20 | Valkey(standalone) 설치, 인메모리 카운터→Valkey INCR로 교체. ch6 CPU 예산 확보를 위해 replicas 2→1 선제 축소(B/G) |
| ch6 | 6.2 시크릿 관리 | ✅ | 2026-07-20 | Workload Identity + GKE Secret Manager CSI로 Valkey 비밀번호를 파일 마운트 방식으로 전환. CPU 부족으로 Alertmanager/Grafana/operator/kube-state-metrics 임시 축소(replicas 0) — ch7.2 노드풀 추가 후 복원 필요 |
| ch6 | 6.3 Canary 전환 | ✅ | 2026-07-20 | `rollout.yaml` strategy `blueGreen`→`canary`(20→50→80→100%, 각 30초 pause) 전환, api:v0.4.0 배포로 e2e 검증 |
| ch6 | 6.4 아키텍처 스냅샷 | ✅ | 2026-07-20 | `claude-context/architecture.md` 신규 작성 — 3층 지식 구조(CLAUDE.md/claude-context/ADR), 클러스터 토폴로지, 컴포넌트 다이어그램, 배포 파이프라인, 관측 가능성, 주요 네임스페이스 표 |
| ch7 | 7.2 멀티 노드풀 | ⬜ | | |
| ch7 | 7.3 App of Apps | ⬜ | | |
| ch7 | 7.4 멀티테넌시 | ⬜ | | |
| ch8 | 8.1 메시징 | ⬜ | | |
| ch8 | 8.2 트레이싱 | ⬜ | | |
| ch8 | 8.3 CronJob | ⬜ | | |
| ch9 | 9.1 저장소 분석 | ⬜ | | |
| ch9 | 9.2 회고 | ⬜ | | |
| ch9 | 9.3 온보딩 문서 | ⬜ | | |
| ch9 | 9.4 GitAIOps 분석 | ⬜ | | |
| ch9 | 9.5 마무리 | ⬜ | | |

## 도구 선택 기록

독자가 3-프롬프트 패턴(탐색→비교→실행)에서 실제로 선택한 도구와 이유를 기록한다.

| 영역 | 선택 | 검토한 대안 | 선택 이유 |
|------|------|-----------|----------|
| GitOps 도구 (3.2) | ArgoCD | Flux, Jenkins X, Spinnaker | Web UI로 배포 상태를 시각적으로 확인 가능, e2-medium 노드에서 감당 가능한 리소스(~500m) |
| CI 도구 (3.4) | GitHub Actions | Cloud Build, GitLab CI, Jenkins | GitHub 네이티브라 별도 서버 설치/관리 불필요, `.github/workflows/ci.yaml` YAML 선언 하나로 파이프라인 정의, 프라이빗 저장소도 월 2,000분 무료, `google-github-actions/auth`로 GCP 서비스 계정 연동 간편 |
| 메트릭 모니터링 (4.2) | kube-prometheus-stack | Google Cloud Managed Service for Prometheus, Prometheus+Grafana 개별 차트 설치 | Prometheus Operator가 제공하는 ServiceMonitor/PrometheusRule CRD를 매니페스트로 관리해 GitOps(ArgoCD) 흐름에 그대로 편입 가능, Grafana·Alertmanager·kube-state-metrics·node-exporter가 한 번에 설치되어 4.3(로그)·4.4(알림)에서 재사용 가능, `helm-values/kube-prometheus.yaml`로 requests를 낮춰 e2-medium 2노드에서도 감당 가능 |
| 로그 수집 (4.3) | Loki + Fluent Bit | ELK Stack, CloudWatch Logs, Google Cloud Logging | 경량(Loki 128Mi, Fluent Bit 64Mi)이라 e2-medium에서 ELK(2Gi+)는 불가능, 4.2에서 설치한 Grafana에 데이터소스만 추가하면 메트릭과 같은 UI에서 로그 조회 가능, 라벨 기반 인덱싱이라 풀텍스트 인덱싱 대비 저장 비용 낮음 |
| 알림 (4.4) | PrometheusRule + Alertmanager(kube-prometheus-stack 기본값) | Grafana Alerting, Cloud Monitoring 알림 정책 | 4.2에서 Alertmanager가 이미 함께 설치되어 추가 리소스 없이 바로 사용 가능, PrometheusRule이 CRD 매니페스트라 `k8s/monitoring/`에 버전관리 및 ArgoCD 배포 가능, Grafana Alerting·Cloud Monitoring은 UI/콘솔 설정 위주라 GitOps 원칙에서 벗어남 |
| 외부 트래픽 관리 (5.2) | Gateway API | Ingress NGINX, Istio, Traefik | GKE Standard에서 별도 Controller 설치 없이 네이티브로 지원(`gke-l7-regional-external-managed`), Gateway(인프라)/HTTPRoute(앱) 역할 분리, Ingress를 대체하는 K8s 공식 표준, 5.3 Blue/Green의 HTTPRoute backendRefs 트래픽 분배와 연동 |
| 무중단 배포 (5.3) | Argo Rollouts — Blue/Green | Flagger, K8s native Rolling Update | 같은 Argo 생태계라 ArgoCD와 통합이 매끄럽고 Rollout 상태를 UI에서 확인 가능, CRD 기반 YAML 선언이라 GitOps 호환, 6장에서 Canary로 전략을 바꿀 때 Rollout CRD만 수정하면 되는 점진적 진화 경로, 2 replica 규모라 Blue/Green의 리소스 2배 부담이 크지 않음 |
| 배포 전략 고도화 (6.3) | Argo Rollouts — Canary | Blue/Green 유지 | 20%→50%→80%→100% 점진 전환으로 검증 실패 시 노출 범위를 최소화, Blue/Green(2x) 대비 리소스 효율(1.2x), 같은 Rollout CRD에서 strategy 필드만 변경해 새 도구 설치 불필요 |
| 캐시/상태 공유 (6.1) | Valkey | Redis, Memcached, DragonflyDB | INCR로 원자적 순차 ID 생성 필요 + Pod 재시작 후에도 카운터 유지(영속성) 필요 → Memcached는 INCR 미지원·영속성 없어 탈락. Redis는 2024년 SSPL 라이선스 전환으로 상용 제한 있음. Valkey는 Redis 포크로 100% 호환되면서 BSD 라이선스 유지, Bitnami 차트로 간편 설치, e2-medium 2노드 환경에서 50m 경량 운영 가능 |
| Secret 관리 (6.2) | Secrets Store CSI Driver + GCP Secret Manager | Sealed Secrets, External Secrets Operator, kubectl create secret | GKE 환경이라 Workload Identity로 SA 키 없이 네이티브 인증 가능. CSI 파일 마운트 방식은 K8s Secret 리소스를 만들지 않아(secretObjects 미사용) etcd/RBAC 노출 표면이 ESO(K8s Secret 자동 생성)보다 작음 — 단, 이 이점을 실제로 살리려면 앱이 env var가 아니라 마운트된 파일을 직접 읽어야 해서 `app/main.go`를 파일 읽기 방식으로 변경 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 초기 설정 (2.6) |
| Notiflex 이미지 | api:sha-b2940b7 (내부 버전 v0.4.0) | CI-CD e2e 테스트로 배포, 이후 태그는 git SHA 기반 (3.5). Blue/Green 전환 후 v0.2.0으로 배포 테스트, `/version` 응답 갱신 (5.3). Valkey INCR로 ID 생성 방식 전환, `/version` 응답 v0.3.0으로 갱신 (6.1). Valkey 비밀번호를 CSI 파일 마운트(`VALKEY_PASSWORD_FILE`)로 읽도록 변경 (6.2). Canary 전환 후 `/version` 응답 v0.4.0으로 갱신 (6.3) |
| Valkey | chart 6.2.0 (appVersion 9.1.0) | standalone 모드 최초 설치, `notiflex` 네임스페이스 (6.1) |
| ArgoCD | v3.4.5 | 초기 설치 (3.2) |
| kube-prometheus-stack | chart 87.15.1 | 초기 설치, requests 축소 (4.2). CPU 부족으로 Alertmanager/Grafana/operator/kube-state-metrics 임시 replicas=0 (6.2, ch7.2 복원 예정) |
| Loki | 3.6.7 (chart 7.0.0, SingleBinary) | 초기 설치 (4.3). CPU 부족으로 6.2에서 helm uninstall, ch7.2 복원 예정 |
| Fluent Bit | grafana/fluent-bit-plugin-loki 2.1.0-amd64 | 초기 설치 (4.3). CPU 부족으로 6.2에서 helm uninstall, ch7.2 복원 예정 |
| Argo Rollouts | v1.9.1 | 초기 설치, Deployment→Rollout Blue/Green 전환 (5.3). Blue/Green→Canary 전략 전환 (6.3) |
| Kafka | | |
| OTel SDK | | |

## 현재 리소스

⚠️ **2026-07-20 비용 절감을 위해 GCP 리소스 전체 정리됨** (ch6.4 완료 시점 → 재정리). 저장소 코드(`app/`, `k8s/`, `helm-values/`, `argocd/`, `claude-context/`)는 유지.
  - 삭제 순서: GKE 클러스터 `notiflex-cluster`(async 삭제, Compute 인스턴스 2개 + 노드 부트 디스크 2개는 클러스터 삭제 시 함께 정리됨) → IAM 바인딩 제거(`notiflex-secrets@`의 WI 바인딩, `github-ci@`의 프로젝트 레벨 `artifactregistry.writer`) → Secret Manager `valkey-password` 삭제(binding도 함께 제거됨) → Artifact Registry `notiflex` 삭제(이미지 전부 포함) → Service Account `notiflex-secrets@my-gitaiops.iam.gserviceaccount.com`, `github-ci@my-gitaiops.iam.gserviceaccount.com` 삭제
  - **Gateway API 고아 리소스 정리** (클러스터 삭제 후에도 자동 삭제되지 않음, ch5.2에서 생성): forwarding-rule, target-http-proxy, url-map, address, backend-service ×3(notiflex-api/gw-serve404/gw-serve500), health-check ×3을 의존성 역순으로 수동 삭제
  - **PVC 고아 Persistent Disk 정리** (ch6.1 Valkey standalone PVC, 클러스터 삭제로 자동 삭제되지 않음): `pvc-72288e44-3b66-4054-b552-8a5d6247d502` 삭제
  - 로컬 kubeconfig 컨텍스트/클러스터/유저 `gke-sysnet4admin_book_gitaiops` 삭제 (다른 로컬 컨텍스트는 유지)
  - **전체 스윕 검증 완료**: GKE 클러스터·Artifact Registry·Service Account(notiflex/github-ci)·Secret Manager·Compute 인스턴스·Persistent Disk·Forwarding Rule·Target Proxy·URL Map·Address·Backend Service·Health Check → 전부 없음 확인. (`Compute Engine default SA`는 시스템 계정으로 유지)
  - **재점검에서 발견 — GCS 버킷 고아 리소스** (가드레일/resource-budget.md에 명시 안 된 신규 발견 사항): 1차 정리 후 독자 요청으로 리전 제한 없이 프로젝트 전체를 재스윕한 결과 `gs://my-gitaiops_cloudbuild`(Cloud Build가 CI 빌드마다 자동 생성하는 소스 아카이브 버킷, `source/*.tgz` 3개, 4KB 미만) 발견 → 버킷째로 삭제. 이 버킷은 2026-07-12에도 한 번 정리된 이력이 있는데(당시엔 가드레일에 없던 항목으로 발견) 이후 CI 실행마다 자동 재생성되므로, **다음 정리 때도 `gcloud storage ls --project=<PROJECT>`로 별도 확인이 필요하다** — GKE/Artifact Registry 삭제 체크리스트에는 포함되지 않는 리소스.
  - ⚠️ **재개 시 주의** ([[project_gcp_resource_cleanup_cycle]]): GitHub Secret `GCP_SA_KEY`는 이번 정리로 삭제되지 않고 예전 값 그대로 남아 있다. `github-ci` SA를 재생성하기 전까지는 값이 유효하지 않으므로, 다음 재개 시 CI를 처음 돌리기 전 SA 재생성(`roles/artifactregistry.writer` 재부여) 및 `gh secret set GCP_SA_KEY` 갱신이 필요하다.
> 재개 시 다음 진행 지점: ch7.2(멀티 노드풀)부터 시작 가능. ch2.5~ch6.4 리소스를 처음부터 재생성해야 한다.

✅ **2026-07-20 ch6.3 Canary 전환**: `k8s/smb/rollout.yaml`의 `strategy.blueGreen`을 `strategy.canary`(canaryService `notiflex-api-preview`, stableService `notiflex-api`, steps 20→50→80→100%·각 30초 pause)로 변경 → git push 후 `kubectl delete rollout` → ArgoCD `refresh=hard`로 전략 전환 적용. `app/main.go` 버전을 v0.4.0으로 올려 CI가 `api:sha-b2940b7` 빌드·`rollout.yaml` 이미지 자동 갱신 → e2e 테스트 진행. CPU 한계(2 노드 모두 requests 근접)로 신규 Canary Pod이 `Insufficient cpu`로 Pending → `kubectl argo rollouts promote --full` + 구 ReplicaSet `replicas=0`으로 해결(가드레일 트러블슈팅 항목과 동일 패턴). 최종 Rollout Healthy, `/version` 응답 v0.4.0 확인. preview Service는 Canary의 canaryService로 계속 재사용 중이므로 삭제 금지.

⚠️ **2026-07-20 ch6.2 Secret 관리 도입 — CPU 한계 도달, 모니터링 일부 컴포넌트 임시 비활성화**: GKE Workload Identity(클러스터+노드풀) + `--enable-secret-manager` CSI addon 활성화. `gcloud secrets create valkey-password`로 GCP Secret Manager에 저장(`printf`로 개행 없이), `notiflex-secrets` GSA 생성 후 `roles/secretmanager.secretAccessor` 부여, `notiflex` 네임스페이스의 `notiflex-api` KSA를 Workload Identity로 GSA와 바인딩. `k8s/smb/secretproviderclass.yaml`(provider: gke) + `k8s/smb/serviceaccount.yaml` 추가, `rollout.yaml`에 CSI 볼륨(`secrets-store-gke.csi.k8s.io`) 마운트 및 `VALKEY_PASSWORD_FILE` 환경변수 추가. `app/main.go`가 마운트된 파일에서 비밀번호를 읽도록 변경(`secretObjects` 미사용 — K8s Secret 리소스를 만들지 않아 노출 표면 최소화).
  - **예상보다 CPU 여유가 부족했던 이유**: Workload Identity 활성화가 `gke-metadata-server` DaemonSet(노드당 100m, 합계 200m)을 추가로 배포한다는 점이 `shared/resource-budget.md`/가드레일에 명시되어 있지 않았음(신규 발견 사항). CSI DaemonSet(240m)과 합쳐 예상보다 더 많은 CPU가 소진됨.
  - ch6.1에서 이미 Loki+FluentBit를 제거하고 monitoring requests를 축소했음에도 CSI 설치 후 두 노드 모두 CPU 96~99%에 도달, Blue/Green이 신규+기존 Pod을 동시에 띄우지 못해 배포가 멈춤.
  - **임시 조치**: `kubectl scale`로 `kube-prometheus-kube-prome-operator`, `kube-prometheus-grafana`, `kube-prometheus-kube-state-metrics` Deployment와 `alertmanager` StatefulSet을 `replicas=0`으로 내려 Blue/Green 배포를 완료시킴. Prometheus 본체(5m)는 계속 Running 유지(메트릭 수집은 지속). 배포 완료 후 원복을 시도했으나 다시 Pending 발생 확인 → **ch7.2 노드풀 추가 전까지 replicas=0 유지**가 확정. Loki/FluentBit(ch6.1에서 제거)와 함께 ch7.2에서 반드시 복원해야 할 목록에 추가.
  - 검증: `/id` 호출로 CSI 파일 마운트 비밀번호로 Valkey 인증 성공 확인(Pod 로그에 재시도 없이 즉시 연결 성공), ID 카운터가 이전 값(34→35)에서 정상 이어짐.

✅ **2026-07-20 ch6.1 Valkey 캐시 도입**: `shared/resource-budget.md`가 경고한 ch6 CPU 위험 구간(노드 CPU requests 99%/79%로 이미 타이트) 확인 후, Valkey 설치 전 `k8s/smb/rollout.yaml`의 `replicas: 2 → 1`로 선제 축소(Git 경유, ArgoCD `argocd.argoproj.io/refresh=hard`로 강제 동기화). `helm install valkey bitnami/valkey -n notiflex`(standalone, resourcesPreset=none, requests cpu 50m/mem 64Mi) 설치 → `valkey-primary-0` Running, Secret 이름 `valkey`/key `valkey-password` 확인(가드레일 문서와 일치). `app/main.go`의 인메모리 `atomic.Uint64` 카운터를 `github.com/valkey-io/valkey-go` 클라이언트의 `INCR notiflex:id`로 교체, 10회/3초 간격 연결 재시도 로직 추가, `Dockerfile`을 `COPY go.mod go.sum ./`로 갱신. `git push` → CI가 `api:sha-ccd6ad1` 빌드/푸시 및 `rollout.yaml` 이미지 태그 자동 갱신 → ArgoCD Blue/Green 배포(Synced/Healthy). 검증: `/id` 호출로 `id: 1→2→3` 순차 증가 확인, Pod을 강제 삭제(`kubectl delete pod`) 후 재기동된 새 Pod에서 `/id` 호출 시 `id: 4`로 이어져 카운터가 Valkey에 영속되는 것(인메모리였다면 1로 리셋)을 확인. B/G replicas는 1로 유지 중이며, ch7에서 노드풀 추가 후 2로 복원 예정.

✅ **2026-07-19 재개 완료 — GCP 리소스 전체 재생성 (ch4까지)**: 2026-07-12 정리 이후 비어있던 `my-gitaiops` 프로젝트에 ch2.5~ch4.4 리소스를 처음부터 재생성.
  - GKE `notiflex-cluster`(asia-northeast3-a, e2-medium Spot ×2, disk 30GB, Gateway API standard) RUNNING, 컨텍스트 `gke-sysnet4admin_book_gitaiops`로 등록
  - Artifact Registry `notiflex` 재생성 후 기존 `app/` 소스(HEAD, ch3.5 시점과 코드 동일)를 그대로 빌드하여 매니페스트가 참조하던 태그 `api:sha-aed9c36` 그대로 재푸시 — `k8s/smb/deployment.yaml` 수정 없이 이미지 정합성 유지
  - ArgoCD stable manifest 설치 → 기본 포함된 NetworkPolicy가 repo-server의 GitHub egress를 막는 기존 트러블슈팅 이슈를 선제적으로 삭제 후 전체 rollout restart → `argocd/notiflex-smb.yaml` 적용 → `notiflex-smb` Application **Synced/Healthy**, `notiflex` 네임스페이스 Pod 2/2 Running, `/health`·`/id`·`/version`(v0.1.2) 확인
  - kube-prometheus-stack(`helm-values/kube-prometheus.yaml`) 설치 → Prometheus 타겟 16/18 up(코어DNS 2개 down은 기존과 동일한 GKE 관리형 환경의 정상 현상)
  - Loki(SingleBinary) + Fluent Bit(`helm-values/loki.yaml`·`helm-values/fluent-bit.yaml`) 설치, `k8s/monitoring/loki-datasource.yaml` 적용 → Loki API로 `{namespace="monitoring"}` 로그 실제 수신 확인
  - `k8s/monitoring/pod-restart-alert.yaml`(PrometheusRule) 재적용 → Prometheus 규칙 로드(`notiflex-alerts`/`PodRestartTooMany`) 확인. 규칙 반영까지 kubelet ConfigMap 동기화 주기(~1분) 만큼 지연되는 것을 관찰(신규 트러블슈팅 이력 아님, 정상 동작)
  - 모든 네임스페이스(`argocd`/`monitoring`/`notiflex`/`kube-system` 등) Pod 전수 Running 확인
> 재개 시 다음 진행 지점: ch5.2(트래픽 관리)부터 시작 가능. ch5 이후 리소스는 아직 미생성 상태.

✅ **2026-07-19 ch5.3 Argo Rollouts Blue/Green 전환**: `argo-rollouts` 네임스페이스에 컨트롤러(v1.9.1) 설치. 기존 `k8s/smb/deployment.yaml`을 삭제하고 `k8s/smb/rollout.yaml`(Rollout, `strategy.blueGreen`: activeService `notiflex-api` / previewService `notiflex-api-preview`, autoPromotionSeconds 30) + `k8s/smb/service-preview.yaml` 추가. `.github/workflows/ci.yaml`의 이미지 태그 sed 대상을 `deployment.yaml` → `rollout.yaml`로 갱신(두 곳: sed 라인, `git add` 라인). `api:v0.2.0`으로 배포해 `kubectl argo rollouts get rollout`으로 30초 후 auto-promote되어 active 전환되는 것을 확인, `/version` 응답이 v0.2.0으로 갱신됨을 검증.

✅ **2026-07-19 ch5.2 Gateway API 구성**: `k8s/smb/gateway.yaml`(Gateway+HTTPRoute), `k8s/smb/healthcheckpolicy.yaml` 추가 후 ArgoCD auto-sync로 배포. Gateway 최초 생성 시 `An active proxy-only subnetwork is required` 에러 발생 → `proxy-only-subnet`(172.16.0.0/23, REGIONAL_MANAGED_PROXY, asia-northeast3) 신규 생성으로 해결. Gateway 외부 IP `35.216.114.48` 할당(PROGRAMMED=True), `/health`·`/id`·`/version`(v0.1.2) 외부 접속 확인.

⚠️ **2026-07-12 비용 절감을 위해 GCP 리소스 전체 정리됨** (ch4 완료 시점 → 재정리). 저장소 코드(`app/`, `k8s/`, `helm-values/`, `argocd/`)는 유지.
  - 삭제: GKE 클러스터 `notiflex-cluster`(Compute 인스턴스 2개 + 30GB 디스크 2개 함께 정리) + Artifact Registry `notiflex`(이미지 전부 포함) + Service Account `github-ci@my-gitaiops.iam.gserviceaccount.com` + Cloud Build 소스/로그 버킷 `my-gitaiops_cloudbuild`
  - 로컬 kubeconfig 컨텍스트 `gke-sysnet4admin_book_gitaiops` 삭제 (`home-server` 컨텍스트는 유지)
✅ **2026-07-12 전체 스윕 검증 완료** — `my-gitaiops` 프로젝트에서 실습 생성 리소스가 하나도 남아있지 않음을 확인. 점검 항목: GKE 클러스터·Artifact Registry·Compute 인스턴스/디스크·Service Account(github-ci)·GCS 버킷·Forwarding Rule·Target Proxy·URL Map·Address·Backend Service·Health Check → 전부 없음. (`Compute Engine default SA`는 시스템 계정으로 유지). Secret Manager API는 비활성 상태로 별도 리소스 없음.

<details><summary>이전 이력 (2026-07-11 환경 이전 → 재생성, 2026-07-05 재개 → 재정리)</summary>

✅ **2026-07-11 다른 컴퓨터로 환경 이전 완료**: 기존 GCP 리소스(GKE `notiflex-cluster`, Artifact Registry `notiflex`)는 그대로 재사용. 새 컴퓨터에서 로컬 환경만 재구성 — gcloud 기본 zone/region 설정(asia-northeast3-a/asia-northeast3), Docker Artifact Registry 인증(`gcloud auth configure-docker`), `gke-gcloud-auth-plugin` 설치, kubectl 컨텍스트 `gke-sysnet4admin_book_gitaiops` 등록(기존 로컬 `home-server` 컨텍스트는 유지). notiflex-platform 저장소는 이미 클론되어 있었고 origin과 동기화 상태 확인. `notiflex` 네임스페이스 Pod 2/2 Running, `/health`·`/id` 재확인. 3장부터 이어서 진행 가능.

✅ **2026-07-11 재개 완료 — GCP 리소스 재생성**: 2.5(클러스터) → 2.6(Artifact Registry + 빌드/배포) 재실행. GKE `notiflex-cluster`(asia-northeast3-a, e2-medium Spot × 2, Gateway API standard) RUNNING, 컨텍스트 `gke-sysnet4admin_book_gitaiops`로 등록. Artifact Registry `notiflex` 재생성, 이미지 `api:v0.1.0` 빌드·푸시(Cloud Build). `notiflex` 네임스페이스 Pod 2/2 Running, `/health`·`/id` 정상 확인. 이제 3장부터 이어서 진행 가능.

<details><summary>이전 이력 (2026-07-05, 재개 → 재정리)</summary>

> ⚠️ **2026-07-05 비용 절감을 위해 GCP 리소스 전체 정리됨** (재개 → 재정리). 저장소 코드(`app/`, `k8s/`)는 유지.
>   - 삭제: GKE 클러스터 `notiflex-cluster`(노드 인스턴스·부트 디스크 함께 정리) + Artifact Registry `notiflex`
>   - 로컬 kubeconfig 컨텍스트 `gke-sysnet4admin_book_gitaiops` 삭제
> ✅ **2026-07-05 전체 스윕 검증 완료** — `my-gitaiops` 프로젝트에서 실습 생성 리소스가 하나도 남아있지 않음을 확인. 점검 항목: GKE 클러스터·Artifact Registry·Service Account(notiflex/github-ci)·Secret Manager·Compute 인스턴스/디스크·Forwarding Rule·Target Proxy·URL Map·Address·Backend Service·Health Check → 전부 없음. (`Compute Engine default SA`는 시스템 계정으로 유지). ch3+ 리소스는 애초에 미생성.
>
> ✅ **재개 완료 — GCP 리소스 재생성**: 2.5(클러스터) → 2.6(Artifact Registry + 빌드/배포) 재실행. GKE `notiflex-cluster`(e2-medium Spot × 2) RUNNING, 이미지 `api:v0.1.0` 빌드·푸시, `notiflex` 네임스페이스 Pod 2/2 Running, `/health`·`/id` 정상 확인.
> ⚠️ 이후 다시 비용 절감을 위해 위와 같이 전체 정리함.
> </details>

| 노드풀 | 머신 타입 | 노드 수 | 주요 워크로드 |
|--------|----------|---------|-------------|
| default-pool | e2-medium (Spot) | 2 | notiflex-api, ArgoCD, monitoring(Prometheus/Grafana/Alertmanager/Loki/Fluent Bit) |

✅ **2026-07-12 monitoring 네임스페이스 구성**: kube-prometheus-stack(Prometheus+Grafana+Alertmanager+operator+kube-state-metrics) + Loki(SingleBinary) + Fluent Bit(DaemonSet ×2) 설치 완료. `shared/resource-budget.md`의 ch4 완료 시점 예산(~320m)과 대체로 일치.

</details>

## 트러블슈팅 이력

독자가 겪은 문제와 해결 방법을 기록한다. 같은 문제를 다시 겪지 않도록 한다.

| 챕터 | 문제 | 해결 |
|------|------|------|
| 3.2 | notiflex-smb Application이 `Repository not found` 에러로 Sync Unknown 상태 지속 | 저장소가 private으로 되어 있어 ArgoCD가 익명 접근 실패. GitHub 저장소를 Public으로 전환 후 `kubectl annotate application notiflex-smb -n argocd argocd.argoproj.io/refresh=hard`로 강제 refresh하여 Synced/Healthy 전환 |
| 4.3 | Loki Pod이 `mkdir /var/loki: read-only file system`로 CrashLoopBackOff | `singleBinary.persistence.enabled: false`로 하면 `/var/loki`에 볼륨 자체가 마운트되지 않는다(가드레일에 없는 이슈). `singleBinary.extraVolumes`/`extraVolumeMounts`로 emptyDir을 직접 마운트하여 해결 (Spot VM 환경이라 PVC 대신 emptyDir로도 충분, 재시작 시 로그 유실은 허용) |
| 4.3 | Fluent Bit가 `dial tcp: lookup fluent-bit-loki ... no such host`로 전송 실패 반복 | `grafana/fluent-bit` 차트(`fluent-bit-plugin-loki` 이미지 기반)는 `config.outputs` raw 설정이 아니라 `loki.serviceName` 값으로 호스트를 결정하며, 비워두면 `${RELEASE}-loki`로 추정한다. Loki 릴리스명이 `loki`인데 Fluent Bit 릴리스명(`fluent-bit`) 기준으로 `fluent-bit-loki`를 찾아 실패. `helm-values/fluent-bit.yaml`에 `loki.serviceName: "loki"` 명시하여 해결 |
| 4.3 | 일부 Pod 로그 전송 시 `400 Bad Request ... has 16 label names; limit 15` | `config.autoKubernetesLabels: true`로 하면 Pod에 붙은 모든 k8s 라벨이 그대로 Loki 스트림 라벨이 되어, 라벨이 많은 Helm 배포 Pod(kube-prometheus-operator, loki 등)에서 Loki의 라벨 개수 제한(15개)을 초과. `autoKubernetesLabels: false`로 되돌리고 기본 `labelMap`(namespace/pod/container 등)만 사용하도록 해결 |
| 4.3 | 설치 후 Grafana에 Loki 데이터소스가 자동으로 나타나지 않음 | `helm-values/loki.yaml`에 넣은 `grafana.datasource.isDefault: false`는 실제로는 `grafana/loki` 차트(v7.0.0)에 존재하지 않는 키라 조용히 무시된다(`helm show values`로 확인). kube-prometheus-stack Grafana의 sidecar가 `grafana_datasource: "1"` 라벨 붙은 ConfigMap을 감시하는 구조임을 확인하고, `k8s/monitoring/loki-datasource.yaml`을 직접 작성해 Loki datasource ConfigMap을 추가 배포하여 해결. Grafana Explore에서 `{namespace="monitoring"}` 등으로 조회 가능 확인 |
| 4.4 | 가드레일의 `kubectl delete pod`로는 `PodRestartTooMany` 알림이 실제로 발동하지 않음 | `increase(kube_pod_container_status_restarts_total[5m])`는 같은 Pod 안에서 컨테이너가 크래시 후 제자리 재시작될 때만 증가한다. `kubectl delete pod`는 restartCount 0인 새 Pod을 만들 뿐이라 카운터가 오르지 않음(Prometheus `ALERTS{alertname="PodRestartTooMany"}` 조회로 미발동 확인). `notiflex-api`는 ArgoCD(`selfHeal: true`)가 관리 중이라 실제 크래시 재현(livenessProbe 실패 주입)은 auto-sync를 끄고 진행해야 하며 일시적 응답 불가 리스크가 있어 이번엔 규칙 설정·로드 확인까지만 진행하고 실전 크래시 테스트는 생략함 |
| 6.3 | ArgoCD `refresh=hard` 1회로는 최신 커밋(CI가 이미지 태그 갱신한 커밋)이 바로 반영되지 않음 | Rollout 삭제 직후 refresh는 그 시점의 git 커밋(전략 변경 커밋)까지만 반영. CI의 "ci: update image to..." 커밋(뒤이어 push됨)은 별도로 `refresh=hard`를 한 번 더 보내야 `status.sync.revision`이 최신 SHA로 갱신됨(가드레일에 명시되지 않은 이슈, `kubectl get application -o jsonpath='{.status.sync.revision}'`으로 확인) |
| 6.3 | Canary 신규 Pod이 `0/2 nodes are available: 2 Insufficient cpu`로 계속 Pending | 가드레일 트러블슈팅 항목과 동일한 CPU 한계 상황 재현(stable 1 + canary 1 = 2 Pod 필요하나 두 노드 모두 CPU 부족). `kubectl argo rollouts promote --full`로 즉시 승격 후 구 ReplicaSet(`notiflex-api-65df798c5`)을 `kubectl scale replicaset --replicas=0`으로 내려 CPU 확보 → 신규 Pod 정상 스케줄링·Running 확인 |
