#!/bin/sh

cd $INPUT_DIRECTORY

echo "Processing file $COMET_INPUT_FILE"
if [ -f $COMET_PARAMS ]; then
	echo "Running with a custom params file $COMET_PARAMS"
    /bin/comet -P$COMET_PARAMS $COMET_INPUT_FILE
else
	echo "Running with a standard params file comet.params"
	/bin/comet $COMET_INPUT_FILE
fi

echo "Deleting temp files"
rm "$BASE_INPUT_FILE.pin"
rm "$BASE_INPUT_FILE.txt"
rm "$BASE_INPUT_FILE.sqt"
rm -rf "$BASE_INPUT_FILE"

echo "Starting the promec-indexer process"
/bin/promec-indexer -pepxml "$BASE_INPUT_FILE.pep.xml" -host $ELS_HOST -index $INDEX_NAME