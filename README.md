# Go Rest Api Template

# Migrations


# Locale Development 

## Docker Compose Local environments
Docker Compose Local file contains the following environment variables:

* `POSTGRES_USER`
* `POSTGRES_PASSWORD`
* `PGADMIN_DEFAULT_EMAIL`
* `PGADMIN_DEFAULT_PASSWORD`

## Access to postgres:
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** test_password (as a default)

## Access to PgAdmin:
* **URL:** `http://localhost:5050`
* **Username:** admin@service.example (as a default)
* **Password:** admin (as a default)

## Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`
* **Password** as `POSTGRES_PASSWORD`