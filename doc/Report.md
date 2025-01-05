# Naslov Projekta: BugBase

# Člani in številka skupine
- Anže Arhar
- Kristjan Kostanjšek
- Nejc Ločičnik
- Skupina 22

# Povezava do repozitorija
[Github link](https://github.com/RSOPMS)

# Naslov URL, kjer je aplikacija dostopna (in morebitni testni prijavni podatki)
TODO: Link do aplikacije. Aplikacija naj bo dostopna vsaj med zagovorom.

# Kratek opis projekta
BugBase je aplikacija za vodenje in upravljanje projektov, zasnovana za razvijalce, ki se pogosto soočajo s kompleksnimi projekti in potrebujejo pomoč pri organizaciji. Poleg registracije in prijave omogoča dodajanje novih projektov in spremljanje njihovega napredka, pri čemer lahko pod vsak projekt dodajamo naloge (issues), ki jih je mogoče povezati z razvojnimi vejami, določiti prioriteto (npr. High, Low) ter spremljati njihovo stanje (npr. Open, Closed). Poleg tega aplikacija podpira dodajanje komentarjev k posameznim nalogam, kar olajša sodelovanje in komunikacijo znotraj ekip.

# Ogrodje in razvojno okolje
TODO: Navedite tehnologije, programska ogrodja in razvojna okolja, ki so bila uporabljena (npr. KumuluzEE, Spring Boot, Node.js, Docker, Kubernetes).

# Shema arhitekture
TODO: Priložite končno shemo arhitekture aplikacije, ki prikazuje interakcije med vsemi elementi (mikrostoritve, podatkovne baze, zunanja orodja in API-ji, ...). V shemi označite tudi uporabljene komunikacijske protokole.

# Seznam funkcionalnosti mikrostoritev
V pričujočem poglavju bomo za vsako mikrostoritev navedli funkcionalnosti, ki jih implementira.

## database
TODO

## app-static
TODO

## app-issue
TODO

## app-login
TODO

## app-bulk
- prejemanje večje količine podatkov (novih projektov in nalog) preko POST zahtevka
- pošiljanje prejetih podatkov mikrostoritvi app-ingress preko sporočilnega sistema NATS

## app-ingress
- prejem podatkov (novih projektov in nalog) od mikrostoritve app-bulk preko sporočilnega sistema NATS
- vstavljanje prejetih podatkov v podatkovno bazo

# Primeri uporabe
TODO: Navedite seznam primerov uporabe, ki jih aplikacija podpira. Opišite tudi en kompleksnejši primer uporabe, kjer pri obdelavi sodeluje več mikrostoritev.

# Seznam vključenih zahtev
TODO: Za vsako zahtevo na kratko (okvirno do 500 znakov) opišite kako ste zahtevo implementirali/naslovili. Lahko vključite tudi slike. (Jaz sem napisal kar vse, treba je pobrisat tiste, ki jih nimamo).

## Repozitorij
TODO

## Mikrostoritve in "cloud-native aplikacija"
TODO

## Dokumentacija
TODO

## Dokumentacija API
TODO

## Cevovod CI/CD
TODO

## Helm charts
TODO

## Namestitev v oblak
TODO

## "Serverless" funkcija
TODO

## Zunanji API
TODO

## Večnajemništvo
TODO

## Preverjanje zdravja
TODO

## GraphQL in gRPC
TODO

## Sporočilni sistemi
TODO

## "Event sourcing" in CQRS
TODO

## Centralizirano beleženje dnevnikov
TODO

## Zbiranje metrik
TODO

## Izolacija in toleranca napak
TODO

## Upravljanje s konfiguracijo
TODO

## Grafični vmesnik
TODO

## Terraform
TODO

## API Gateway
TODO

## Ingress Controller
TODO

## IAM, OAuth2, OIDC
TODO
