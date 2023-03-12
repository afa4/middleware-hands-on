# MIDDLEWARE HANDS ON

This simple project is a composition of three docker containers:

- rabbitMQ broker
- publisher api
- consumer server

### lessons learned

for some reason, the rabbitmq container required
specific user grants (chown) that wasn't in the container entrypoint user.
So I had to configure UID and GUI variables so the container would
start whith current OS user permissions.

### build

```bash
chmod +x build.sh
chmod +x start.sh
./build.sh
./start.sh
```

### playing

After running de app with previous section instructions
run a POST request on the root resource / with the body below:

```json
{
  "Name": "Alice"
}
```

Check the logs of consumer container

```bash
docker-compose logs consumer -f
```

You should see a greeting with the requested name

````
consumer  | Hello Alice
```
