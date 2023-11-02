# Learing_architecture_web

## Objectif :
- 1. Apprendre les bases de l'architecture web
- 2. Passer d'une architecture monolite en client/serveur 3 tiers
    - 2.0. Technologie au choix
    - 2.1. Créer un serveur web avec Go
    - 2.2. Créer un client web avec Vue.js
    - 2.3. Créer une base de données avec PostgreSQL
- 3. Optimiser l'architecture pour la rendre scalable
- 4. Passer d'une architecture client/serveur à une architecture en microservices
    - 4.1 Rajouter une api coté backend avec Fastify

### Passage du monolite au client/serveur 3 tiers :
- Backend : Go avec le framework Gin pour l'API
- Frontend : Vue.js avec Nuxt pour le rendu côté serveur
- Database : PostgreSQL avec GORM pour l'ORM

### Passage du client/serveur 3 tiers au microservice :
- Docker : Conteneurisation de chaque service
- Frontend : Vue.js avec Nuxt pour le rendu côté serveur
- Database : PostgreSQL avec GORM pour l'ORM
- Backend :
    - Go avec le framework Gin pour l'API
    - Fatify pour une deuxième API pour la recherche d'éléments
