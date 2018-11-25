import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

class Code {

	public static void main(String[] args) throws IOException {
		String input = new String(Files.readAllBytes(Paths.get("input"))).trim();
		int n = 0;
		for (int i = 0; i < input.length(); i++) {
			switch (input.charAt(i)) {
				case '(': n++; break;
				case ')': n--; break;
			}
		}
		System.out.println(n);
	}

}