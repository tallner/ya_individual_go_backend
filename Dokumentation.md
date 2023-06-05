## Min lösning:  
Jag gör inget specifikt use-case, bara en applikation som uppfyller kraven(tror jag :)) och CRUDar sin databas.

## I denna appen gör jag följande
-  backend i go, 
-  MySQL databas 
-  github källkodshantering
-  gihub actions för cicd
-  docker hub som repository
-  Lindode cloud provider  

## Konfigurering och sätta upp applikationerna
1. Linode -> create cluster, manuellt inuti Linode
2. [Create namespace -> semimanuellt, yaml fil](https://github.com/tallner/Linodekluster/blob/master/01-namespace.yaml)
3. Create [nginx](https://github.com/tallner/Linodekluster/blob/master/60-ingress.yaml) [ingresscontroller](https://github.com/tallner/Linodekluster/blob/master/50-ingresscontroller.yaml) -> semimanuellt, yaml fil
4. [Create pvc -> semimanuellt, yaml fil](https://github.com/tallner/Linodekluster/blob/master/20-pvc.yaml)
5. [Create mysql -> semimanuellt, yaml fil](https://github.com/tallner/Linodekluster/blob/master/30-mysql.yml)
6. Create the application in a yaml file specific for the application and deploy with github actions workflow -> automatiskt, [workflow](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml)+[yaml](https://github.com/tallner/ya_individual_go_backend/blob/main/linodedeploy.yaml)

## Gihub actions workflow
1. [Triggas av en push till main från IDE.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L7)  
2. [Sätter upp ett jobb på en Ubuntu runner.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L13)
3. [Checkar ut repot så så att runner kan bygga koden och köra workflown.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L22)
4. [Get commit SHA tp show on webpage.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L27)
5. [Kör unit tester, avbryter om dom failar.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L32)
6. [Loggar in på docker hub med creds som är secrets i github.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L45)
8. [Sätter context med KUBE_CONFIG secrets, dvs Linodes inloggningscreds osv, så att kubectl kan köras sen.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L51)
9. [Docker image byggs och pushas till docker hub.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L57)
10. [DOCKER_CONFIG variablen sätts och sparas som env var i workflown.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L62)
12. [Seddar in värdena på secrets till deploy.yaml.](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L64)
13. [kubectl apply på deployfilen](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L71)
14. [send email att error (does not work yet)](https://github.com/tallner/ya_individual_go_backend/blob/main/.github/workflows/workflow.yml#L73)

## [Deployfilen](https://github.com/tallner/ya_individual_go_backend/blob/main/linodedeploy.yaml)
1. [Skapar secret regcred2 som används för att accessa imagen på docker hub](https://github.com/tallner/ya_individual_go_backend/blob/main/linodedeploy.yaml#L3)
--
2. [Skapar en deployment, pod, hämtar imagen från dockerhub med secret regcred2](ttps://github.com/tallner/ya_individual_go_backend/blob/main/linodedeploy.yaml#L13)
--
3. [kapar en service med intern IP och vidarebefodrar trafik från port 80 till 8080 på poddar som matchar selektorn "app", dvs mygobackend. Detta gör att appen kan nås inom klustret via den interna IP-adressen och port 80.](https://github.com/tallner/ya_individual_go_backend/blob/main/linodedeploy.yaml#L41)

