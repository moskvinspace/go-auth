# simple-web-app

## Run app

```sh
docker compose up -d
```

### [API] POST Register

```sh
http://localhost:8080/api/register
```
```json
{
	"first_name": "John",
	"last_name": "Doe",
	"email": "john@example.com",
	"password_1": "Qwerty!23",
	"password_2": "Qwerty!23"
}
```

### [API] POST Login

```sh
http://localhost:8080/api/login
```
```json
{
	"email": "john@example.com",
	"password": "Qwerty!23"
}
```
### [API] GET Logout

```sh
http://localhost:8080/api/logout
```

### [API] GET CurrentUser

```sh
http://localhost:8080/api/current_user
```
