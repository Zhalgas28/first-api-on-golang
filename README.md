# Rest API for works with News on Go

## Project covers the following technologies:
- Work with framework <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Work with PostgreSQL with <a href="https://github.com/jmoiron/sqlx">sqlx</a> library and writed SQL queries.
- Application configuration using <a href="https://github.com/ilyakaznacheev/cleanenv">ilyakaznacheev/cleanenv</a>, Work with environment variables.
- Registration and Authentication, Work with JWT token, Middleware.
- Running application and Postgres with Docker, Docker-Compose.

---
### To run an application:
```
make build && make run
```
### If app is running for the first time:
```
make migrate
```
