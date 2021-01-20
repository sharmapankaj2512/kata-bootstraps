public class OcheReader {
    String input;
    public OcheReader(String echo) {
      input = echo;
    }

    public String getResponse() {
        StringBuilder sb = new StringBuilder(input);
        return sb.reverse().toString();
    }
}
