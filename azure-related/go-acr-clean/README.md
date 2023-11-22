## We need .env file with the following

```
ACR_USERNAME=<ACR username>
ACR_PASSWORD=<ARC password>
ACR=full name of the ACR without https like acrname.azurecr.io
```

## Usage

# For help
go-acr-clean help

# options
```
go-acr-clean or go-acr-clean.exe --list-all-repos   to list all repositories
go-acr-clean or go-acr-clean.exe --list-repo-tags   <repository name> list all tags fopr the repository
go-acr-clean or go-acr-clean.exe --delete-tags --repo <repository name> --start-tag <the oldest tag to delete> --end-tag <the newest tag to delete>
```