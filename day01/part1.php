<?php
$n = 0;
while ($c = fgetc(STDIN)) {
    switch ($c) {
        case '(':
            $n++;
            break;
        case ')':
            $n--;
            break;
    }
}
echo $n;
