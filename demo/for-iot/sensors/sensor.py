import random
import time
import threading
import logging
FILES = ["temperature", "humidity", "radar"]
FORMAT = "%(asctime)s: %(message)s"


def thread_function(name):
    data=random.randint(1,100)
    while True:
        
        logging.info(f"Updating {name}")
        with open(name,"w") as file:
            new_data=data+random.randint(-10,10)
            file.write(str(new_data))
        
        delay=random.randint(2,5)
        time.sleep(delay)
        data=new_data

if __name__ == "__main__":
    format = "%(asctime)s: %(message)s"
    logging.basicConfig(format=format, level=logging.INFO,
                        datefmt="%H:%M:%S")

    threads = list()
    for f in FILES:
        logging.info(f"Main    : create and start thread {f}")
        x = threading.Thread(target=thread_function, args=(f,))
        threads.append(x)
        x.start()

    for index, thread in enumerate(threads):
        logging.info(f"Main    : before joining thread {index}")
        thread.join()
        logging.info(f"Main    : thread {index} done")
