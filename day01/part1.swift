import Foundation

let fileURL = Bundle.main.url(forResource: "input", withExtension: "")
let input = try String(contentsOf: fileURL!, encoding: String.Encoding.utf8)

var n = 0
for c in input {
  if c == "(" {
    n += 1
  }
  else if c == ")" {
    n -= 1
  }
}

print(n)