# Software Agent

## Tech Stack

- Commands: Markdown
- Workflows: Python
- Frontend: Nuxt (Vue)
- Backend: Python
- Databas: PostgreSQL
- ORM: Drizzle
- Deployment / Orkestrering: Kubernetes
- CI (GitHub Actions)

## Kör igång dev

```bash
cd environments/dev
tilt up
```

Detta startar alla tjänster via Docker Compose:

- **PostgreSQL** på port 5432
- **Hasura** (GraphQL) på port 8080
- **Backend** (Python/FastAPI) på port 8000
- **Frontend** (Nuxt) på port 3000

För att stänga av: `tilt down`

---

## Databas

Schema Drawing:
https://www.drawdb.app/editor/diagrams/8b11f8df-9402-4654-931c-5f20f78fb253

## Detta händer när ett nytt meddelande kommer in

1. Spara meddelande i databas
   - MeddelandeID
   - Meddelande
   - KonversationsID
   - Tid
   - RepoURL
2. Bestäm vilka workflows som ska köras - **workflow-selector**
   - Behövs ny feature? - **Skriv ny feature**
   - Behövs nytt scenario? - **Skriv ny scenario**
   - Behövs en plan? - **Skapa plan**
     - plan_feature
     - plan_bug
     * plan_chore
   - Behöver du bara svara på en fråga?
3. Skapa nytt worktree
4. Kör workflows som selector bestämt

### Vilka workflows behövs?

#### SDLC

- Skriv tester
  - Ny feature
  - Nytt scenario
- Planera förändringar
- Implementera förändringar
- Refactor loop
  - While not green
    - Kör tester
    - Refactor både kod och step definitions
- Review
  - Skapa PR
- Dokumentera
- Deploy
