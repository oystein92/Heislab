#include <pthread.h>
#include <stdio.h>
// Note the return type: void*
static int i=0;
void* thread_1_function(){
	for(int j=0;j<=1000000;j++)
	{
		i++;
	}	
	return NULL;
}

void* thread_2_function(){
	for(int j=0;j<=1000000;j++)
	{
		i--;
	}	
	return NULL;
}

int main(){
pthread_t thread_1;
pthread_create(&thread_1, NULL, thread_1_function, NULL);
pthread_t thread_2;
pthread_create(&thread_2, NULL, thread_2_function, NULL);
// Arguments to a thread would be passed here ---------^
pthread_join(thread_1, NULL);
pthread_join(thread_2, NULL);
printf("%i",i);
return 0;
}
