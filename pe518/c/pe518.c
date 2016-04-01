#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, j, n, p;
  int limit = 100000000;

  // prime sieve
  int *prime_bool;
  prime_bool = (int *)malloc(sizeof(int)*limit);
  prime_bool[0] = 0;  prime_bool[1] = 0;  prime_bool[2] = 1;
  for (i = 3; i < limit; i += 2) prime_bool[i] = 1;
  for (i = 4; i < limit; i += 2) prime_bool[i] = 0;
  for (p = 3; p*p < limit; p += 2)
    {
      if (prime_bool[p]) // if p is prime
	{
	  for (n = p*p; n < limit; n += p)
	    {
	      prime_bool[n] = 0;
	    }
	}
    }
  printf("check\n");
  int nprimes = 0;
  for (p = 0; p < limit; p++) if (prime_bool[p]) nprimes++;
  long *primes;
  primes = (long *)malloc(sizeof(long)*nprimes);
  printf("check\n");
  // populate primes array
  j = 0;
  for (i = 0; i < limit; i++)
    {
      if (prime_bool[i])
	{
	  primes[j] = i + 1;
	  j++;
	}
    }

  int low, high, mid, z;
  long output = 0;
  for (i = 0; i < nprimes; i++)
    {
      printf("p = %ld\n", primes[i]);
      //printf("started, ");
      for (j = i + 1; primes[j]*primes[j]/primes[i] < limit; j++)
	{
	  if (primes[j]*primes[j] % primes[i] == 0)
	    {
	      z = primes[j]*primes[j]/primes[i];
	      low = 0;
	      high = nprimes;
	      while (low <= high)
		{
		  mid = (high + low)/2;
		  if (z < primes[mid]) high = mid - 1;
		  else if (z == primes[mid])
		    {
		      //printf("(%d, %d, %d): ", primes[i]-1, primes[j]-1, z-1);
		      output += primes[i] + primes[j] + z - 3;
		      //printf("passed\n");
		      break;
		    }
		  else low = mid + 1;
		}
	    }
	}
    }






  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

}
