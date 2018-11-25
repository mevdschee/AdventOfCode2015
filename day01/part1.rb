File.open("input", "r") do |f|
    n = 0
    while c = f.getc do
        n += case c
            when '(' then +1
            when ')' then -1
        end
    end
    puts n
end