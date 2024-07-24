#!/usr/bin/env bash

# Ensure details were provided
if [ $# -ne 3 ]; then
  echo "Usage: $0 <username> <password> <postgres_password>"
  exit 1
fi

USERNAME=$1
PASSWORD=$2
export PGPASSWORD=$3

# Create the new user/role
psql -U postgres -c "CREATE USER $USERNAME WITH PASSWORD '$PASSWORD';"

# Check if the user was created successfully
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to create user $USERNAME"
  exit 2
fi

# Create the database for the user
psql -U postgres -c "CREATE DATABASE $USERNAME;"

# Check if the database was created successfully
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to create database $USERNAME"
  exit 3
fi

# Grant all privileges on database
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $USERNAME TO $USERNAME;"

# Check if the privileges were granted successfully
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to grant privileges on database $USERNAME"
  exit 4
fi

# Grant all privileges on schema
psql -U postgres -d $USERNAME -c "GRANT ALL ON SCHEMA public TO $USERNAME;"

# Check if the privileges were granted successfully
if [ $? -ne 0 ]; then
  unset PGPASSWORD
  echo "Failed to grant privileges on schema $USERNAME"
  exit 5
fi

unset PGPASSWORD
