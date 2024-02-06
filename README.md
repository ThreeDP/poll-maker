# poll-maker
This project is focused on developing a poll API.

## Config project

### Create .env
> Create a `.env` file and insert the following environment variables:

```env
    POSTGRES_DB=<db_name>
    POSTGRES_USER=<db_username>
    POSTGRES_PASSWORD=<db_password>
    ALLOW_EMPTY_PASSWORD=<allow_empty_pass_on_redis[true / false]>
```

### Up Containers
> Here are the commands for managing containers:

**Build and Start Containers**
```sh
make
```

**Build Containers**
```
make build
```

**Start Containers**
```sh
make up
```

**Stop Containers**
```sh
make down
```

**Cleaning Images and Volumes**
```sh
make clean
```

**Terminating Containers Process**
````sh
make fclean
```

Feel free to use these commands to set up and manage the project environment easily.