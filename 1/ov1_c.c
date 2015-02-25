#include <pthread.h>
#include <stdio.h>

static int i = 0;


void* Thread1func(){
	for(int j = 0; j < 1000000; j++){
		i++;
	}
	return NULL;
}

void* Thread2func(){
	for(int j = 0; j < 1000000; j++){
		i--;
	}
	return NULL;
}


int main(){
	pthread_t Thread1,Thread2;

	pthread_create(&Thread1, NULL, Thread1func, NULL);
	pthread_create(&Thread2, NULL, Thread2func, NULL);

	pthread_join(Thread1, NULL);
	pthread_join(Thread2, NULL);

	printf ("i = %d\n",i);
	return 0;
}




