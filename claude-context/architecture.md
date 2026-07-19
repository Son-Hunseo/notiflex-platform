# Notiflex 아키텍처 스냅샷

> 최종 갱신: 2026-07-20 (ch6.4)

## 3층 지식 구조

이 저장소는 AI가 참조하는 정보를 세 층으로 분리한다. 섞이면 AI가 "현재 상태"와 "과거 결정 이유"를 혼동하게 되므로 역할을 엄격히 구분한다.

- **`CLAUDE.md`** — AI에게 프로젝트 메타데이터를 제공한다. 매 대화 시작 시 자동 로드되며, 가드레일 참조 규칙·프로젝트 컨텍스트 등 "AI가 어떻게 행동해야 하는가"를 담는다.
- **`claude-context/`(이 문서)** — AI가 참조할 **현재 시점의 아키텍처 스냅샷**이다. "지금 무엇이 어떻게 동작하는가"만 담고, 항상 최신 상태로 덮어쓴다. 결정 이유나 과거 이력은 담지 않는다.
- **`docs/architecture-decisions.md`(ADR)** — "왜 이 결정을 내렸는가"의 누적 기록이다. 시간순으로 쌓이며 삭제하지 않는다. 사람과 AI가 함께 검토한다.

세부 진행 이력·트러블슈팅은 `JOURNEY.md`에 남는다. 이 문서(`claude-context/architecture.md`)는 항상 **한 페이지 현재 상태 요약**으로만 유지한다.

## 클러스터 토폴로지

| 항목 | 값 |
|------|-----|
| 클러스터 이름 | `notiflex-cluster` |
| 리전/존 | `asia-northeast3-a` (Zonal) |
| K8s 버전 | v1.35.5-gke.1241004 |
| 노드풀 | `default-pool` — `e2-medium` × 2, Spot VM |
| Workload Identity | 활성화 (`my-gitaiops.svc.id.goog`) |
| Gateway API | 활성화 (Standard, GatewayClass `gke-l7-regional-external-managed`) |
| Secret Manager CSI addon | 활성화 (`secrets-store-gke.csi.k8s.io`, provider: gke) |
| kubectl context | `gke-sysnet4admin_book_gitaiops` |

## 컴포넌트 다이어그램

```
[외부 클라이언트]
       │  HTTP (35.216.114.48)
       ▼
[Gateway: notiflex-gateway]  (gke-l7-regional-external-managed, HealthCheckPolicy)
       │
       ▼
[HTTPRoute: notiflex-route]  PathPrefix "/" → backendRef notiflex-api:80
       │
       ▼
[Service: notiflex-api] (stable) ──┬── [Service: notiflex-api-preview] (canary)
       │                           │
       ▼                           ▼
[Rollout: notiflex-api]  (Argo Rollouts, strategy.canary)
  steps: 20% → 50% → 80% → 100% (각 단계 30s pause)
       │
       ▼
[Pod: notiflex-api]  (replicas=1, scratch 이미지)
   ├── env VALKEY_ADDR → [Service: valkey-primary] → [StatefulSet: valkey-primary] (Bitnami Valkey, standalone)
   └── volumeMount(CSI) ← [SecretProviderClass: notiflex-secrets] ← GCP Secret Manager (valkey-password)
         (ServiceAccount notiflex-api ↔ Workload Identity ↔ GSA notiflex-secrets)
```

## 배포 파이프라인

```
로컬/PR: app/** 변경 → git push origin main
       │
       ▼
GitHub Actions (.github/workflows/ci.yaml)
  1. docker build (app/) → asia-northeast3-docker.pkg.dev/my-gitaiops/notiflex/api:sha-<7자리>
  2. docker push → Artifact Registry
  3. sed로 k8s/smb/rollout.yaml 이미지 태그 갱신 → git commit "ci: update image to ..." → push
       │
       ▼
ArgoCD (Application: notiflex-smb, syncPolicy: automated{prune,selfHeal})
  git 변경 감지 → Sync → k8s/smb/*.yaml 적용
       │
       ▼
Argo Rollouts 컨트롤러가 Rollout 스펙 변경 감지
  → 신규 ReplicaSet(canary) 생성 → steps대로 setWeight 점진 증가 → stable 승격
```

> ⚠️ 전략(`blueGreen`↔`canary`) 자체를 바꿀 때는 `kubectl delete rollout`으로 컨트롤러 상태를 초기화해야 하며, ArgoCD auto-sync와 충돌하지 않도록 **git push를 먼저** 한 뒤 삭제한다.

## 관측 가능성

| 도구 | 역할 | 현재 상태 |
|------|------|----------|
| Prometheus | 메트릭 수집, PrometheusRule(`pod-restart-alert`) 평가 | Running (monitoring 네임스페이스) |
| Alertmanager | 알림 라우팅 | ⚠️ replicas=0 (ch6.2 CPU 확보를 위해 임시 비활성화, ch7.2 노드풀 추가 후 복원 예정) |
| Grafana | 메트릭/로그 대시보드 | ⚠️ replicas=0 (동일 사유) |
| kube-state-metrics / operator | Prometheus Operator 지원 | ⚠️ replicas=0 (동일 사유) |
| Loki + Fluent Bit | 로그 수집/조회 | ⚠️ ch6.1에서 helm uninstall (CPU 확보), ch7.2 복원 예정 |
| Tempo | 분산 트레이싱 | 미도입 (ch8.2 예정) |

> 현재 CPU가 타이트해 모니터링 스택 일부가 축소된 상태다. 새 리소스를 배포하기 전에 `kubectl top nodes`로 여유를 확인할 것.

## 주요 네임스페이스

| 네임스페이스 | 주요 워크로드 |
|-------------|--------------|
| `notiflex` | Rollout `notiflex-api`(Canary), Service `notiflex-api`/`notiflex-api-preview`, StatefulSet `valkey-primary`, Gateway/HTTPRoute, SecretProviderClass `notiflex-secrets` |
| `argocd` | `argocd-server`, `argocd-repo-server`, `argocd-application-controller`(StatefulSet), `argocd-redis`, `argocd-dex-server` — Application `notiflex-smb` 관리 |
| `argo-rollouts` | Argo Rollouts 컨트롤러(v1.9.1) |
| `monitoring` | Prometheus(Running), Grafana/Alertmanager/operator/kube-state-metrics(replicas=0, 임시 비활성화) |
| `kube-system` | GKE 관리형 컴포넌트(kube-dns, metrics-server, konnectivity-agent 등) |
