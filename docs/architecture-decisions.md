# Architecture Decision Records

## ADR-001: 배포 자동화는 ArgoCD (3장)
**시점**: 2026-07 / **결정**: GitOps 도구로 ArgoCD를 채택한다. Flux, Jenkins X, Spinnaker는 쓰지 않는다.
**이유**:
- Web UI로 배포 상태를 실시간·시각적으로 확인 가능 (학습 과정에서 "지금 무슨 일이 일어나는지" 눈으로 확인)
- Application CRD로 "어떤 Git 경로 → 어떤 네임스페이스" 선언적 관리
- Self-Heal: 누군가 `kubectl edit`으로 직접 수정해도 Git 상태로 자동 복구
- e2-medium 노드에서 감당 가능한 리소스(~500MB), GKE Standard와 네이티브 호환

## ADR-002: CI 도구는 GitHub Actions (3장)
**시점**: 2026-07 / **결정**: CI 파이프라인은 GitHub Actions로 구성한다. Cloud Build, GitLab CI, Jenkins는 쓰지 않는다.
**이유**:
- GitHub 네이티브 — 코드 저장소와 CI가 같은 플랫폼이라 별도 서버 설치/관리 불필요
- `.github/workflows/ci.yaml` YAML 선언 하나로 파이프라인 정의
- 프라이빗 저장소도 월 2,000분 무료 크레딧
- `google-github-actions/auth` 액션으로 GCP 서비스 계정 연동이 간편

## ADR-003: 메트릭은 kube-prometheus-stack (Prometheus + Grafana) (4장)
**시점**: 2026-07 / **결정**: 메트릭 모니터링은 kube-prometheus-stack으로 구성한다. Google Cloud Managed Service for Prometheus, 개별 차트 설치는 쓰지 않는다.
**이유**:
- Prometheus Operator의 ServiceMonitor/PrometheusRule CRD를 매니페스트로 관리해 ArgoCD GitOps 흐름에 그대로 편입
- Grafana·Alertmanager·kube-state-metrics·node-exporter가 한 번에 설치되어 이후 로그(4.3)·알림(4.4)에서 재사용
- `helm-values/kube-prometheus.yaml`로 requests를 낮춰 e2-medium 2노드에서도 감당 가능

## ADR-004: 로그는 Loki와 Fluent Bit (4장)
**시점**: 2026-07 / **결정**: 로그 수집은 Loki + Fluent Bit로 구성한다. ELK Stack, CloudWatch Logs, Google Cloud Logging은 쓰지 않는다.
**이유**:
- 경량(Loki 128Mi, Fluent Bit 64Mi) — e2-medium에서 ELK(2Gi+)는 불가능
- 4.2에서 설치한 Grafana에 데이터소스만 추가하면 메트릭과 같은 UI에서 로그 조회 가능
- 라벨 기반 인덱싱이라 풀텍스트 인덱싱 대비 저장 비용이 낮음

## ADR-005: 알림은 PrometheusRule과 Alertmanager (4장)
**시점**: 2026-07 / **결정**: 알림은 PrometheusRule + Alertmanager(kube-prometheus-stack 기본값)로 구성한다. Grafana Alerting, Cloud Monitoring 알림 정책은 쓰지 않는다.
**이유**:
- 4.2에서 Alertmanager가 이미 함께 설치되어 추가 리소스 없이 바로 사용 가능
- PrometheusRule이 CRD 매니페스트라 `k8s/monitoring/`에 버전관리 및 ArgoCD 배포 가능
- Grafana Alerting·Cloud Monitoring은 UI/콘솔 설정 위주라 GitOps 원칙에서 벗어남
- 규칙이 코드로 남아 `git blame`으로 "언제·왜 이 임계값을 정했는지" 추적 가능

## ADR-006: 외부 진입점은 Gateway API (5장)
**시점**: 2026-07 / **결정**: GKE Gateway API로 외부 진입점을 만든다. Ingress는 사용하지 않는다.
**이유**:
- GKE 네이티브 지원: 별도 Controller 설치 불필요 (`gke-l7-regional-external-managed`)
- 역할 분리: Gateway(인프라팀) / HTTPRoute(앱팀)로 관심사 분리
- K8s 공식 표준: Ingress의 후속 세대 (GA since K8s 1.27)
- Blue/Green 연동: HTTPRoute의 backendRefs로 트래픽 분배 가능

## ADR-007: 무중단 배포는 Blue/Green (5장)
**시점**: 2026-07 / **결정**: Argo Rollouts의 Blue/Green 전략을 사용한다. Canary는 아직 도입하지 않는다.
**이유**:
- Rolling Update 한계: 배포와 검증이 분리되지 않아 문제 발견이 늦음
- 같은 Argo 생태계라 ArgoCD와 통합이 매끄럽고, Rollout 상태를 UI·`kubectl argo rollouts`로 확인 가능
- 2 replica 규모에서 Blue/Green의 리소스 2배 부담이 크지 않음
- Canary는 메트릭 기반 자동 판정이 선행되어야 안전한데 현재는 미확보 — Rollout CRD 구조상 6장에서 Canary로의 전환 경로는 열려 있음
