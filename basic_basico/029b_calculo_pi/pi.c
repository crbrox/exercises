#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char *argv[]){

	if(argc!=2) {
		puts("Uso: pi <num_lanzamientos>");
		exit(-1);
	}

	long n = strtol(argv[1],NULL,10); // número lanzamientos
	long m = 0;
	srand((unsigned)time(NULL));
	for (long i=0; i < n; i++) {
		double X=((double)rand()/(double)RAND_MAX);
		double Y=((double)rand()/(double)RAND_MAX);
		if (X*X + Y*Y < 1) {
			m++;
		}
	}
  	printf("Si lanzas %ld dardos obtenemos\n", n);
	printf("un valor aproximado de Pi = %f\n", 4.0 * (double)m/(double)n);
	return 0;
}
