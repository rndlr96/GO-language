
import java.util.Scanner;

public class ChainedMatrix2 {
	   public static void main(String args[])
	   {
	      int start;
	      int minimum;
	      System.out.println("행렬의 개수를 입력하세요");
	      Scanner scanner = new Scanner(System.in);

	      int n = scanner.nextInt();
	      int d[] = new int[n+1];

	      for(int i=0; i<=n; i++) {

	         System.out.println("d["+i+"]를 입력하세요.");
	         Scanner scanner1 = new Scanner(System.in);
	         d[i] = scanner1.nextInt();
	      }

	      minimum = Matrix1(1,n,d);
	      System.out.print("행렬의 최소 곱셈 횟수는:"+minimum);

	   }


	   static int Matrix1(int start,int end, int d[]) {

	      int k;
	      int minimum = 10000000;
	      if(start == end) {
	         return 0;
	      }

	      for(k = start; k < end; k++) {
	         int count = (Matrix1(start,k,d)+ Matrix1(k+1,end,d)+d[start-1]*d[k]*d[end]);
	         if(minimum > count) {
	            minimum = count;
	         }

	      }

	      return minimum;

	   }

}
