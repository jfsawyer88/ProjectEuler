#include <stdio.h>
#include <sys/time.h>

// highest power of p that divides n
int npk(int n, int p)
{
  int out = 1;
  while (0 == n % p)
    {
      n /= p;
      out *= p;
    }
  return out;
}

// largest s sucht that s^2 < n
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


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, j, n, p, pk;
  long a;
  int N = 1 + 10000000;
  int sqrtN = isqrt(N);

  // prime sieve
  int prime_bool[sqrtN];
  prime_bool[0] = 0;
  prime_bool[1] = 0;
  for (i = 2; i < sqrtN; i++)
    {
      prime_bool[i] = 1; // set all but first
    }                    // positions to 1

  int lpf0[sqrtN];
  for (i = 0; i < sqrtN; i++)
    {
      lpf0[i] = 1;
    }

  for (p = 2; p < sqrtN; p++)
    {
      if (1 == lpf0[p])
	{
	  lpf0[p] = p;
	  for (n = 2 * p; n < sqrtN; n += p)
	    {
	      prime_bool[n] = 0;
	      pk = npk(n, p);
	      if (lpf0[n] < pk)
		{
		  lpf0[n] = pk;
		}
	    }
	}
    }

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

  // sum of M(n) for all n < sqrt(N)
  long output = 0;
  for (n = 2; n < sqrtN; n++)
    {
      // if n is a prime power then it
      // only contributes 1 to the sum
      if (n == lpf0[n])
	{
	  output++;
	}
      else
	{
	  for (a = n - lpf0[n]; a > n/2 - 1; a -= lpf0[n])
	    {
	      if ((a+1) == ((a+1) * (a+1)) % n)
		{
		  output += a + 1;
		  break;
		}
	      else if ( a == (a * a) % n)
		{
		  output += a;
		  break;
		}
	    }
	}
    }

  int segment_size = 1 << 15;
  int low = sqrtN;
  int next_indx[n_primes];
  int lpf[segment_size];
  int T[segment_size]; // to divide fully by primes

  // initialize next_indx array
  for (i = 0; i < n_primes; i++)
    {
      next_indx[i] = (primes[i] - (low % primes[i])) % primes[i];
    }

  // segment loop
  while (low < N)
    {      
      // initialize lpf and T
      for (i = 0; i < segment_size; i++)
	{
	  lpf[i] = 1;
	  T[i] = low + i;
	}

      // find highest prime power for each prime
      // next_indx tells us where ot start for each prime
      for (i = 0; i < n_primes; i++)
	{
	  while (next_indx[i] < segment_size)
	    {
	      j = next_indx[i]; // current integer
	      n = low + j;
	      p = primes[i];    // current prime

	      pk = npk(n, p);
	      if (lpf[j] < pk)
		{
		  lpf[j] = pk;
		}

	      T[j] /= pk;
	      next_indx[i] += p;
	    }
	}
	

      // go through current segment
      i = 0;
      n = low;
      while ( (i < segment_size) && (n < N) )
	{

	  // fix lpf when largest p^k
	  // dividing n has p > sqrtN
	  if (lpf[i] < T[i])
	    {
	      lpf[i] = T[i]; 
	    }

	  // do the M(n) algorithm like usual
	  if (n == lpf[i])
	    {
	      output++;
	    }
	  else
	    {
	      for (a = n - lpf[i]; a > n/2 - 1; a -= lpf[i])
		{
		  if ((a+1) == ((a+1) * (a+1)) % n)
		    {
		      output += a + 1;
		      break;
		    }
		  else if ( a == (a * a) % n)
		    {
		      output += a;
		      break;
		    }
		}
	    }
	  i++;
	  n++;
	}

      // re-initiaize nex_indx for next segment
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
