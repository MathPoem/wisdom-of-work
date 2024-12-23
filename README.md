![alt text](./assets/image.png)
# wisdom-of-work

This is a client and a protected with the Proof of Work protocol TCP server.

Hashcash chosen as a proof-of-work algorithm as it is a simple and effective way to prevent spam and DDoS attacks.

Hashcash is a proof-of-work system that requires a client to solve a cryptographic puzzle before it can send data to the server.
___
## **Usage**
#### Enshure you have [*docker desktop*](https://www.docker.com/products/docker-desktop/) installed
___
### 👜 👜 👜 *copy .env file* 👜 👜 👜
```shell
cp ./server/.env.example ./server/.env
```
```shell
cp ./client/.env.example ./client/.env
```
___
### 🚒🚒🚒 *create SERVER container* 🚒🚒🚒
```shell
docker compose up -d server
```
___
### 🔌 🔌 🔌 *create CLIENT container* 🔌 🔌 🔌
```shell
docker compose up -d client
```
___
### 🎬🎬🎬 *look at the logs* 🎬🎬🎬
```shell
docker compose logs -f server
```
```shell
docker compose logs -f client
```