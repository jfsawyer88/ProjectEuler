#include <stdio.h>
#include <sys/time.h>

int isqrt(int n)
{
  int xk0 = n;
  int xk1 = (xk0 + n / xk0) / 2;
  while ( ((xk1 - xk0) < -1) || (1 < (xk1-xk0)) )
    {
      xk0 = xk1;
      xk1 = (xk0 + n / xk0) / 2;
    }
  return xk1;
}

int full_div(int n, int p)
{
  while (0 == n % p)
    {
      n /= p;
    }
  return n;
}

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  //int N = 1 + 1000000;
  int N = 1 + 500000000;
  int sqrtN = isqrt(N);
  int i, j;

  // prime sieve
  int prime_bool[sqrtN];
  prime_bool[0] = 0;
  prime_bool[1] = 0;
  for (i = 2; i < sqrtN; i++)
    {
      prime_bool[i] = 1;  //set all but first
    }                      //positions to 1

  // initialize array for totient values
  int phi0[sqrtN];
  for (i = 0; i < sqrtN; i++)
    {
      phi0[i] = i;
    }
  for (i = 2; i < sqrtN; i++)
    {
      // if i is prime
      if (i == phi0[i])
	{
	  phi0[i] -= 1;
	  for (j = 2 * i; j < sqrtN; j += i)
	    {
	      phi0[j] = (phi0[j] / i) * (i - 1); // adjust the phi value
	      prime_bool[j] = 0;                 // sieve out the prime
	    }
	}
    }

  // number of primes
  // to make primes array
  int n_primes = 0;
  for (i = 0; i < sqrtN; i++)
    {
      n_primes += prime_bool[i];
    }

  // array of all primes below sqrt(N)
  int primes[n_primes];
  j = 0;
  for (i = 0; i < sqrtN; i++)
    {
      if (prime_bool[i])
	{
	  primes[j] = i;
	  j++;
	}
    }
  
  // the sum of phi values below sqrtN
  // we will add the rest in a segmented sieve
  long output = 0;
  for (i = 1; i < sqrtN; i += 2)
    {
      output += phi0[i];
    }

  int segment_size = 1 << 15;
  int low = sqrtN;
  int next_indx[n_primes];
  int phi[segment_size];
  int T[segment_size]; // to divide fully by primes
  for (i = 0; i < n_primes; i++)
    {
      next_indx[i] = (primes[i] - (low % primes[i])) % primes[i];
    }


  while (low < N)
    {
      //initialize current phi segment
      for (i = 0; i < segment_size; i++)
	{
	  phi[i] = low + i;
	  T[i]   = low + i;
	}

      // adjust each multiple of each prime
      // next_indx tells us where to start for each prime
      for (i = 0; i < n_primes; i++)
	{
	  while (next_indx[i] < segment_size)
	    {
	      phi[next_indx[i]] = (phi[next_indx[i]] / primes[i]) * (primes[i] - 1);
	      T[next_indx[i]]   = full_div(T[next_indx[i]], primes[i]);
	      next_indx[i] += primes[i];
	    }
	}

      // fix phi values (those j's with prime factors > sqrt(i)
      // and add the odd values to the output
      i = 0;
      j = low;
      while ( (i < segment_size) && (j < N))
	{
	  // phi must be fixed if we failed to divide j down to 1
	  // NB: the remaining value will be the last prime
	  if ( T[i] != 1 )
	    {
	      phi[i] = (phi[i] / T[i]) * (T[i] - 1);
	    }

	  // add phi to sum if j is odd
	  if ( 1 == j % 2 )
	    {
	      output += phi[i];
	    }

	  //don't forget to increment
	  i += 1;
	  j += 1;
	}

      // re-initialize next_indx for next segment
      for (i = 0; i < n_primes; i ++)
	{
	  next_indx[i] -= segment_size;
	}
      low += segment_size;
    }


  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;
}
