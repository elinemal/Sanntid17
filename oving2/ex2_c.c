#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

int i = 0;
pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;

void* add(){
	for (int a = 1; a < 1000000; a++) //can start on 0 and get 0 printed out as i in the end
	{
		pthread_mutex_lock(&lock);
		i++;
		pthread_mutex_unlock(&lock);
	}
	return NULL;
}

void* sub(){
	for (int a = 0; a < 1000000; ++a)
	{
		pthread_mutex_lock(&lock);
		i--;
		pthread_mutex_unlock(&lock);
	}
	return NULL;
}


int main(){
    pthread_t thread1;
    pthread_create(&thread1, NULL, add, NULL);
    // Arguments to a thread would be passed here in the end where we have NULL

    pthread_t thread2;
    pthread_create(&thread2, NULL, sub, NULL);
    
    pthread_join(thread2, NULL);
    pthread_join(thread1, NULL);
    pthread_mutex_destroy(&lock);
    printf("Done: %i\n",i);
    return 0;
}