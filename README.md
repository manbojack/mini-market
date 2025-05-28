# 🛒 mini-market

**`mini-market`** — учебный проект интернет-магазина с фокусом на практику DevOps. Включает микросервисы на Python и Go, PostgreSQL, Redis, Kafka, CI/CD через GitHub Actions, контейнеризацию с Docker, оркестрацию через Minikube и автоматизацию с Ansible.

---

## 📦 Стек технологий

- **Backend:** Python (FastAPI), Go
- **БД:** PostgreSQL
- **Кэш:** Redis
- **Брокер сообщений:** Apache Kafka
- **CI/CD:** GitHub Actions
- **Инфраструктура:** Docker, Minikube, Ansible

---

## 🗂 Структура проекта

```plaintext
mini-market/
│
├── ansible/                  # Playbook для развертывания
│   └── site.yml
│
├── ci-cd/                    # GitHub Actions workflows
│   └── main.yml
│
├── deploy/                   # YAML-файлы для Minikube/Kubernetes
│   ├── postgres.yaml
│   ├── redis.yaml
│   ├── kafka.yaml
│   ├── python-app.yaml
│   ├── go-app.yaml
│   └── ingress.yaml
│
├── go-app/                   # Сервис на Go
│   ├── main.go
│   └── Dockerfile
│
├── python-app/               # Сервис на Python (FastAPI)
│   ├── main.py
│   ├── requirements.txt
│   └── Dockerfile
│
├── docker-compose.yml        # Локальный запуск
├── Makefile                  # Утилиты
├── README.md                 # Документация
└── .github/
    └── workflows/
        └── ci.yml
```
