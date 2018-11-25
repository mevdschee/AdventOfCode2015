<?php
$n = 0;
while (($c = fgetc(STDIN)) !== false) {
    $n += ($c == '(' ? 1 : ($c == ')' ? -1 : 0));
}
echo $n;
