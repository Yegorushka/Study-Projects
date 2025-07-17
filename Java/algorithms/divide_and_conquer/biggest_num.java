// Найти самое большое число в массиве
import java.util.Arrays;

public class biggest_num{

    private static int biggestNum(int[] array){
        if (array.length == 2){
            if (array[0] > array[1]) {
                return array[0]; } 
            return array[1];
        }

        int sub_max = biggestNum(Arrays.copyOfRange(array, 1, array.length));
        if (array[0] > sub_max){
            return array[0]; }
        return sub_max;
    }

    public static void main(String[] args){
        int[] array = {1, 3, 5, 8, 9};
        System.out.println(biggestNum(array));
    }
}
