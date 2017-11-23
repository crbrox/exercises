#include <math.h>
#include <stdlib.h>
#include <stdio.h>


double
distance (double a, double b)
{
  return sqrt (a * a + b * b);
  //return hypot(dx,dy);
}

double
randDouble ()
{
  return 1.0 * random () / ((long) (1 << 31) - 1L);
}

int
main (int argc, char *argv[])
{
  double cnt = 0.0;
  for (long j = 0; j < 100000L * 100000L; j++)
    {
      cnt += distance ((double) j, (double) (j + 1));

    }
  printf ("%f\n", cnt);
}
