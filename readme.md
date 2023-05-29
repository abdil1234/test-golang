##GOLANG TEST

# Installation
* git clone repo-url
* docker-compose up
* app running on port 9010

# Configuration

```mysql
CREATE USER 'root'@'%' IDENTIFIED BY 'root';
GRANT ALL ON *.* TO 'root'@'%';
FLUSH PRIVILEGES;

```

# Game Post

    POST http://localhost:9010/golang-test/api/v1/game

<img width="500" alt="Screen Shot 2023-05-29 at 17 45 40" src="https://github.com/abdil1234/test-golang/assets/31970269/481c4a23-432e-46a9-84dd-362756fe1c5b">


# Game List

    GET http://localhost:9010/golang-test/api/v1/games

<img width="500" alt="Screen Shot 2023-05-29 at 17 44 56" src="https://github.com/abdil1234/test-golang/assets/31970269/163e9e72-20fe-4650-b596-b399a3886e9b">
