# ya_individual_go_backend -> [Dokumentation av appen](https://github.com/tallner/ya_individual_go_backend/blob/main/Dokumentation.md)

### Webtjänster
- [x] Två stycken; ena är skriven i Go  eller Python
- [x] En webbsajt
- [x] Databasdrivna (förslag kör MySQL i ert kluster)
	- [x] Dessa ska köras i eget Kubernetes-kluster
	- [x] set up controller och DNS-subdomains
	- [x] manuell deploy med kubectl rollout restart deployment

### Inlämning
- [x] Kod i privata git-repository
- [x] bjud in stefan.holmberg@systementor.se som collaborator
- [x] Skriftlig dokumentation (ca 3 sidor) – beskriv kortfattat appliktionernas arkitektur

### Godkänt
- [x] Källkod i valfri Cloud-leverantör (Github, Dockerhub, Linode)

##### Continous integration:
- [x] Kompilering
- [x] Enhetstester 
- [x] skapa container 
- [x] push container -> container repository

### Väl godkänt
- [x] Minst ett till CRUD API i andra språket....njaaa
  - [x] Som databas ska du välja nåt annat än det ovan (ex om MySQL ovan, MS SQLServer)...njaa här med 
- [x] Frontend
- [ ] Databas backup -> S3 lagring – tex backblaze.com
https://www.youtube.com/watch?v=VtVaK5BE52k
