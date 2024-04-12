#Install pre-requirements
  ##Postgres 
  ```dcoker-compose.yml
  version: '3'
  services:
    db:
      image: postgres
      container_name: postgres
      ports:
        - '5432:5432'
      environment:
        - POSTGRES_USER=postgres
        - POSTGRES_PASSWORD=password
        - POSTGRES_DB=Product
      volumes:
        - "pgdata:/var/lib/postgresql/data"
      networks:
        - db
      restart: always

      networks: 
        - db

  networks:
    db:
      external: true
  
  volumes:
    pgdata:
  ```

#Config Project
  ##Create .env file
  ```.env
  DB_HOST=127.0.0.1
  DB_PORT=5432
  DB_USER=postgres
  DB_PASSWORD=password
  DB_NAME=Product
  ```

#Build and Run 
``` go build && ./product```
