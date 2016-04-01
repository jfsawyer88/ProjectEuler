#include <stdio.h>
#include <sys/time.h>


int mult_order(int p)
{
  int k = 1;
  int mul = 10 % p;
  while (1 != mul)
    {
      k++;
      mul = (mul * 10) % p;
    }
  return k;
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int max = 1000000;
  int i, p, n, cl;

  int primes[max]; // prime bool array

  //initialize it!
  primes[0] = 0;
  primes[1] = 0;
  for (i = 2; i < max; i++)
    {
      primes[i] = 1;
    }

  // sieve it!
  for (p = 2; p*p < max; p++)
    {

      // if p is prime
      if (1 == primes[p])
	{
	  for (n = p * p; n < max; n += p)
	    {
	      primes[n] = 0;
	    }
	}
    }

  for (n = max - 1; n >= 0; n--)
    {
      if (1 == primes[n])
	{
	  cl = mult_order(n);
	  if (n == cl + 1)
	    {
	      break;
	    }
	}
    }

  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", n);

  return 0;
}
