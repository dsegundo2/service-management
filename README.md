# Service-Management
API for keeping track of various Services

### Database
- Uses a postgres database
- To create a postgres container, use the following command
```
docker-compose up -d
```
- (Optional) To Connect to the server in the command line.
```
docker exec -ti service-management-postgres psql -U postgres -d service_management
```