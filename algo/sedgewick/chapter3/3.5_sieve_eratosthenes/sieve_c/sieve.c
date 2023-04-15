#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define N 10000000

//bool a[N];

int main(void) {
  bool *a = malloc(sizeof(bool) * N);
  // for (long long i = 2; i < N; i++) { a[i] = true;}
  memset(a, true, sizeof(bool) * N);
  for (long long i = 2; i < N; i++)
    if (a[i])
      for (long long j = i; i * j < N; j++)
        a[i * j] = false;
  long long count = 0;
  for (long long i = 2; i < N; i++)
    if (a[i])
      count++;

  printf("%lld primes\n", count);
}