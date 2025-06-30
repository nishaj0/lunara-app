# Database Migration Usage Guide

This guide explains how to use the migration system in your project using the `migrate` tool and the provided [Makefile](Makefile).

---

## Prerequisites

- Ensure you have [golang-migrate](https://github.com/golang-migrate/migrate) installed and available in your PATH.
- Ensure your `.env` file is configured with the correct database credentials.
- The `migrations` directory should exist in your project root.

---

## Important: Loading Environment Variables on Windows

The Makefile does **not** automatically load variables from your `.env` file into the shell environment on Windows.  
You must set these variables manually before running migration commands, or use a tool to load them.

### Option 1: Set Variables in PowerShell

```powershell
$env:POSTGRES_USER="your_user"
$env:POSTGRES_PASSWORD="your_password"
$env:POSTGRES_HOST="your_host"
$env:POSTGRES_PORT="your_port"
$env:DB_NAME="your_db"
```

### Option 2: Use dotenv-cli

If you have [Node.js](https://nodejs.org/) installed, you can use [dotenv-cli](https://www.npmjs.com/package/dotenv-cli):

```powershell
npx dotenv -e .env -- make migrate-up
```

---

## Common Commands

### 1. Create a New Migration

To create a new migration file, run:

```powershell
make migrate-create name=your_migration_name
```

Replace `your_migration_name` with a descriptive name, e.g., `create_users_table`.

This will generate two files in the `migrations` directory:

- `000001_create_users_table.up.sql` (for applying changes)
- `000001_create_users_table.down.sql` (for rolling back changes)

**Example:**

```
PS C:\MyProjects\lunara-client-next\lunara-server> make migrate-create name=create_users_table
```

---

### 2. Apply All Pending Migrations

To apply all new migrations to your database, run:

```powershell
make migrate-up
```

This will execute all `.up.sql` files that have not yet been applied.

---

### 3. Rollback the Last Migration

To undo the most recent migration, run:

```powershell
make migrate-down
```

This will execute the latest `.down.sql` file, reverting the last change.

---

### 4. Check Current Migration Version

To see the current migration version applied to your database, run:

```powershell
make migrate-version
```

---

### 5. Force Set Migration Version

If your migrations get out of sync, you can force the migration version:

```powershell
make migrate-force
```

You will be prompted to enter the version number to force.

---

## Example Workflow

1. **Create a migration:**
    - `make migrate-create name=add_phone_to_users`
    - Edit the generated `.up.sql` and `.down.sql` files as needed.

2. **Apply the migration:**
    - `make migrate-up`

3. **If you need to rollback:**
    - `make migrate-down`

---

## Notes

- Only edit migration files that have **not** been pushed or shared with others.  
  Once a migration is merged or used by others, always create a new migration for further changes.
- Always test both `up` and `down` migrations locally before pushing.

---