#!/usr/bin/env bash

# Check if the user exists
if [ -z "$2" ]; then
    echo "Usage: $0 <username> <postgres_password>"
    exit 1
fi

USERNAME=$1
export PGPASSWORD=$2

# Delete owned tables
# This might not be required
#psql -U postgres -d $USERNAME -c "DROP OWNED BY $USERNAME;"

# Revoke privileges
psql -U postgres -d $USERNAME -c "REVOKE ALL ON SCHEMA public FROM $USERNAME;"

# Ensure the privileges were revoked
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to revoke privileges from $USERNAME"
  exit 5
fi

# Delete the database
psql -U postgres -c "DROP DATABASE $USERNAME;"

# Ensure the database was deleted
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to delete database $USERNAME"
  exit 3
fi

# Delete the user
psql -U postgres -c "DROP ROLE $USERNAME;"

# Ensure the user was deleted
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to delete role $USERNAME"
  exit 2
fi

# unset PGPASSWORD
