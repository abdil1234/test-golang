##GOLANG TEST

# Installation
* git clone repo-url
* docker-compose up
* app running on port 9010

# Configuration

Enter into mysql container

```
docker exec -it database-mysql bash
```

Run query

```mysql
CREATE USER 'root'@'%' IDENTIFIED BY 'root';
GRANT ALL ON *.* TO 'root'@'%';
FLUSH PRIVILEGES;

```

# Game Post

    POST http://localhost:9010/golang-test/api/v1/game

<img width="500" alt="Screen Shot 2023-05-30 at 01 18 41" src="https://github.com/abdil1234/test-golang/assets/31970269/67dfb8fc-0f09-488f-b372-0d647f2fdb9c">



# Game List

    GET http://localhost:9010/golang-test/api/v1/games

<img width="500" alt="Screen Shot 2023-05-30 at 01 19 34" src="https://github.com/abdil1234/test-golang/assets/31970269/e28fcff9-438a-448d-9345-cad44cb19ee9">

