# MIDDLEWARE HANDS ON

This simple project is a composition of three docker containers:

- rabbitMQ broker
- api
- consumer server

Requests are receive by API and content is published to rabbitMq broker queue and returs response.
Consumer receives de message a prints the result.

## Lessons Learned

Simple archtecture for asynchronous processing

For some reason, the rabbitmq container required
specific user grants (chown) that wasn't in the container entrypoint user.
So I had to configure UID and GUI variables so the container would
start whith current OS user permissions.

## Building

```bash
chmod +x build.sh
chmod +x start.sh
./build.sh
./start.sh
```

## Playing

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
