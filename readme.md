# CHUCK NORRIS CHALLENGE

Build
-----

Steps to build a Docker image:

1. Clone the project
```bash
  git clone https://github.com/NilsParedes/chuck-norris-challenge
```

2. Go to the project directory

```bash
  cd chuck-norris-challenge
```

3. Build the image

```bash
  docker build --rm -t chuck-norris-challenge .
```

This will take a few minutes.


4. Run the image's default command, which should start everything up. The `-p` option forwards the container's port 8081 to port 8081 on the host.

```bash
  docker run -p 8081:8081 chuck-norris-challenge
```

5. Once everything has started up, you should be able to access via [http://localhost:8081/](http://localhost:8000/) on your host machine.

```bash
  http://localhost:8081/
```

## API Reference

#### Get items

```http
  GET /api/items
```