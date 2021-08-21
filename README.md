# sirclo-technical-test-backend

Sirclo Technical Test - Backend.
Created and initiated by @judahsuryaputra

## Running on Local

-   Make sure you have installed go locally, check [here](https://golang.org/doc/install)
-   Make sure you have installed postgresql locally, check[here](https://www.postgresql.org/download/)
-   Make sure you have created a local database named `sirclo-dev`
-   run `go run main.go`
-   I provide the requests that I use to test my code through Insomnia named `sirclo-backend-requests-INSOMNIA`.

-   Query that I use for this project.
-   ```
    CREATE TABLE weights(
        id SERIAL PRIMARY KEY,
        max int,
        min int,
        difference int,
        date timestamp
    );
    ```