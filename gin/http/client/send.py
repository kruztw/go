import requests

res = requests.get("http://localhost:8888?title=a&content=b", headers={"username":"user1", "password":"pwd"})
print(res.text)
