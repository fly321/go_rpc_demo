<?php

class JsonRpc{
    private $conn = null;
    public function __construct(string $host, int $port)
    {
        $this->conn = fsockopen($host, $port, $errno, $errMsg,3);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call(string $method,$params){
        if (!$this->conn) {
            return false;
        }

        $err = fwrite($this->conn, json_encode(array('method' => $method, 'params'=>array($params),'id'=>0)).'\n');

        if ($err === false) {
            return false;
        }

        stream_set_timeout($this->conn, 0 , 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }

        return json_decode($line,true);

    }
}

$jsonRpc = new JsonRpc("127.0.0.1", 8080);
$res = $jsonRpc->Call("helloServer.SayHello", "我是php请求的数据");
print_r($res);