lines <- suppressWarnings(readLines("input"))
input <- paste(lines, collapse="\n")
characters <- strsplit(input, "")[[1]]
n <- 0
for (ch in characters) {
    n <- n + switch(ch, '(' = 1, ')' = -1)
}
cat(n)