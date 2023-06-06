##Usage
* With default ssh key ~/.ssh/id_rsa
```
remotecmd  -h ip_address1,ip_address2 -u user -c "df -h" #for example running df -h
```

* With custom private key
```
remotecmd  -h ip_address1,ip_address2 -u user -i /home/user1/ssh_key -c "df -h" #for example running df -h
```
* With password
```
remotecmd  -h ip_address1,ip_address2 -p -u user -c "df -h"  #for example running df -h
```