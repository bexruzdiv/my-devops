from minio import Minio
from minio.error import S3Error

# Initialize MinIO client
client = Minio(
    "135.181.91.126:9000",
    access_key="uKrvnXQDDNPJ0Q5v3eyI",
    secret_key="f2JIQ9HGZ8cTBDDzByJ1SZpuUgaHHQ1mgRrANABy",
    secure=False  # Change to True if MinIO server uses HTTPS
)

# Specify bucket name and file name
bucket_name = "my-bucket"
file_name = "test.txt"

try:
    # Get object data
    data = client.get_object(bucket_name, file_name)
    
    # Read and print object content
    content = data.read().decode('utf-8')
    print("Content of test.txt:")
    print(content)
    
except S3Error as err:
    print(f"MinIO error occurred: {err}")
