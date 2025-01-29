<?php

$page = $_GET['page'] ?? 'home';

include __DIR__ . '/pages/' . $page . '.php';
