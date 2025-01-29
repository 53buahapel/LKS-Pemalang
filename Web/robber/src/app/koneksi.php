<?php

$conn = mysqli_connect('mysql', 'root', 'root', 'sqli');
mysqli_select_db($conn, $db);

$sql = "CREATE TABLE IF NOT EXISTS `user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(255) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    `email` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;";

$conn->query($sql);
$checklength = "SELECT COUNT(*) FROM user";
$result = $conn->query($checklength);
$length = $result->fetch_row()[0];

if ($length == 0) {
    $sql = "INSERT INTO user (username, password, email) VALUES ('LKS Pemalang', 'sUp3rs3cr3tBR0WnoneC4nthack', 'admin@localhost')";
    $conn->query($sql);
}

$sql = "CREATE TABLE IF NOT EXISTS `post` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `judul` varchar(255) DEFAULT NULL,
    `konten` text,
    `tanggal` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;";

$conn->query($sql);
$checklength = "SELECT COUNT(*) FROM post";
$result = $conn->query($checklength);
$length = $result->fetch_row()[0];

if ($length == 0) {
    $sql = "INSERT INTO post (judul, konten, tanggal) 
    VALUES 
    ('Ebook Hacker', 'Kalo mau jadi hengker harus belajar. Belajar dengan tekun, rajin, dan memiliki komitmen tinggi terhadap pekerjaan adalah hal yang penting. Menjadi seorang hengker bukan hanya soal teori, tapi juga praktik langsung di lapangan. Pahami bahwa setiap langkah dalam hidup ini adalah proses pembelajaran yang tak ada habisnya. Jangan lupa untuk terus berusaha dan jangan pernah berhenti untuk belajar.\n\n\n\n\n\n\n\n\n\n', NOW()),
    ('Ebook Hacker v2', 'WAF paling keren itu bukan dwaf. WAF (Web Application Firewall) adalah salah satu alat yang digunakan untuk melindungi aplikasi web dari berbagai jenis serangan. Memilih WAF yang tepat sangat penting untuk memastikan aplikasi kita tetap aman. Bukan hanya soal fitur, tetapi juga soal bagaimana WAF tersebut dapat mengidentifikasi dan mencegah serangan yang lebih kompleks seperti SQL injection, XSS, dan lainnya. WAF yang bagus haruslah memiliki kemampuan untuk memfilter trafik secara real-time tanpa mempengaruhi kinerja aplikasi.\n\n\n\n\n\n\n\n\n\n', NOW()),
    ('LKS Pemalang secret', 'Bg, plis jagan di hack hehe. Tapi kalo nemu vuln report aja bang karena sangat membantu. Keamanan sistem informasi adalah salah satu aspek yang paling penting di dunia digital saat ini. Banyak orang yang menganggap remeh masalah keamanan, padahal serangan terhadap sistem informasi bisa sangat merusak. Oleh karena itu, penting bagi kita untuk selalu berhati-hati dan menjaga data pribadi serta informasi penting kita dengan baik. Jika menemukan celah atau kerentanannya, jangan ragu untuk melaporkannya agar kita semua bisa belajar dan meningkatkan sistem kita bersama-sama.\n\n\n\n\n\n\n\n\n\n', NOW()),
    ('Bendera', 'Benderanya yang punya cuman admin. Bendera ini melambangkan sesuatu yang lebih dari sekedar simbol. Di baliknya terdapat nilai-nilai yang sangat penting untuk dipahami oleh setiap orang yang terlibat dalam sistem ini. Hanya admin yang dapat mengakses informasi lebih dalam tentang bendera ini, karena peran admin sangat vital dalam menjaga integritas dan keamanan dari sistem yang ada. Jangan coba-coba untuk memanipulasi atau mencoba mengambil alih bendera ini, karena konsekuensinya bisa sangat berbahaya.\n\n\n\n\n\n\n\n\n\n', NOW());";
    $conn->query($sql);
}

$sql = "CREATE TABLE IF NOT EXISTS `guestbook` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `tanggal` datetime NOT NULL,
    `nama` varchar(255) NOT NULL,
    `pesan` text NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=2 ;";

$conn->query($sql);
$checklength = "SELECT COUNT(*) FROM guestbook";
$result = $conn->query($checklength);
$length = $result->fetch_row()[0];

if ($length == 0) {
    $sql = "INSERT INTO guestbook (tanggal, nama, pesan) 
    VALUES (NOW(), 'LKS Pemalang', 'bg keknya salah jalan, ga ada flag disini. baca desc nya :v');";
    $conn->query($sql);
}
