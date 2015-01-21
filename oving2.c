#include <pthread.h>
#include <stdio.h>

static int i = 0;
pthread_mutex_t mtx;

void* Thread1func(){
	for(int j = 0; j < 999999; j++){

		pthread_mutex_lock(&mtx);
		i++;
		pthread_mutex_unlock(&mtx);
	}
	return NULL;
}
void* Thread2func(){
	for(int j = 0; j < 1000000; j++){
		pthread_mutex_lock(&mtx);
		i--;
		pthread_mutex_unlock(&mtx);
	}
	return NULL;
}
int main(){
	pthread_mutex_init(&mtx,NULL);	

	pthread_t Thread1,Thread2;

	pthread_create(&Thread1, NULL, Thread1func, NULL);
	pthread_create(&Thread2, NULL, Thread2func, NULL);

	pthread_join(Thread1, NULL);
	pthread_join(Thread2, NULL);

	pthread_mutex_destroy(&mtx);

	printf ("i = %d\n",i);
	return 0;
}
