#include <stdio.h>
#include <sys/time.h>


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long i;
  long sqrsum = 0;
  long sumsqr = 0;
  for (i = 1; i <= 100; i++)
    {
      sumsqr += i * i;
      sqrsum += i;
    }

  long res = sqrsum * sqrsum - sumsqr;

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", res);

  return 0;
}
