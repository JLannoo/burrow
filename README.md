# Burrow
Burrow is a Go CLI tool for managing your passwords.

It uses the filesystem to store your passwords in an organized way.  
You only need to remember one password to unlock all your other passwords.

> [!CAUTION]
> THIS IS PROBABLY NOT ACTUALLY SECURE.  
> I AM NOT A SECURITY EXPERT AND THIS IS JUST A FUN PROJECT.  
> USE AT YOUR OWN RISK.

## Features
- Store passwords encrypted on your filesystem
- Generate secure passwords

## Installation
```bash
go get github.com/JLannoo/burrow
```

## Usage
```bash
burrow [command]
```

## Commands
- `init` - Initialize the password store. This will create a new directory in your home directory called `.burrow`
- `add` - Add a new password
- `generate` - Generate a new password
- `list` or `ls` - List all the passwords in the store
- `get` - Get a password
- `remove` or `rm`  - Remove a password
- `update` - Update a password
