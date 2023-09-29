# go-cart
An e-commerce multiple microservices application build using Go-Gin framework and MongoDB.

# Microservices:
    - user-management:
        - Exposing Rest APIs to manage user registration, jwt token based authentication 
          and profile management.
    - product-management:
        - Exposing Rest APIs to manage product offerings.
    - order-management: [In-Development]
        - Exposing Rest APIs to manage cart and order.

# Run application:
    - execute make commands from project root to start/stop docker containers.
        - 'make up' to launch all microservices
        - 'make down' to stop all docker containers.

        Note 'order-management' service not supported by docker compose yet.
