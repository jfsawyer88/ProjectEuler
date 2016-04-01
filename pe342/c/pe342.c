#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>


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

int is_cube(long n)
{
  long l = 1; // low
  long h = 1; // high
  for (h = 1; h*h*h <= n; h *= 2);
  while (1 + l != h)
    {
      //printf("l = %ld, h = %ld, n = %ld\n", l, h, n);
      long m = (l+h)/2; //mid
      if (m*m*m < n) l = m;
      else if (m*m*m == n) return 1;
      else h = m;
    }
  return 0;
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

  long N = 10000000000;
  long sqrtN = isqrt(N);
  int j, n, p;
  long i;

  /*  
  for (n = 1; n <= 125; n++)
    {
      if (is_cube(n)) printf("%d is a cube\n", n);
      else printf("%d is not a cube\n", n);
    }
  */

  //is_cube(2148183452);

  // prime sieve
  int prime_bool[sqrtN];
  prime_bool[0] = 0;
  prime_bool[1] = 0;
  for (i = 2; i < sqrtN; i++)
    {
      prime_bool[i] = 1;  //set all but first
    }                     //positions to 1

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
	      // or phi0[j] = phi0[j] - phi[0]/i;
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
  long output = 1;
  for (i = 2; i < sqrtN; i++)
    {
      if (is_cube(i * phi0[i]))
	{
	  //printf("phi[%d] = %d\n", i, phi0[i]);
	  output += i;
	}
    }

  int segment_size = 1 << 15;
  int low = sqrtN;
  int next_indx[n_primes];
  long phi[segment_size];
  int T[segment_size]; // to divide fully by primes

  for (i = 0; i < n_primes; i++)
    {
      next_indx[i] = (primes[i] - (low % primes[i])) % primes[i];
    }

  double progress;
  while (low < N)
    {
      progress = 100 * (double) low / N;
      gettimeofday(&tv2, NULL);
      printf("progress: %f percent, in %f seconds\n", progress,
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));

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
	      // or phi[next_indx[i]] = phi[next_indx[i]] - phi[next_indx[i]]/primes[i]
	      n = next_indx[i];
	      p = primes[i];

	      phi[n] = (phi[n] / p) * (p - 1);
	      T[n] = full_div(T[n], p);
	      next_indx[i] += p;

	      //phi[next_indx[i]] = (phi[next_indx[i]] / primes[i]) * (primes[i] - 1);
	      //T[next_indx[i]]   = full_div(T[next_indx[i]], primes[i]);
	      //next_indx[i] += primes[i];
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

	  // add phi to the output
	  // (for all, not just evens)
	  if (is_cube((i + low) * phi[i]))
	    {
	      //printf("phi[%d] = %d\n", i + low, phi[i]);
	      output += (i + low);
	    }

	  //don't forget to increment
	  i += 1;
	  j += 1;
	}

      // re-initialize next_indx for next segment
      for (i = 0; i < n_primes; i++)
	{
	  next_indx[i] -= segment_size;
	}
      low += segment_size;
    }


  // use sum of phis to get answer
  //output = (3 * ((long) N) * ((long) N - 1)) - (6 * output);

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;
}
