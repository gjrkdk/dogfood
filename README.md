# dogfood

Dit is de dogfooding-repo voor **Polder**, een background coding agent die
zelfstandig taken uitvoert en de resultaten aanbiedt als pull requests.

## Doel

Deze repo dient als testomgeving om Polder in de praktijk te gebruiken op
zijn eigen ontwikkeling: taken worden hier uitgevoerd zoals een gebruiker
dat ook zou doen, zodat problemen en verbeterpunten vroeg aan het licht
komen.

## Hoe het werkt

1. Polder krijgt een taak toegewezen.
2. De taak wordt uitgevoerd op een wegwerp-VM: een tijdelijke, geïsoleerde
   omgeving die na afloop wordt weggegooid.
3. Wijzigingen worden gecommit op een branch.
4. De control plane pusht deze branch en opent een pull request, zodat de
   wijzigingen beoordeeld kunnen worden voordat ze worden samengevoegd.

## demo-app (`app/`)

Een kleine Next.js-app met Postgres, bedoeld om te bewijzen dat Polder tegen
een **draaiende** database kan werken (v1.3 Environments).

```sh
docker compose up -d --wait      # Postgres, host-poort 55432 (niet 5432: vaak bezet)
cd app && npm ci
DATABASE_URL="postgres://demo:demo@localhost:55432/demo" npm test
```

De tests praten met een echte database — zonder container falen ze. Dat is
opzet: het bewijst dat de omgeving klopt voordat er code gemerged wordt.
