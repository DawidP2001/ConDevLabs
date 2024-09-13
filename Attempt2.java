import java.util.Map;
import java.util.Scanner;

// Name: Dawid Pionk
// Student Number: C00273530
// Date: 13/09/24

class Attempt2{
    public static void main(String args[]){
        Scanner in = new Scanner(System.in);
        String input; // Holds the Roman numerical
        int output = 0; // Holds the int output
        

        System.out.println("Enter number to convert");
        input = in.nextLine();
        input = input.toUpperCase();
        input += " "; // empty space added so strign array wont run out of bounds in the loop below

        // Main loop
        // Checks for the 6 exceptions and adds the rest
        for(int i=0;i<input.length();i++){
            if(input.charAt(i) == 'I'){
                if(input.charAt(i+1) == 'V'){
                    output +=4;
                    i+=1;
                }
                else if(input.charAt(i+1)=='X'){
                    output += 9;
                    i+=1;
                }
                else {
                    output += 1;
                }
            }
            else if(input.charAt(i)=='V'){
                output += 5;
            }
            else if(input.charAt(i) == 'X'){
                if(input.charAt(i+1) == 'L'){
                    output += 40;
                    i+=1;
                }
                else if(input.charAt(i+1)=='C'){
                    output += 90;
                    i+=1;
                }
                else {
                    output += 10;
                }
            }
            else if(input.charAt(i)=='L'){
                output += 50;
            }
            else if(input.charAt(i)=='C'){
                if(input.charAt(i+1) == 'D'){
                    output += 400;
                    i+=1;
                }
                else if(input.charAt(i+1)=='M'){
                    output += 900;
                    i+=1;
                }
                else {
                    output += 100;
                }
            }
            else if(input.charAt(i)=='D'){
                output += 500;
            }
            else if(input.charAt(i)=='M'){
                output += 1000;
            }

        }
        System.out.println(input + "as an integer is: " + output);
    }
}