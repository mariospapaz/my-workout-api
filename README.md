# my-workout-api
https://hub.docker.com/r/mariospapaz/my-workout

## Plan

![image](https://user-images.githubusercontent.com/30930688/173666083-5b4dbfca-bf83-46df-b38a-17c2c6a0c857.png)

## Installation Guide

**for servers**

`docker pull mongodb:5.0` (hostname must be named `bibi`, api will not work without this spec and a mongodb image)

`docker pull mariospapaz/api-workout:1.0`

You can either run them in your own, or take reference my `docker-compose.yaml`


**for users that have docker installed**

`$ git clone https://github.com/mariospapaz/uni_scheduler.git `

`$ docker compose up`
