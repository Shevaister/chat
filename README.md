This is a chat build with docker, nginx, mysql, go and vue.

Chat has authorization signin/signup, using jwt for security, has user search, freindlist page and friend add/remove possibility and block user opportunity, chat use websocket to provide real-time chat connection, you can also change your login/password/avatar.

To download the project use "git clone https://github.com/Shevaister/chat"

Project start and setup:
    Start project for development:
        Set up all .env files and nginx.conf.

        Start all parts of server:
        nginx: "start nginx"
        go: "go run main.go"
        vue: "npm run serve"

    Start project:
        You should delete all .env files except docker-compose.yml' .env file and delete go config loader "go/config" and delete its import in main.go. 

        To start project as docker network u need to build all images:
        nginx: "docker build -t nginx ."
        go: "docker build -t backend ."
        vue: "docker build -t frontend ." 

        After image building u need to set up docker-compose.yml file.

        After setuping you need to write "docker-compose up" in console, and chat starts after that.