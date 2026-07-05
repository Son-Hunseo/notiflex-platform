# Notiflex Platform

## 프로젝트 개요

**Notiflex** — B2B 알림 SaaS 플랫폼. 기업 고객에게 이메일·SMS·푸시 알림을 안정적으로 전송하는 서비스다.

## 기술 스택

- **언어**: Go 표준 라이브러리 (외부 프레임워크 없음)
- **컨테이너**: scratch 베이스 이미지 (최소 크기)
- **인프라**: GKE Standard (Zonal), Spot VM
- **GitOps**: ArgoCD

## GCP 설정

- **프로젝트 ID**: my-gitaiops
- **리전**: asia-northeast3 (서울)
- **존**: asia-northeast3-a
- **Artifact Registry**: `asia-northeast3-docker.pkg.dev/my-gitaiops/notiflex`

## 저장소 구조

```
notiflex-platform/
├── CLAUDE.md
├── app/           # Go 애플리케이션
├── k8s/
│   └── smb/       # K8s 매니페스트
└── .github/
    └── workflows/ # CI 파이프라인
```

## 행동 규칙

1. 변경 전 항상 현재 상태를 확인한다.
2. kubectl 명령에는 반드시 `--context gke-sysnet4admin_book_gitaiops`를 지정한다.
3. 리소스 삭제 전 반드시 확인한다.
