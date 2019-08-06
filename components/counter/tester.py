import requests
import random
from time import sleep
for i in range(0,100):
    r=requests.get("http://localhost:8080/count",params={"payload":"atakan"})
    sleep(random.randint(1,3)*0.1)
    print(i)
