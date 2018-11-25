import java.io.File

fun main(args: Array<String>) {
    val input = File("input").readLines()
    var n = 0
    input.forEach { str ->
        str.forEach { c ->
            when (c) {
                ')' -> {
                    n--
                }
                '(' -> {
                    n++
                }
            }
        }
    }
    println("${n}\n")
}