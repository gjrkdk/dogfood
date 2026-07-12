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
