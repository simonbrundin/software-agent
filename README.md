# Software Agent

## Tech Stack

- Commands: Markdown
- Workflows: Python
- Frontend: Nuxt (Vue)
- Backend: Go
- Databas: PostgreSQL
- ORM: Drizzle
- Deployment / Orkestrering: Kubernetes
- CI (GitHub Actions)

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
