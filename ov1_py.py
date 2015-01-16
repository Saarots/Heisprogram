from threading import Thread


i = 0

global i	

def main():
	Thread1 = Thread(target = Thread1func, args = (),)
	Thread2 = Thread(target = Thread2func, args = (),)

	Thread1.start()
	Thread2.start()

	Thread1.join()
	Thread2.join()
	print i

def Thread1func():
	for i in range(0,1000000):
		i += 1
def Thread2func():
	for i in range(0,1000000):
		i -= 1

main()
