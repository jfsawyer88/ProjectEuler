#include <stdio.h>
#include <sys/time.h>


// input must be > 1
long pe003(long N)
{
  while (0 == N / 2) N /= 2;
  if (1 == N) return 2;

  long factor = 3;
  long last_factor;
  
  while (factor * factor < N)
    {
      if ( 0 == N % factor)
	{
	  last_factor = factor;
	  N /= factor;
	  while (0 == N % factor)
	    {
	      N /= factor;
	    }
	}
      factor += 2;
    }

  if (1 == N) return last_factor;
  else return N;
}




main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long N = 600851475143;
  //long N = 600851475149;
  //long N = 346815156100580087;
  //long N = 1030724545140160933;
  //long N = 412082356787843821;
  //long N = 9223372036854775783;
  long output = pe003(N);

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
	 (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
	 (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;
}
