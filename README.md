# Medods Interview Assignment

**Technologies used:**

- Go
- JWT
- PostgreSQL

**Task:**

Write a part of the authentication service.

Two REST routes:

- The first route issues a pair of Access, Refresh tokens for the user with the identifier (GUID) specified in the request parameter
- The second route performs a Refresh operation on a pair of Access, Refresh tokens

**Requirements:**

An access token of the JWT type, the SHA512 algorithm, is strictly prohibited from being stored in the database.

Refresh token is an arbitrary type, base64 transfer format, stored in the database exclusively as a bcrypt hash, must be protected from client-side changes and reuse attempts.

Access, Refresh tokens are mutually related, the Refresh operation for an Access token can only be performed with the Refresh token that was issued with it.

The payload of tokens must contain information about the IP address of the client to whom it was issued. If the ip address has changed, during the refresh operation, you need to send an email warning to the user's email (for simplification, you can use the mock data).

**Result:**

The result of the task must be provided as source code on Github. It will be a plus if you can use Docker and cover the code with tests.

P.S. Friends! Tasks completed in whole or in part using chatGPT are immediately visible. If you are not ready to solve this test task on your own, then please, let's appreciate each other's time and not even try :)

**Instructions:**

You need to create a .env file in the root folder

```env
SYS_NAME=APITITLE
SECRET_KEY=KEY
ENV=DEV

POSTGRES_USER=pg
POSTGRES_PASSWORD=pass
POSTGRES_DB=db

SERVER_ADDRESS=0.0.0.0:8000
POSTGRES_CONNECTION=postgres://pg:pass@db:5432/db

PGADMIN_DEFAULT_EMAIL=mzinyoni7@mail.com
PGADMIN_DEFAULT_PASSWORD=mzinyoni7
```

`docker compose -f "docker-compose.yml" up -d --build`
