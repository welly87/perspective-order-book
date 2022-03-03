# java 11
java -cp tools/aeron-all-1.37.0.jar -Daeron.dir=./aeron-dir io.aeron.driver.MediaDriver

# java 16
java -cp tools/aeron-all-1.37.0.jar --add-opens java.base/sun.nio.ch=ALL-UNNAMED io.aeron.driver.MediaDriver

# java 17+
java -cp tools/aeron-all-1.37.0.jar -Daeron.dir=./aeron-dir --add-opens java.base/sun.nio.ch=ALL-UNNAMED io.aeron.driver.MediaDriver

aws s3 ls s3://scifin-aeron/

aws s3 ls s3://scifin-aeron --recursive --human-readable --summarize

aws s3 rm s3://scifin-aeron/aerons.parquet --recursive

scp file.txt remote_username@10.10.0.2:/remote/directory/newfilename.txt