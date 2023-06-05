# ya_individual_go_backend -> [Dokumentation av appen](https://github.com/tallner/ya_individual_go_backend/blob/main/Dokumentation.md)

### Webtjänster
- [ ] Två stycken; ena är skriven i Go  eller Python
- [ ] En webbsajt
- [ ] Databasdrivna (förslag kör MySQL i ert kluster)
	- [ ] Dessa ska köras i eget Kubernetes-kluster
	- [ ] set up controller och DNS-subdomains
	- [ ] manuell deploy med kubectl rollout restart deployment

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
- [ ] Minst ett till CRUD API i andra språket
  - [ ] Som databas ska du välja nåt annat än det ovan (ex om MySQL ovan, MS SQLServer)
- [ ] Frontend
- [ ] Databas backup -> S3 lagring – tex backblaze.com
https://www.youtube.com/watch?v=VtVaK5BE52k
