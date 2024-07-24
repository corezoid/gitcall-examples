<?php

function handle($data) {

    $n=10;
    $factorial=$n;
    while (--$n>1) $factorial=bcmul($factorial,$n);
    
    $data['factorial'] = $factorial;
    return $data;
}