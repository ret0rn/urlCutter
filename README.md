# urlCutter

A simple web API for link shortening

## Set Up


### Docker Compose	<img src="https://www.svgrepo.com/show/353659/docker-icon.svg" alt="generate short url" width="20">

```
dokcer-compose up
```
### Manual	<img src="https://www.svgrepo.com/show/144271/settings.svg" alt="generate short url" width="20">


1. edit database config in `./configs/database.json`
2. use `psql -U postgres -f db.sql` for create database
3. use `go run .` to start app

## Usage

### Request

<img src="https://i.imgur.com/AWL5NTs.jpg" alt="generate short url">

### Result
<img src="https://i.imgur.com/lLhT20j.jpg" alt="redirect to long  url">