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