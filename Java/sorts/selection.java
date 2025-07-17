import java.util.ArrayList;

public class selection{
    
    public static int findSmallest(ArrayList<Integer> array){
        int smallest = array.get(0);
        int smallest_index = 0;

        for(int i = 1; i < array.size(); i++){
            if (smallest < array.get(i)) {
                smallest = array.get(i);
                smallest_index = i;
            }
        }
        return smallest_index;
    }

    public static ArrayList<Integer> selectionSort(ArrayList<Integer> array){
        ArrayList<Integer> new_arr = new ArrayList<>();
        int size = array.size();       // Если array.size() вставить в цикл, то цикл будет брать длину каждую иттерацию

        for(int i = 0; i < size; i++){
            int smallest = findSmallest(array);
            new_arr.add(array.remove(smallest));
        }
        return new_arr;
    }

    public static void main(String[] args){

        ArrayList<Integer> arr = new ArrayList<>();
        arr.add(7);
        arr.add(8);
        arr.add(9);
        arr.add(4);
        arr.add(6);
        arr.add(1);
        arr.add(0);

        System.out.println(arr);
        System.out.println(selectionSort(arr));

    }
}