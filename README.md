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

#### Clean up command
- https://gist.github.com/bastman/5b57ddb3c11942094f8d0a97d461b430
- https://linuxize.com/post/how-to-remove-docker-images-containers-volumes-and-networks/

#### Tests

- docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
- docker-compose -f docker-compose.test.yml down

The tests is currently not working on docker. Possible solution: Issue fixed. I managed to find my config setting from my previous GOGS install. Since I’m using the docker version of GOGS, instead of the MySQL host being 127.0.0.1:3306, it should be 172.17.0.1:3306 since 172.17.0.1 being the bridge gateway IP for the container.

My guess is that the default settings for first time install assumes that MySQL and GOGS are both installed on the host server.

Credited to https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121?source=friends_link&sk=13d172e0aec3b14f758c0968a8798145