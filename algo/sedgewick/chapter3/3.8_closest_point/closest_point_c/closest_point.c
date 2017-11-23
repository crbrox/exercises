#include <math.h>
#include <stdlib.h>
#include <stdio.h>
#include <time.h>

typedef struct point {
	float x, y;
} point;

double distance(point a, point b)
{
	double dx = a.x - b.x, dy = a.y - b.y;
	return sqrt(dx * dx + dy * dy);
	//return hypot (dx, dy);
}

double randDouble()
{
	return 1.0 * random() / ((long)(1 << 31) - 1L);
}

int main(int argc, char *argv[])
{
	//double d = atof (argv[2]);
	double d = 0.5;
	int i, j, cnt = 0, N = 100000;	//= atoi (argv[1]);
	point *a;
	{
		clock_t t;
		t = clock();
		a = (point*)malloc(N * sizeof(*a));
		for (i = 0; i < N; i++) {
			a[i].x = randDouble();
			a[i].y = randDouble();
		}
		t = clock() - t;
		double time_taken = ((double)t) / CLOCKS_PER_SEC;	// in seconds
		printf("filling %f seconds \n", time_taken);
	}

	{
		clock_t t;
		t = clock();
		{
			for (i = 0; i < N; i++) {
				for (j = 0; j < N; j++) {
					if (distance(a[i], a[j]) < d)
						cnt++;
				}
			}
		}
		t = clock() - t;
		double time_taken = ((double)t) / CLOCKS_PER_SEC;	// in seconds
		printf("searching %f seconds \n", time_taken);
	}

	printf("%d edges shorter than %f\n", cnt, d);
}
