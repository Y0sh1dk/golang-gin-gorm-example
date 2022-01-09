# golang-gin-gorm-example

## TODO:

 - Add server config
 - Logging (To DB?)
 - Seed DB with data
 - Build docker image (from makefile)
 - Different environments/configs
 - Input validation
 - User authentication (jwt)
 - CI/CD
 - IaC w/ Terraform

    select pid as process_id,
        usename as username,
        datname as database_name,
        client_addr as client_address,
        application_name,
        backend_start,
        state,
        state_change
    from pg_stat_activity;