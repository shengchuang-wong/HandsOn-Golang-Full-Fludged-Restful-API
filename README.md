Credited to https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121?source=friends_link&sk=13d172e0aec3b14f758c0968a8798145

#### Basic Command

- docker-compose up
- docker-compose up -d <<< run in the background
- docker-compose up --build <<< to build again
- docker-compose down
- docker-compose down --remove-orphans --volumes <<< also remove the volumes

- docker ps -a
- docker container stop [container_id]
- docker container prune
- docker system prune -a <<< remove all unused images, not just dangling ones
- docker system prune -a --volumes <<< remove all unused images not just dangling ones and volumes

- docker image ls
- docker image rm [image_id]
- docker image prune

#### Tests

- docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
- docker-compose -f docker-compose.test.yml down

