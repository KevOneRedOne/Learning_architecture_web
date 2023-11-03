# Web Architecture Learning Project

## Objective

The main objective of this project is to provide a practical learning experience about web architecture. It demonstrates the evolution of a web application's architecture, starting from a monolithic client/server 3-tier design and transitioning to a more scalable microservices-based architecture.

## Technologies Used

- **Backend**:
  - Server: Go with the Gin framework for the API
  - Additional API using Fastify for element searching

- **Frontend**:
  - Web client: Vue.js with Nuxt for server-side rendering

- **Database**:
  - PostgreSQL with GORM for the Object-Relational Mapping (ORM)

## Project Evolution

### Transition from Monolith to Client/Server 3-Tier

In this initial phase, the project will focus on transforming the existing monolithic architecture into a structured client/server 3-tier architecture. The main components involved are:

- **Backend**: A Go server built with the Gin framework to handle API requests.
- **Frontend**: Vue.js with Nuxt for server-side rendering.
- **Database**: PostgreSQL with GORM used for ORM.

### Transition from Client/Server 3-Tier to Microservices

The subsequent phase will involve transitioning from a client/server 3-tier architecture to a more scalable and flexible microservices architecture. Key changes in this phase include:

- **Dockerization**: Each service will be containerized using Docker, allowing for easier deployment and scaling.
- **Frontend**: The Vue.js with Nuxt setup for server-side rendering remains consistent.
- **Backend**: Go with the Gin framework for the primary API and Fastify for a secondary API dedicated to element searching.
