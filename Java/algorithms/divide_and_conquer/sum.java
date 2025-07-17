// Найти сумму массива
import java.util.Arrays;

public class divide_and_conque{

    private static int sum(int[] array){
        if (array.length == 0){
            return 0;
        }
        return array[0] + sum(Arrays.copyOfRange(array, 1, array.length)); // Аналогично срезу в питон [1:]
    }

    public static void main(String[] args){
        int[] array = {1, 3, 5, 8};
        System.out.println(sum(array));
    }
}
