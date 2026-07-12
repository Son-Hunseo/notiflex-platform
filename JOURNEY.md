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
| GitOps 도구 (3.2) | ArgoCD | Flux, Jenkins X, Spinnaker | Web UI로 배포 상태를 시각적으로 확인 가능, e2-medium 노드에서 감당 가능한 리소스(~500m) |
| 메트릭 모니터링 (4.2) | kube-prometheus-stack | Google Cloud Managed Service for Prometheus, Prometheus+Grafana 개별 차트 설치 | Prometheus Operator가 제공하는 ServiceMonitor/PrometheusRule CRD를 매니페스트로 관리해 GitOps(ArgoCD) 흐름에 그대로 편입 가능, Grafana·Alertmanager·kube-state-metrics·node-exporter가 한 번에 설치되어 4.3(로그)·4.4(알림)에서 재사용 가능, `helm-values/kube-prometheus.yaml`로 requests를 낮춰 e2-medium 2노드에서도 감당 가능 |
| 알림 (4.4) | PrometheusRule + Alertmanager(kube-prometheus-stack 기본값) | Grafana Alerting, Cloud Monitoring 알림 정책 | 4.2에서 Alertmanager가 이미 함께 설치되어 추가 리소스 없이 바로 사용 가능, PrometheusRule이 CRD 매니페스트라 `k8s/monitoring/`에 버전관리 및 ArgoCD 배포 가능, Grafana Alerting·Cloud Monitoring은 UI/콘솔 설정 위주라 GitOps 원칙에서 벗어남 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 초기 설정 (2.6) |
| Notiflex 이미지 | api:sha-aed9c36 (v0.1.2) | CI-CD e2e 테스트로 배포, 이후 태그는 git SHA 기반 (3.5) |
| ArgoCD | v3.4.5 | 초기 설치 (3.2) |
| kube-prometheus-stack | chart 87.15.1 | 초기 설치, requests 축소 (4.2) |
| Loki | 3.6.7 (chart 7.0.0, SingleBinary) | 초기 설치 (4.3) |
| Fluent Bit | grafana/fluent-bit-plugin-loki 2.1.0-amd64 | 초기 설치 (4.3) |
| Kafka | | |
| OTel SDK | | |

## 현재 리소스

⚠️ **2026-07-12 비용 절감을 위해 GCP 리소스 전체 정리됨** (ch4 완료 시점 → 재정리). 저장소 코드(`app/`, `k8s/`, `helm-values/`, `argocd/`)는 유지.
  - 삭제: GKE 클러스터 `notiflex-cluster`(Compute 인스턴스 2개 + 30GB 디스크 2개 함께 정리) + Artifact Registry `notiflex`(이미지 전부 포함) + Service Account `github-ci@my-gitaiops.iam.gserviceaccount.com` + Cloud Build 소스/로그 버킷 `my-gitaiops_cloudbuild`
  - 로컬 kubeconfig 컨텍스트 `gke-sysnet4admin_book_gitaiops` 삭제 (`home-server` 컨텍스트는 유지)
✅ **2026-07-12 전체 스윕 검증 완료** — `my-gitaiops` 프로젝트에서 실습 생성 리소스가 하나도 남아있지 않음을 확인. 점검 항목: GKE 클러스터·Artifact Registry·Compute 인스턴스/디스크·Service Account(github-ci)·GCS 버킷·Forwarding Rule·Target Proxy·URL Map·Address·Backend Service·Health Check → 전부 없음. (`Compute Engine default SA`는 시스템 계정으로 유지). Secret Manager API는 비활성 상태로 별도 리소스 없음.
> 재개 시 2.5(클러스터) → 2.6(Artifact Registry + 빌드/배포)부터 재실행 필요. ch4까지의 매니페스트·Helm values는 저장소에 남아있으므로 클러스터 재생성 후 ArgoCD Sync + Helm 재설치로 복원 가능.

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
