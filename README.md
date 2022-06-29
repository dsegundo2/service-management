# Service Management

---
## Task
  - Implement a services API that can be use to implement this dashboard widget. It should support:
    - Returning a list of services. Support filter sorting, pagination
    - Fetching a Particular Service. This Including a method for retrieving its versions.

## Local Use (run in order)
  - Export the listed environment variables listed in `configs/example` or run `source configs/example.env`
  - `make resources` command will create a postgres database with the sql migration file in `cmd/migration`. It will also start up a server for the swagger endpoint documentation.
    - The swagger endpoint documentation is located at [localhost:8082](localhost:8082). If you already have something running on that server, you can change the port in the `docker-compose.yaml` file
    - Note that the swagger interactive functionality is not working yet for this appliation and is only meant for documentation at this time
    - The current swagger-ui image only supports `amd64` chips. If you are using local machine with `arm64`, this image might work for it `spryker/swagger-ui:v3.24.3`, However, I have not confirmed this
  - If you want to run the database service tests, run `make test`
  - After exporting envs and starting the database, you can run the service locally with `make run`
  - For an api client to test the service with, import the `openapi-doc.yaml` in to [insomnia](https://insomnia.rest/download). To import:
    - Open Insomnia Dashboard Window -> Create -> Import From File -> Select `openapi-doc.yaml` in the root of this repo

---

## Design Decisions

### Database
  - I used a postgres database since I think it would help with the need to filter & sort efficiently. Also there is a relational association between the two structs
  - Instead of writing manual sql queries, I used Gorm for a database manager to save time
  - A `ServiceDB` struct implements a database interface and contains a gorm database. This will allow the gorm database to be dependency injected in to the database package, and it will ease unit testing. Additionally, it would make it easier to switch out the database manager if we decided to.
  - There is a `docker-compose.yaml` to set up a postgres database locally for testing. There is also a makefile command `create-database` to help.
  - The database is migrated each time the service is started or tested. This might not be ideal for the long term, but it did help with rapid development of the database package.
  - The Automigrate function is being used, which will make any schema rolebacks more work. However, at the time this should work.
  - There are database integration tests that can be ran with `make test`
  - In order to handle the total number of versions for a Service, there is a database hook that queries the amount each time one is queried from the database. This does result in slower read times and might need to be addressed in the future. However it does keep the database accurate and saves on memory and write time to the database.

### API
  - Structure follows the [Google API Design Guide](https://cloud.google.com/apis/design/standard_methods)
  - Uses an openapi spec to declare what methods are exposed through the api. This was done to try and save time while writing the models, documentation and creating a file for insomnia import.
    - I went back and forth on taking this approach and using the database models for the api as well. One downside to this would have been losing the future flexibility of using different fields in the database and the UI. However, by using different structs, I did have to write conversion functions in the rest/models package, which took more time
  - Not only did the openapi doc allow for the generation of the api struct models, but it also allowed for [method documentation generation](/openapi/docs/). However, I think it these documents are not incredibly clean and it would be worth more time in the future to fine tune the documentation generation for easy api consumption.

---

### Known Issues
  - The DeletedAt Time will return with the default time when it should be omitted from the response messages
  - Adding a service version takes in the same ID in both the url and the query params, and they must match
  - Unsafe marshalling in the handlers. The methods should function correctly, but the handler should ideally use the correct incoming structures moving forward

---

### Future Implementations Issues to Consider
  - A major one is that authorization has not been addressed. A user service should be integrated with the API
  - Security around printing the database password should be handled and probably have the value stored in some kind of secure secret manager
  - Set up a swagger docs server so that it interacts with the applicaiton
  - It would be helpful to add more logs that have a granular logging level. That way we can limit the amount of logs until we need to debug
  - There is currently no update Service Version endpoint. This would be a useful function in the future, but for now, the user can Remove the current version and add a new updated version if necessary.
  - Better format the error messages returned so the user can marshall them more effectively
  - Add smoke tests that test the server instead of just manually testing through insomnia

