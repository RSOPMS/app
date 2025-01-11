# Naslov Projekta: BugBase

# Člani in številka skupine

Skupina 22:
- Anže Arhar
- Kristjan Kostanjšek
- Nejc Ločičnik

# Povezava do repozitorija

[GitHub link](https://github.com/RSOPMS)

# Naslov URL, kjer je aplikacija dostopna (in morebitni testni prijavni podatki)

[BugBase](http://72.146.53.48/login/)

# Kratek opis projekta

BugBase je aplikacija za vodenje in upravljanje projektov, zasnovana za razvijalce, ki se pogosto soočajo s kompleksnimi projekti in potrebujejo pomoč pri organizaciji.
Poleg registracije in prijave omogoča dodajanje novih projektov in spremljanje njihovega napredka, pri čemer lahko pod vsak projekt dodajamo naloge (issues), ki jih je mogoče povezati z razvojnimi vejami, določiti prioriteto (npr. High, Low) ter spremljati njihovo stanje (npr. Open, Closed).
Poleg tega aplikacija podpira dodajanje komentarjev k posameznim nalogam, kar olajša sodelovanje in komunikacijo znotraj ekip.

# Ogrodje in razvojno okolje

V skopu razvoja, izdelave in zagona aplikacije smo uporabili naslednja programska ogrodja, orodja, razvojna okolja ter storitve.
Celotna aplikacija je napisana v programskem jeziku go.
Za osnovno delovanje storitev smo uporabili le standardno knjižnico.
Dodatno smo namestili štiri knjižnice: [joho/godotenv](https://github.com/joho/godotenv) za branje okoljskih spremenljivk iz `.env` datotek; [lib/pq](https://github.com/lib/pq) sql gonilnik za PostgreSQL podatkovno bazo; [golang-jwt/jwt](https://github.com/golang-jwt/jwt) kot osnovno implementacijo JWT standarda; [nats-io/nats.go](https://github.com/nats-io/nats.go) za pošiljanje in branje sporočil iz NATS strežnika.
Za interaktivno delovanje spletnega dela aplikacije (API klici in preusmeritve uporabnikov) smo uporabili ogrodje [htmx](https://htmx.org/).
Orodje `make` smo uporabili za organizacijo in uporabo pogosto uporabljenih ukazov (postavljanje razvojnega okolja, migracije podatkovne baze, ...).
Z orodjem [air](https://github.com/air-verse/air) smo lokalno poganjali strežnike.
Migracije (in testne podatke) na podatkovni bazi smo upravljali z orodjem [goose](https://github.com/pressly/goose).
Za iskanje napak v aplikaciji smo uporabljali orodje [delve](https://github.com/go-delve/delve).
Vsako storitev smo pretvorili v samostojno datoteko in jo namestili v lasten Docker vsebnik.
S Kubernetesom smo zagotovili visoko razpoložljivost izdelanih vsebnikov.
Lokalno testiranje konfiguracije Kubernetesa znotraj okolja Docker nam je omogočilo orodje [minikube](https://minikube.sigs.k8s.io/docs/).
Za komunikacijo med nekaterimi storitvami smo tudi uporabili [NATS](https://nats.io/).
Za spremljanje osnovnih statistik podatkovne baze smo namestili orodje [Grafana](https://grafana.com/).
Izdelavo `CHANGELOG.md` datoteke, kjer so zapisane ključne spremembe izvorne kode, je omogočilo orodje [git-cliff](https://github.com/orhun/git-cliff).
Vso izvorno kodo smo hranili na platformi GitHub.
Tam smo tudi izvajali zvezno integracijo in zvezno dostavo s GitHub actions.
Izdelane Docker vsebnike smo hranili na GitHub-ovi platformi GitHub Container Registry.
Vse verzije aplikacije smo poganjali na Azure Kubernetes Service.

# Shema arhitekture

![arhitektura](assets/arhitektura.png)
![workflow](assets/workflow.png)

# Seznam funkcionalnosti mikrostoritev

V pričujočem poglavju bomo za vsako mikrostoritev navedli funkcionalnosti, ki jih implementira.

## database

Po meri nadgrajena slika PostgreSQL podatkovne baze.
Ob zagonu se avtomatično namestijo ključne nadgradnje (migracije) podatkovne baze.
Prav tako smo implementirali avtomatično namestitev testnih podatkov.

## app-static

TODO

## app-issue

TODO - Nejc

## app-login

TODO - Nejc

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

## 1. Repozitorij

TODO

## 2. Mikrostoritve in "cloud-native aplikacija"

TODO

## 3. Dokumentacija

TODO

## 4. Dokumentacija API

TODO - Nejc

## 5. Cevovod CI/CD

TODO

## 6. Helm charts

TODO - Nejc

## 7. Namestitev v oblak

TODO

## 8. "Serverless" funkcija

TODO - Nejc

## 9. Zunanji API

TODO - Nejc

## 10. Večnajemništvo

TODO - Nejc

## 11. Preverjanje zdravja

TODO

~~## 12. GraphQL in gRPC~~

## 12. Sporočilni sistemi

TODO

~~## "Event sourcing" in CQRS~~

~~## Centralizirano beleženje dnevnikov~~

## 13. Zbiranje metrik

TODO

## 14. Izolacija in toleranca napak

TODO

## 15. Upravljanje s konfiguracijo

TODO - Nejc

## 16. Grafični vmesnik

TODO

~~## Terraform~~

~~## API Gateway~~

TODO

## 17. Ingress Controller

TODO

~~## IAM, OAuth2, OIDC~~
