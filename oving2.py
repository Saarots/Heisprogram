from threading import Thread, Lock

i = 0


mtx = Lock()

def main():
	
	
	Thread1 = Thread(target = Thread1func, args = (),)
	Thread2 = Thread(target = Thread2func, args = (),)
	
	Thread1.start()
	Thread2.start()
	
	Thread1.join()
	Thread2.join()
	
	
	print i

def Thread1func():
	global i
	for j in range(0,999999):
		mtx.acquire()
		i += 1
		mtx.release()
def Thread2func():
	global i	
	for j in range(0,1000000):
		mtx.acquire()		
		i -= 1
		mtx.release()
main()
