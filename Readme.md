# 🧊 Auth-JWT 

### Сервис авторизации с использованием jwt токенов

# ✨ Используемые технологии

> Backend

<table style="width: 100%" >
     <td align="center" width="130" height="90">
      <a href="#">
        <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" width="70" alt="Go" />
      </a>
      <br>Go
    </td>
     <td align="center" width="130" height="90">
      <a href="#">
        <img  src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg" width="70" alt="Postgresql" />
      </a>
      <br>Postgresql
    </td>
     <td align="center" width="130" height="90">
      <a href="#">
        <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/swagger/swagger-original.svg"  width="70" alt="Swagger" />
      </a>
      <br>Swagger
    </td>    
<td align="center" width="130" height="90">
      <a href="#">
        <img  src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg" width="70" alt="Docker" />
      </a>
      <br>Docker
    </td>
   
</table>

> Frontend

<table style="width: 100%" >
    <tr>
    <td align="center" width="110" height="90">
      <a href="#">
        <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/react/react-original.svg"  width="70" alt="React" />
      </a>
      <br>React
    </td>
    <td align="center" width="110" height="90">
      <a href="#">
        <img src="https://raw.githubusercontent.com/devicons/devicon/1119b9f84c0290e0f0b38982099a2bd027a48bf1/icons/typescript/typescript-original.svg" width="70" alt="TypeScript" />
      </a>
      <br>TypeScript
    </td>
     <td align="center" width="110" height="90">
      <a href="#">
        <img  src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vitejs/vitejs-original.svg" width="70" alt="Vite" />
      </a>
      <br>Vite
    </td>       
    <td align="center" width="110" height="90">
      <a href="#">
        <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/redux/redux-original.svg" width="70" alt="Redux" />
      </a>
      <br>Redux
    </td>
   <td align="center" width="110" height="90">
      <a href="#">
        <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/yarn/yarn-original.svg" width="70" alt="Yarn" />
      </a>
      <br>Yarn
    </td>
    <td align="center" width="110" height="90">
      <a href="#">
        <img  src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/eslint/eslint-original.svg" width="70" alt="Eslint" />
      </a>
      <br>Eslint
    </td>
    <td align="center" width="110" height="90">
      <a href="#">
        <img  src="https://brandeps.com/icon-download/P/Prettier-icon-vector-02.svg" width="70" alt="Prettier" />
      </a>
      <br>Prettier
    </td>
<td align="center" width="130" height="90">
      <a href="#">
        <img  src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg" width="70" alt="Docker" />
      </a>
      <br>Docker
    </td>
</tr>

</table>


# 🚀 Быстрый старт

1. Склонируй репозиторий

```bash 
  git clone https://github.com/limona77/Auth-JWT
```
2. Собери мигратор
```bash 
  docker build -t migrator .\migrator
 ``` 
3. запусти postgresql 
```bash 
  docker-compose up postgres 
``` 
4. запусти миграции
```bash 
  Пример моей ссылки для postgres: postgres://postgres:postgres@postgres:5432/auth?sslmode=disable
```
```bash 
docker run  --network host migrator -path=/migrations/  -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/auth?sslmode=disable {up/down 2}
```
5. запусти backend и frontend
```bash
  docker-compose up
```
## 🌐 API документация

### Регистрация

- `POST /auth/register`
- **Запрос**:
  ```json
  {
    "email": "user1@gmail.com",
    "password": "12345"
  }
  ```

- **Ответ**:
  ```json
  {
    "user": {
        "ID": 3,
        "Email": "user1@gmail.com",
        "Password": ""
    },
    "refreshToken": "token",
    "accessToken": "token"
  }
  ```
  

### Логин
- `POST /auth/login`
- **Запрос**:
  ```json
  {
    "email": "user1@gmail.com",
    "password": "12345"
  }
  ```

- **Ответ**:
  ```json
  {
    "user": {
        "ID": 3,
        "Email": "user1@gmail.com",
        "Password": ""
    },
    "refreshToken": "token",
    "accessToken": "token"
  }
  ```

### Обновление токенов
- `GET auth/refresh`
- **Запрос**: body не нужно
- **Заголовок**:
- *Authorization=Bearer accessToken*
- **Ответ**:
  ```json
  {
    "user": {
        "ID": 3,
        "Email": "user1@gmail.com",
        "Password": ""
    },
    "refreshToken": "token",
    "accessToken": "token"
  }
  ```
### Выход из аккаунта
- `GET /auth/logout`
- **Запрос**: body не нужно
- **Cookie**: должен храниться refreshToken
- **Ответ**:
  ```json
  {"userID": 3}
  ```

### Получить свой id
- `GET /me`
- **Запрос**: body не нужно
- *Authorization=Bearer accessToken*

- **Ответ**:
  ```json
  {
    "user": {
        "ID": 3,
        "Email": "user1@gmail.com",
        "Password": ""
    }
  }
  ```