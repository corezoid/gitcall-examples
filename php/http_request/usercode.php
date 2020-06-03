<?php

require '/libs/vendor/autoload.php';
use GuzzleHttp\Client;

function handle($data) {

    $client = new Client();

    $response = $client->request('GET', 'https://reqres.in/api/users?page=1');
    
    $data['res'] = json_decode($response->getBody(), true);

    return $data;
}