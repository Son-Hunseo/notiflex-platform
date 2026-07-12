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
| ch3 | 3.2 GitOps 도구 | ⬜ | | |
| ch3 | 3.3 기능 추가 | ⬜ | | |
| ch3 | 3.4 CI | ⬜ | | |
| ch3 | 3.5 CI-CD 연결 | ⬜ | | |
| ch4 | 4.2 메트릭 모니터링 | ⬜ | | |
| ch4 | 4.3 로그 수집 | ⬜ | | |
| ch4 | 4.4 알림 | ⬜ | | |
| ch5 | 5.2 트래픽 관리 | ⬜ | | |
| ch5 | 5.3 무중단 배포 | ⬜ | | |
| ch6 | 6.1 캐시 | ⬜ | | |
| ch6 | 6.2 시크릿 관리 | ⬜ | | |
| ch6 | 6.3 Canary 전환 | ⬜ | | |
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
| | | | |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 초기 설정 (2.6) |
| Notiflex 이미지 | api:v0.1.0 | 초기 배포 (2.6) |
| ArgoCD | | |
| Kafka | | |
| OTel SDK | | |

## 현재 리소스

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
| default-pool | e2-medium (Spot) | 2 | notiflex-api |

## 트러블슈팅 이력

독자가 겪은 문제와 해결 방법을 기록한다. 같은 문제를 다시 겪지 않도록 한다.

| 챕터 | 문제 | 해결 |
|------|------|------|
| | | |
