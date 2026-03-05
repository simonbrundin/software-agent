# agents

Ett experimentellt repository för agent‑baserade arbetsflöden och verktyg. Denna README innehåller en kort introduktion, snabbstart, samt en integrerad sammanfattning av Agent OS‑produktinformationen (mission, tech stack och roadmap).

Innehåll

- Vad är detta
- Agent OS (mission, tech stack, roadmap)
- Snabbstart
- Installation
- Användning
- Utveckling
- Bidra
- Licens

Vad är detta

`agents` är ett utvecklingsrepo för att bygga och utvärdera agent‑baserade system som automatiserar kodunderhåll, testning och pull request‑arbete. Projektet kombinerar orkestrering, ephemeral deploys och LLM‑providers för att skala rutinuppgifter i mjukvaruutveckling.

**Agent OS — Produktöversikt**

Mission

Problem

Automatisera kontinuerlig förbättring och underhåll av mjukvaruprojekt genom agenter som arbetar dygnet runt och skapar pull requests för att implementera förbättringar, fixa buggar och skriva tester.

Målgrupp

Projektägare och utvecklingsteam som vill skala kodunderhåll, snabba upp feedback‑loopar och avlasta rutinuppgifter.

Lösning (högnivå)

- Starta en agent per GitHub‑issue och orkestrera flera agenter parallellt.
- Automatisera worktree/branch‑creation och PR‑skapande.
- Implementera test‑ och refactor‑loopar (code → tests → CI → refactor).
- Hantera modell‑provider‑failover och bevara kontext mellan providers.
- Skapa ephemeral miljöer för manuell testning före merge.
- Säkerhetsmodell: agenter körs isolerat med separata konton och arbetsmiljöer, utan åtkomst till användarens personliga maskin.

Tech Stack

- Frontend: Nuxt (Vue)
- Backend: Go
- Databas: PostgreSQL
- ORM: Drizzle
- Deployment / Orkestrering: Kubernetes
- Övrigt: CI (GitHub Actions), provider SDKs (OpenAI eller andra LLM), observability (Prometheus/Grafana föreslaget), messaging/queueing vid behov (RabbitMQ/NATS)

Dev & Ephemeral detaljer

- Lokal/preview tooling: `tilt` med `local` och `kubernetes` modes.
- Namespace‑konvention för previews: `pr-<pr-number>-<short-sha>` eller `agent-<issue>-<id>`.
- Dev‑workflow: `tilt up --namespace $namespace --mode=<local|kubernetes>`.
- K8s‑åtkomst: kontrollera `kubectl cluster-info` och fallback till Teleport om nödvändigt.
- Snabba notiser: `ntfy.sh` kan användas för dev‑notifieringar (exempel: `curl -d "Ephemeral ready: $URL" ntfy.sh/your-topic`).
- CI: GitHub Actions för build, ephemeral deploy och e2e‑körningar.
- Observability: exponera Prometheus‑metrics i ephemeral miljöer när det är möjligt.

Roadmap (högnivå)

Phase 1: MVP

- Auto‑starta en agent per GitHub‑issue
- Worktree/branch automation + PR‑skapande
- Grundläggande implementeringsloop (code → unit tests → CI)
- Ephemeral PR‑environment för manuell testning och preview (tilt lokalt, GH Actions remote)
- Notifieringar till GitHub (issues/PR‑kommentarer) och kostnadsrapport per issue
- Säkerhetsmodell med isolerade konton och miljöer

Phase 2: Post‑launch

- Agenter som periodiskt föreslår nya issues
- Modell‑prioritetslista och provider‑fallback vid kreditbrist
- Automatisk provider‑switching och kontextmigrering
- Browser‑driven UI‑testning i isolerade miljöer
- Multi‑agent orkestreringsdashboard med kostnads‑ och statusvisning
- Förbättrad observability och avancerade policies för när agenter ska be om mänsklig input

Snabbstart

1. Klona repot:

```bash
git clone https://github.com/<ditt-användarnamn>/agents.git
cd agents
```

2. Utforska agent‑mappen:

```bash
ls -la agent-os
```

Installation

Anpassa efter vilket språk/komponent du vill köra. Exempel:

- Go (backend):

```bash
go mod download
```

- Node (frontend):

```bash
npm install
# eller
pnpm install
```

- Python (verktyg/skript):

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

Användning

Exempel (lokal utveckling / demo):

```bash
# Starta tilt för lokala previews
tilt up --namespace pr-demo --mode=local

# Kör backend (Go)
go run ./services/orchestrator

# Kör frontend (Nuxt)
npm run dev --workspace frontend
```

Utveckling

- Kör tester:

```bash
# Go
go test ./...

# Node
npm test
```

- Formattering / lint:

```bash
gofmt -w .
eslint . --fix
black .
```

- Lokala ephemeral deploys: använd `tilt` och följ namespace‑konventionen ovan.

Bidra

1. Forka repot
2. Skapa branch: `feature/<kort‑beskrivning>` eller `bugfix/<nummer>-<kort‑beskrivning>`
3. Skriv tester för ny funktionalitet
4. Skicka PR med tydlig beskrivning och referenser till issues

Se till att tester körs och passerar lokalt och att PR innehåller korta ändringsbeskrivningar.

Licens

Ange lämplig licens i `LICENSE` (exempel: MIT). Se filen `LICENSE` om den finns.

Kontakt

För frågor, buggrapporter eller diskussion: öppna en issue i repot.

Anpassning

Vill du att jag:

- Lägger till badges (CI, coverage)?
- Infogar exakta kommandon/skript från projektets `Makefile` eller `tasks`?
- Skapar en dedikerad `CONTRIBUTING.md` och `ROADMAP.md` från materialet i `agent-os/product`?

Säg vad du vill ha så uppdaterar jag vidare.
