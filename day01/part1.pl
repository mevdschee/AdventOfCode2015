open my $input, '<', "input";
my $n = 0;
while (my $line = <$input>) {
    chomp $line;
    @characters = split(//, $line);
    for my $ch (@characters) {
        if ($ch eq "(") { 
            $n++;
        } elsif ($ch eq ")") { 
            $n--; 
        }
    }
}
print $n;