import requests

r = requests.post("http://127.0.0.1:8080/good/create?projectId=1", json={"name": "string 6"})

print(r.text)