<?php

function handle($taskId, $data) {
    $data['php'] = "Hello, world!";
    $data['phpTaskId'] = $taskId;
    return $data;
}
