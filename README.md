# circleci-playground
Repository for hands-on exploration and experimentation with CircleCI. Learn Continuous Integration (CI) workflows, test automation, and deployment strategies in a sandbox environment

The project is a simple application That responds `Hello World, The secret env is $SECRET_KEY` on hitting / endpoint on 8080. Where `SECRET_KEY` is an environment variable.

The circleci config will:
- Run tests
- Build the image
- Push to dockerHub     


## CircleCi config
```yaml
version: 2.1

jobs:
  build:
    docker:
      - image: cimg/go:1.21.8  # Using a CircleCI maintained Go image

    working_directory: ~/circleci-playground  # Set the working directory

    steps:
      - checkout # checks out the source code of the project into the working directory

      # Run Go tests
      - run:
          name: Run Go tests
          command: go test -cover ./...  

      # Use setup_remote_docker to set up a remote Docker environment
      - setup_remote_docker

      # Build Docker image (Make sure Dockerfile is in the root folder of the project)
      - run:
          name: Build Docker image
          command: docker build -t breeze5690/hello-world:latest . 

      # Login to Docker Hub using environment variables
      - run:
          name: Login to Docker Hub
          command: docker login --username $DOCKER_USERNAME --password $DOCKER_PASSWORD

      # Push Docker image to Docker Hub
      - run:
          name: Push Docker image to Docker Hub
          command: docker push breeze5690/hello-world:latest

workflows:
  version: 2
  build:
    jobs:
      - build

```