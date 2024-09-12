import java.util.Map;
import java.util.Scanner;

class a{
    static Map<Character, Integer> map = Map.of(
        'I', 1,
        'V', 5,
        'X', 10,
        'L', 50,
        'C', 100,
        'D', 500,
        'M', 1000
    );
    public static char checkIfLarger(char largest, char check){
        int largestInt = map.get(largest);
        int checkInt = map.get(check);
        if (checkInt > largestInt){
            return check;
        }
        return largest;
    }
    public static void main(String args[]){
        Scanner in = new Scanner(System.in);
        String input;
        int output = 0;
        

        System.out.println("Enter number to convert");
        input = in.nextLine();
        
        char largest = input.charAt(0);
        for(int i=0;i<input.length();i++){
            char check = input.charAt(i);
            largest = checkIfLarger(largest, check);
        }
        boolean before = true;
        for(int i=0;i<input.length();i++){
            int num = map.get(input.charAt(i));
            if(input.charAt(i)==largest){
                before = false;
            } 
            if(before){
                output -= num;
            } else {
                output += num;
            }
        }
        System.out.println(output);
    }
}