import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Code {

	public static void main(String[] args) throws IOException {
		String input = new String(Files.readAllBytes(Paths.get("input"))).trim();
		int n = 0;
		for (int i = 0; i < input.length(); i++) {
			n+=(input.charAt(i)=='('?1:(input.charAt(i)==')'?-1:0));
		}
		System.out.println(n);
	}

}