<?php
declare(strict_types = 1);

$inserts = $argv[1] ?? 1_000_000;

$host = $_ENV['DB_HOST'];
$db   = $_ENV['DB_NAME'];
$user = $_ENV['DB_USER'];
$pass = $_ENV['DB_PASSWORD'];
$charset = 'utf8mb4';

$dsn = "mysql:host=$host;dbname=$db;charset=$charset";
$connection = new \Pdo\Mysql($dsn, $_ENV['DB_USER'], $_ENV['DB_PASSWORD']);


if (!empty($connection->errorCode())) {
    die("Connection failed: " . json_encode($connection->errorInfo()));
}

$preparedQuery = $connection->prepare("INSERT INTO test (name) VALUES (?)");
for ($i = 0; $i < $inserts; $i++) {
    $preparedQuery->execute(['test ' . $i]);

    if ($i % 5_000 === 0) {
        $percentage = $i / $inserts * 100;
        echo date('Y/m/d H:i:s') . " Inserted $i/$inserts records ($percentage% )\n";
    }
}


