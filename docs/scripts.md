# Script Documentation: User and Database Management

## Table of Contents
- #### [Create User](#Create-User)
- #### [Delete User](#Delete-User)

## Create User
`create_user.sh`

### Description
This script creates a new PostgreSQL user with the specified username and password, 
creates a database for the user, and grants all necessary privileges. The `postgres`
user is required to run this script as well as `sudo` privileges. The password for 
the `postgres` user is required to run this script and should be provided in the 
3rd argument.

### Usage
```bash
./scripts/create_user.sh <username> <password> <postgres_password>
```

### Exit Codes

| Exit Codes | Description                            |
| ---------- | -------------------------------------- | 
| 0          | Script exectued successfully           |
| 1          | Invalid arguments provided             |
| 2          | Failed to create user                  |
| 3          | Failed to create database              |
| 4          | Failed to grant privileges on database |
| 5          | Failed to grant privileges on schema   |


## Delete User
`delete_user.sh`

### Description
This script deletes a PostgreSQL user with the specified username, deletes the database
associated with the user, and revokes their permissions on the schema. The `postgres`
user is required to run this script as well as `sudo` privileges. The password for 
the `postgres` user is required to run this script and should be provided in the
2nd argument.

### Usage
```bash
./scripts/delete_user.sh <username> <postgres_password>
```

### Exit Codes

| Exit Codes | Description                            |
| ---------- | -------------------------------------- | 
| 0          | Script exectued successfully           |
| 1          | Invalid arguments provided             |
| 2          | Failed to delete user                  |
| 3          | Failed to delete database              |
| 5          | Failed to revoke privileges on schema  |
