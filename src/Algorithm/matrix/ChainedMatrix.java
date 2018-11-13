
import java.util.Scanner;

public class ChainedMatrix {
   static int operation =0;

   public static void main(String args[])
   {

      System.out.println("행렬의 개수를 입력하세요");
      Scanner scanner = new Scanner(System.in);

      int n = scanner.nextInt();
      int d[] = new int[n+1];
      int M[][] = new int[n][n];
      for(int i=0; i<=n; i++) {

         System.out.println("d["+i+"]를 입력하세요.");
         Scanner scanner1 = new Scanner(System.in);
         d[i] = scanner1.nextInt();

      }

      for(int i=0 ; i < n ; i++) {
    	  M[i][i] = 0;
      }

      for(int i = 1; i < n; i++) {
         for(int j = 0; j < n-i; j++) {
        	 int limit = i + j;
             int minimum = 10000000;

        	 for(int k = j ; k < limit ; k++) {
               int count = (M[j][k] + M[k+1][limit] + d[j]*d[k+1]*d[limit+1]);
               operation++;
               if(minimum > count)
               {
                  minimum = count;
                  M[j][limit] = minimum;
               }
            }
         }
      }

      System.out.println("행렬의 최소 곱셈 횟수는 : "+M[0][n-1]);
      System.out.print("단위연산의 횟수는:"+operation);

   }

}
