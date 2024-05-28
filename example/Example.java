public class Example {
    public static void main(String[] args) {
        // Unused variable (Code Smell)
        int unusedVariable = 10;

        // Example of a potential bug: Division by zero
        int divisor = 0;
        int result = 10 / divisor;

        // Example of code smell: Nested loops
        for (int i = 0; i < 5; i++) {
            for (int j = 0; j < 5; j++) {
                System.out.println("Nested loop i: " + i + ", j: " + j);
            }
        }

        // Example of code smell: Long method
        longMethod();

        // Example of bug: Null pointer dereference
        String str = null;
        System.out.println(str.length());
    }

    public static void longMethod() {
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
        System.out.println("This is a long method with too many statements.");
    }
}
