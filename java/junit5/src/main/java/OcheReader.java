import java.io.PrintStream;

public class OcheReader {
    String input;
    private PrintStream printStream;

    public OcheReader(String echo, PrintStream mock) {
        input = echo;
        printStream = mock;
        printStream.println(new StringBuilder(echo).reverse().toString());
    }

    public String getResponse() {
        StringBuilder sb = new StringBuilder(input);
        return sb.reverse().toString();
    }
}
