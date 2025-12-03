import redis  # type: ignore

r = redis.Redis(host="redis", port=6379, decode_response=True)

r.set("message", "Hello from Redis in Docker!")
print(r.get("message"))
