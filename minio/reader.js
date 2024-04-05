const Minio = require('minio')

// Initialize MinIO client
const client = new Minio.Client({
  endPoint: '135.181.91.126',     //ip of the minio server
  port: 9000,
  useSSL: false,
  accessKey: 'uKrvnXQDDNPJ0Q5v3eyI',  //access key of minio
  secretKey: 'f2JIQ9HGZ8cTBDDzByJ1SZpuUgaHHQ1mgRrANABy' //secret key of minio
})

// Specify bucket name and file name
const bucketName = 'my-bucket'
const fileName = 'test.txt'

// Get object data and print content
client.getObject(bucketName, fileName, (err, dataStream) => {
  if (err) {
    return console.log(`MinIO error occurred: ${err}`)
  }
  
  let content = ''
  dataStream.on('data', (chunk) => {
    content += chunk
  })

  dataStream.on('end', () => {
    const jsonData = JSON.parse(content)   //convert to json
    const name = jsonData.email            //json key 
    console.log('Content of key:')
    console.log(`Value of key: ${name}`)
  })
})
