#Needs the following to be exported in Env so the the SP to be able top access KV
##
```
export AZURE_TENANT_ID="<active_directory_tenant_id"
export AZURE_CLIENT_ID="<service_principal_appid>"
export AZURE_CLIENT_SECRET="<service_principal_password>"

#normally the below have defaults
export REDIS_HOST="some_host"  # default if missing to  localhost 
export REDIS_PORT="6379"    #default if missing to 6379
export REDIS_PASSWORD=""    # default if missing to ""
```