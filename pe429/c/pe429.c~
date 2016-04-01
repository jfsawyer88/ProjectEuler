#include <stdio.h>
#include <sys/time.h>

// modulo exponentiation
long modpow(long a, long b, long m)
{
  if (b == 1) return a;
  if (0 == b % 2)
    {
      return modpow(a * a % m, b/2, m);
    }
  else
    {
      return (a * modpow(a * a % m, b/2, m) % m);
    }
}

// number of factors of p in N!
long nfp(long p, long N)
{
  long res = 0;
  while (N > 0)
    {
      N /= p;
      res += N;
    }
  return res;
}

long isqrt(long n)
{
  long xk0 = n;
  long xk1 = (xk0 + n / xk0) / 2;
  while ( ((xk1 - xk0) < -1) || (1 < (xk1-xk0)) )
    {
      xk0 = xk1;
      xk1 = (xk0 + n / xk0) / 2;
    }
  return xk1;
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long N = 100000000;
  long M = 1000000009;
  long sqrtN = isqrt(N);
  long i, j, k, n, p;

  int prime_bool[sqrtN];
  prime_bool[0] = 0;
  prime_bool[1] = 0;
  for (i = 2; i < sqrtN; i++)
    {
      prime_bool[i] = 1;  //set all but first
    }                     //positions to 1

  // sieve it
  for (i = 2; i < sqrtN; i++)
    {
      // if i is prime
      if (prime_bool[i])
	{
	  for (j = i*i; j < sqrtN; j += i)
	    {
	      //sieve multiplies of the prime
	      prime_bool[j] = 0;
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

  // the sum for primes below sqrtN
  // will add the rest in the segmented sieve
  long output = 1;
  for (i = 0; i < n_primes; i++)
    {
      p = primes[i];
      k = nfp(p, N);
      output = ((1 + modpow(p, 2 * k, M)) * output) % M;
    }


  long segment_size = 1 << 15;
  long low = sqrtN;
  long next_indx[n_primes];
  long is_prime_seg[segment_size];
  // initialize next_indx array
  for (i = 0; i < n_primes; i ++)
    {
      p = primes[i];
      next_indx[i] = (p - (low % p)) % p;
    }
  // the segmented sieve
  while (low < N)
    {
      // initialize current segment
      for (i = 0; i < segment_size; i++)
	{
	  is_prime_seg[i] = 1;
	}

      // cross out each multiple of a prime
      for (i = 0; i < n_primes; i++)
	{
	  while (next_indx[i] < segment_size)
	    {
	      is_prime_seg[next_indx[i]] = 0;
	      next_indx[i] += primes[i];
	    }
	}

      // do the thing on each prime in the segment
      i = 0;
      j = low;
      while ( (i < segment_size) && (j < N) )
	{
	  if (is_prime_seg[i])
	    {
	      p = low + i;
	      k = nfp(p, N);
	      output = ((1 + modpow(p, 2 * k, M)) * output) % M;
	    }
	  i ++;
	  j ++;
	}

      // re-initialize next_indx for next segment
      for (i = 0; i < n_primes; i++)
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
