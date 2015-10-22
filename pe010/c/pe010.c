#include <stdio.h>
#include <sys/time.h>

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int N = 2000000;

  int prime_bool[N];
  int i, j;

  for (i = 2; i < N; i++)
    {
      prime_bool[i] = 1;  //set all but first
    }                     //positions to 1

  long res = 0;
  i = 2;
  //sieve out non-primes
  while (i * i < N)
    {
      if (prime_bool[i])
	{
	  res += i;
	  j = i * i;
	  while (j < N)
	    {
	      prime_bool[j] = 0;
	      j += i;
	    }
	}
      i++;
    }

  //sum up primes
  //long res = 0;
  //  for (i = 2; i < N; i++)
  while (i < N)
    {
      if (prime_bool[i])
	{
	  res += i;
	}
      i++;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", res);

  return 0;
}
