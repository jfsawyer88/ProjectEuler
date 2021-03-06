#include <stdio.h>
#include <sys/time.h>

int isqrt(int n)
{
  int xk = n;
  int xk1 = (xk + (n / xk)) / 2;
  while (xk1 < xk)
    {
      xk = xk1;
      xk1 = (xk + (n /xk)) / 2;
    }
  return xk;
}

//checks if element e is in array arr
//uses binary search (array must be sorted)
int in_array(long e, int arr[], int len)
{
  int low  = 0;
  int high = len - 1;
  int mid;
  while (low <= high)
    {
      mid = (low + high) / 2;
      if (arr[mid] > e)
	{
	  high = mid - 1;
	}
      if (arr[mid] < e)
	{
	  low = mid + 1;
	}
      if (arr[mid] == e)
	{
	  return 1;
	}
    }
  return 0;
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int N = 1000000;

  int prime_bool[1000000] = {0};
  int i, j;

  prime_bool[2] = 1;
  //initialize to zero
  for (i = 3; i < N; i += 2)
    {
      prime_bool[i] = 1;  //set all but first
    }                     //positions to 1

  //sieve out non-primes
  int res = 0;
  for (i = 3; i <= isqrt(N); i += 2)
    {
      if (prime_bool[i])
	{
	  j = i * i;
	  while (j < N)
	    {
	      prime_bool[j] = 0;
	      j += i;
	    }
	}
    }

  //sum up primes
  int num_primes = 0;
  for (i=0; i < N; i++)
    {
      if (prime_bool[i])
	{
	  num_primes += 1;
	}
    }

  int primes[num_primes];
  long prime_accum[num_primes + 1];
  prime_accum[0] = 0;
  j = 0;
  for (i = 0; i < N; i++)
    {
      if (prime_bool[i])
	{
	  primes[j] = i;
	  prime_accum[j + 1] = i + prime_accum[j];
	  j++;
	}
    }

  int k;
  long ans;
  for (k = 0; k < num_primes; k++)
    {
      for (i = 0; i <= k; i++)
	{
	  j = (num_primes - k) + i;
	  if (N < (prime_accum[j] - prime_accum[i]))
	    {
	      break;
	    }

	  if (in_array(prime_accum[j] - prime_accum[i], primes, num_primes))
	    {
	      ans = prime_accum[j] - prime_accum[i];
	      k = num_primes;
	      break;
	    }
	}
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
	 (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
	 (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", ans);

  return 0;
}
