<?php
if (isset($_POST['encrypt'])) {
	$data = $_POST['encrypt'];
	$method = "AES-128-CTR";
	$key = "v3ry_s3cr3t_k3y_cann0t_be_gu3ssed";
	$option = 0;
	$iv = "3318201409900001";
	$encrypted = openssl_encrypt($data, $method, $key, $option, $iv);
	echo "$encrypted";
	exit;
}
