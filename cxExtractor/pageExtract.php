<?php

require_once( 'class.textExtract.php' );
//$iTextExtractor = new textExtract( $_POST['url'] );
$input = $argv[1];
//$output = $argv[2];

$iTextExtractor = new textExtract($input);
$text = $iTextExtractor->getPlainText();

if( $iTextExtractor->isGB ) $text = iconv( 'GBK', 'UTF-8//IGNORE', $text );
echo $text
	//file_put_contents($output,$text)
?>
