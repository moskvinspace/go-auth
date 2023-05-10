# simple-web-app

## Run app

```sh
docker compose up -d
```

### [API] POST Sign-up

```sh
http://localhost:8080/api/sign-up
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

### [API] POST Sign-in

```sh
http://localhost:8080/api/sign-in
```
```json
{
	"email": "john@example.com",
	"password": "Qwerty!23"
}
```
