<?php

function handle($data) {
    throw new Exception('my custom error');
    return $data;
}