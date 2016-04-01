#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, p, n;
  int N = 2000000;

  int *is_prime;
  is_prime = (int *)malloc(sizeof(int)*N);
  is_prime[0] = 0; is_prime[1] = 0;
  for (i = 2; i < N; i++) is_prime[i] = 1;

  long res = 0;
  //sieve out non-primes
  for (p = 2; p*p < N; p++)
    {
      if (is_prime[p])
	{
	  res += p;
	  for (n = p * p; n < N; n += p)
	    {
	      is_prime[n] = 0;
	    }
	}
    }

  //sum up primes
  //continuation on all i > sqrt(N)
  while (p < N)
    {
      if (is_prime[p])
	{
	  res += p;
	}
      p++;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", res);

  return 0;
}
