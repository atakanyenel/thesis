import subprocess
import random
from time import sleep
for i in range(0,100):
    subprocess.run(["./fp"])
    sleep(random.randint(2,15)*0.1)
    print(i)
