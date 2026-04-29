---
allowed-tools: Bash
description: Kör git diff, commitar ändringar med ett lämpligt meddelande och pushar till remote
---

# Push

## Syfte

Detta kommando kör git diff för att visa vilka ändringar som gjorts, skapar ett lämpligt commit-meddelande baserat på ändringarna, commitar dem och pushar till remote. Använd när du vill stage:a, commita och pusha alla ändringar i ett steg.

## Variabler

 Inga dynamiska variabler - kommandot hanterar allt internt

## Workflow

1. Kör `git diff` för att se vilka ändringar som finns i working directory
2. Kör `git diff --staged` för att se redan stagade ändringar
3. Analysera outputten från git diff för att förstå vilka filer som ändrats
4. Kör `git status --short` för en简洁 översikt av ändringar
5. Skapa ett lämpligt commit-meddelande baserat på de ändrade filerna och typen av ändringar
6. Kör `git add -A` för att staga alla ändringar
7. Kör `git commit -m "<meddelande>"` med det skapade meddelandet
8. Kör `git push` för att pusha till remote
9. Rapportera resultatet till användaren

## Relevanta filer

- `.git/` (för att förstå git-konfiguration om det behövs)

## Exempel

```bash
git diff
git status --short
git add -A
git commit -m "Update feature: add user authentication"
git push
```

## Rapport

Kommandot ska svara med:
- Sammanfattning av ändringar som hittades (filer ändrade, tillagda, borttagna)
- Det commit-meddelande som användes
- Bekräftelse på att commit och push lyckades
- Eventuella varningar om inga ändringar hittades