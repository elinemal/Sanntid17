# oving2

from threading import Thread
import threading

i = 0
lock = threading.Lock()
threads = []

def add(): #thread1
    global i
    for a in range(0,100000):
      lock.acquire()
      i += 1
      lock.release()
  
def sub(): #thread2
  global i
  for x in range(0,1000000):
      lock.acquire()
      i-=1
      lock.release()
  
def main():
    Thread1 = Thread(target = add, args = (),) 
    Thread1.start()
    Thread2 = Thread(target = sub, args = (),) 
    Thread2.start()

    threads.append(Thread1)
    threads.append(Thread2)

    for t in threads:
      t.join()
   
    print(i)

main()