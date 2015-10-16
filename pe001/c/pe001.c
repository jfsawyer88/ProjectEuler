#include <stdio.h>
#include <sys/time.h>

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i;
  int res = 0;

  for (i = 1; i < 1000; i++)
    {
      if ((0 == i % 3) || (0 == i % 5))
	{
	  res += i;
	}
    }

  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", res);

  return 0;
}
