#Needs the following to be exported in Env so the the SP to be able top access KV

```
export AZURE_TENANT_ID="<active_directory_tenant_id"
export AZURE_CLIENT_ID="<service_principal_appid>"
export AZURE_CLIENT_SECRET="<service_principal_password>"
export TEAMS_WEBHOOK_URL="some teams webhook URL"
export REDIS_HOST="<dedis host | default loclhost>"
export REDIS_PORT= <port | default to 6379>"
export REDIS_PASSWORD = "<password | default to no password>"
```
