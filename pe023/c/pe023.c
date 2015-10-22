#include <stdio.h>
#include <sys/time.h>


main ()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, j;
  int sod[28123];
  // initialize sod array
  for (i = 0; i < 28123; i++)
    {
      sod[i] = 0;
    }
  // fill sod array
  for (i = 1; i < 28123; i++)
    {
      for (j = 2 * i; j < 28123; j+= i)
	{
	 sod[j] += i;
	}
    }
  // number of abundant numbers
  // to help initialize abundants array
  int num_abundants;
  for (i = 0; i < 28123; i++)
    {
      if (sod[i] > i)
	{
	  num_abundants++;
	}
    }
  int abundants[num_abundants];
  j = 0;
  for (i = 0; i < 28123; i++)
    {
      if (sod[i] > i)
	{
	  abundants[j] = i;
	  j++;
	}
    }
  // initialize with all true
  int not_sum_of_abundants[28123];
  int sum;
  for (i = 0; i < 28123; i++)
    {
      not_sum_of_abundants[i] = 1;
    }

  for (i = 0; i < num_abundants; i++)
    {
      for (j = i; j < num_abundants; j++)
	{
	  sum = abundants[i] + abundants[j];
	  if (sum > 28123)
	    {
	      break;
	    }
	  not_sum_of_abundants[sum] = 0;
	}
    }

  int res = 0;
  for (i = 0; i < 28123; i++)
    {
      if (not_sum_of_abundants[i])
	{
	  res += i;
	}
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", res);

}
