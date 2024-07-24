package scripts

import (
	"fmt"
	"os"
	"os/exec"
)

// Create a user and database in the systems PostgreSQL instance.
// Username and password provided will be used to create the user.
// This function is an interface over the `create_user.sh` script.
// Any issues with the script will be returned as an error.
func CreateUser(username, password, postgres_password string) error {
	cmd := exec.Command("sudo", "./scripts/create_user.sh", username, password, postgres_password)

	cmd.Env = append(os.Environ(), "PGPASSWORD="+postgres_password)

	exitCode := 0
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			return err
		}
	}

	// View `/docs/scripts.md` for more information on the exit codes
	switch exitCode {
	case 1:
		return fmt.Errorf("Invalid arguments provided")
	case 2:
		return fmt.Errorf("Failed to create user")
	case 3:
		return fmt.Errorf("Failed to create database")
	case 4:
		return fmt.Errorf("Failed to grant privileges on database")
	case 5:
		return fmt.Errorf("Failed to grant privileges on schema")
	}
	return nil
}

// RunSQLCommand runs a SQL command using psql
func RunSQLCommand(sqlCommand string) error {
	cmd := exec.Command("psql", "-U", "postgres", "-c", sqlCommand)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", os.Getenv("PGPASSWORD")))

	// Capture and print the output for debugging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}
	return nil
}

func RunSQLCommandOnDB(db, sqlCommand string) error {
	cmd := exec.Command("psql", "-U", "postgres", "-d", db, "-c", sqlCommand)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", os.Getenv("PGPASSWORD")))

	// Capture and print the output for debugging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}
	return nil
}

func CreateUserNoScript(username, password, postgres_password string) error {
	os.Setenv("PGPASSWORD", postgres_password)

	// Create the new user/role
	if err := RunSQLCommand(fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s';", username, password)); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Create the database for the user
	if err := RunSQLCommand(fmt.Sprintf("CREATE DATABASE %s;", username)); err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	// Grant all privileges on database
	if err := RunSQLCommand(fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s;", username, username)); err != nil {
		return fmt.Errorf("failed to grant privileges on database: %v", err)
	}

	// Grant all privileges on schema
	if err := RunSQLCommandOnDB(username, fmt.Sprintf("GRANT ALL ON SCHEMA public TO %s;", username)); err != nil {
		return fmt.Errorf("failed to grant privileges on schema: %v", err)
	}

	os.Unsetenv("PGPASSWORD")
	return nil
}
