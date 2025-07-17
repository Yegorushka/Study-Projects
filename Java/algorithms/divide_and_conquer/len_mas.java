// Найти длину массива
import java.util.Arrays;

public class len_mas{

    private static int len(int[] array){
        if (array.length == 0){
            return 0;
        }
        return 1 + len(Arrays.copyOfRange(array, 1, array.length)); // Аналогично срезу в питон [1:]
    }

    public static void main(String[] args){
        int[] array = {1, 3, 5, 8, 9};
        System.out.println(len(array));
    }
}
